package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"ez4o.com/convert-json-cli/model"
)

var (
	inputPath      string
	outputPath     string
	targetLanguage string
	writers        map[string]model.IWriter = map[string]model.IWriter{
		"c":           &model.CWriter{},
		"go":          &model.GoWriter{},
		"cpp":         &model.CppWriter{},
		"php":         &model.PHPWriter{},
		"dart":        &model.DartWriter{},
		"java":        &model.JavaWriter{},
		"rust":        &model.RustWriter{},
		"scala":       &model.ScalaWriter{},
		"swift":       &model.SwiftWriter{},
		"csharp":      &model.CSharpWriter{},
		"kotlin":      &model.KotlinWriter{},
		"python":      &model.PythonWriter{},
		"protobuf":    &model.ProtobufWriter{},
		"flatbuffers": &model.FlatbuffersWriter{},
		"typescript":  &model.TypeScriptWriter{},
	}
	extensions map[string]string = map[string]string{
		"c":           ".c",
		"go":          ".go",
		"cpp":         ".cpp",
		"php":         ".php",
		"dart":        ".dart",
		"java":        ".java",
		"rust":        ".rs",
		"scala":       ".scala",
		"swift":       ".swift",
		"csharp":      ".cs",
		"kotlin":      ".kt",
		"python":      ".py",
		"protobuf":    ".proto",
		"flatbuffers": ".fbs",
		"typescript":  ".ts",
	}
)

func init() {
	flag.StringVar(&outputPath, "o", "", "Specify ouput file path.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: convjson [OPTIONS] [INPUT_FILE_PATH] [TARGET_LANGUAGE]\n")

	fmt.Fprintf(os.Stderr, "  INPUT_FILE_PATH\n")
	fmt.Fprintf(os.Stderr, "\tSpecify input file path.\n")

	fmt.Fprintf(os.Stderr, "  TARGET_LANGUAGE\n")
	fmt.Fprintf(os.Stderr, "\t")
	for key := range writers {
		fmt.Fprintf(os.Stderr, "[%s] ", key)
	}
	fmt.Fprintf(os.Stderr, "\n")

	flag.PrintDefaults()
}

func handlePanic(err interface{}) {
	if _, ok := err.(error); ok {
		fmt.Fprintf(os.Stderr, "Could not read input file: %s\n", err.(error).Error())
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	os.Exit(2)
}

func main() {
	flag.Parse()

	inputPath = flag.Arg(0)
	targetLanguage = flag.Arg(1)

	if inputPath == "" || targetLanguage == "" {
		flag.Usage()
		return
	}

	jsonFile, err := os.Open(inputPath)
	if err != nil {
		handlePanic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		handlePanic(err)
	}

	if outputPath == "" {
		outputPath = strings.TrimSuffix(inputPath, ".json") + extensions[targetLanguage]
	}

	w := writers[targetLanguage]
	if w == nil {
		handlePanic("Target language is not supported.")
	}
	w.SetOutputPath(outputPath)

	jc := model.JSONConverter{Writer: w}
	err = jc.Convert(string(bytes))
	if err != nil {
		handlePanic(err)
	}

	fmt.Printf("Successfully convert %s to %s.\n", inputPath, outputPath)
}
