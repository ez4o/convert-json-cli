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
		"c":          &model.CWriter{},
		"go":         &model.GoWriter{},
		"cpp":        &model.CppWriter{},
		"php":        &model.PHPWriter{},
		"dart":       &model.DartWriter{},
		"java":       &model.JavaWriter{},
		"rust":       &model.RustWriter{},
		"scala":      &model.ScalaWriter{},
		"swift":      &model.SwiftWriter{},
		"csharp":     &model.CSharpWriter{},
		"kotlin":     &model.KotlinWriter{},
		"python":     &model.PythonWriter{},
		"protobuf":   &model.ProtobufWriter{},
		"typescript": &model.TypeScriptWriter{},
	}
	extensions map[string]string = map[string]string{
		"c":          ".c",
		"go":         ".go",
		"cpp":        ".cpp",
		"php":        ".php",
		"dart":       ".dart",
		"java":       ".java",
		"rust":       ".rs",
		"scala":      ".scala",
		"swift":      ".swift",
		"csharp":     ".cs",
		"kotlin":     ".kt",
		"python":     ".py",
		"protobuf":   ".proto",
		"typescript": ".ts",
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
		panic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		outputPath = strings.TrimSuffix(inputPath, ".json") + extensions[targetLanguage]
	}

	w := writers[targetLanguage]
	if w == nil {
		panic("Target language is not supported.")
	}
	w.SetOutputPath(outputPath)

	jc := model.JSONConverter{Writer: w}
	err = jc.Convert(string(bytes))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully convert %s to %s.\n", inputPath, outputPath)
}
