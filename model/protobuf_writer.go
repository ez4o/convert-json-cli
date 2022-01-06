package model

import (
	"fmt"
	"os"
)

type ProtobufWriter struct{}

func (pw *ProtobufWriter) SetNested(nested bool) {}

func (pw *ProtobufWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create("out/protobuf_result.proto")
	if err != nil {
		return err
	}
	defer file.Close()

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
	case "int":
	case "int8":
	case "int16":
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "float32":
	case "float64":
		return "double"
	case "string":
		return "string"
	case "bool":
		return "bool"
	}

	return typeName
}
