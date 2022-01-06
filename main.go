package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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
		"dart":       &model.DartWriter{},
		"java":       &model.JavaWriter{},
		"kotlin":     &model.KotlinWriter{},
		"python":     &model.PythonWriter{},
		"protobuf":   &model.ProtobufWriter{},
		"typescript": &model.TypeScriptWriter{},
		// "ruby":       &model.RubyWriter{},
		// "csharp":     &model.CSharpWriter{},
		// "swift":      &model.SwiftWriter{},
		// "php":        &model.PhpWriter{},
		// "scala":      &model.ScalaWriter{},
		// "rust":       &model.RustWriter{},
	}
)

func init() {
	flag.StringVar(&outputPath, "o", ".", "Specify ouput file path.")
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

	fmt.Println(inputPath)

	jsonFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
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
}
