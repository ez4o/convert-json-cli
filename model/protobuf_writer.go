package model

import (
	"fmt"
	"os"
)

type ProtobufWriter struct {
	outputPath string
}

func (pw *ProtobufWriter) SetOutputPath(outputPath string) {
	pw.outputPath = outputPath
}

func (pw *ProtobufWriter) SetNested(nested bool) {}

func (pw *ProtobufWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(pw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("syntax = \"proto3\";\n\n")
	if err != nil {
		return err
	}

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(pw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (pw *ProtobufWriter) GetStruct(abstractStruct Struct) string {
	return "message " + abstractStruct.Name + " {\n" + pw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (pw *ProtobufWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += pw.GetField(i, field)
	}

	return result
}

func (pw *ProtobufWriter) GetField(i int, field Field) string {
	var repeated string = ""
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		repeated = "repeated "
		typeName = field.TypeName[:len(field.TypeName)-2]
	} else {
		typeName = field.TypeName
	}

	return "  " + repeated + pw.GetTypeName(typeName) + " " + field.Index + " = " + fmt.Sprintf("%d", i+1) + ";\n"
}

func (pw *ProtobufWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int16":
		return "int32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "float32":
		return "double"
	case "float64":
		return "double"
	case "string":
		return "string"
	case "bool":
		return "bool"
	}

	return typeName
}
