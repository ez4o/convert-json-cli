package model

import (
	"fmt"
	"os"
)

type RustWriter struct {
	outputPath string
	nested     bool
}

func (rw *RustWriter) SetOutputPath(outputPath string) {
	rw.outputPath = outputPath
}

func (rw *RustWriter) SetNested(nested bool) {
	rw.nested = nested
}

func (rw *RustWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(rw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("use serde_derive::{Deserialize, Serialize};\n\n")
	if err != nil {
		return err
	}

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			_, err = file.WriteString("pub type Nested = Vec<AutoGenerated>;\n\n")
			if err != nil {
				return err
			}
			continue
		}

		_, err := file.WriteString(rw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (rw *RustWriter) GetStruct(abstractStruct Struct) string {
	deriveAnnotation := "#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]\n"
	serdesAnnotation := "#[serde(rename_all = \"camelCase\")]\n"

	return deriveAnnotation + serdesAnnotation + "pub struct " + abstractStruct.Name + " {\n" + rw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (rw *RustWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += rw.GetField(i, field)
	}

	return result
}

func (rw *RustWriter) GetField(_ int, field Field) string {
	var typeName string = ""
	var index string = ""
	var serdesAnnotation string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "Vec<" + rw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ">"
	} else {
		typeName = rw.GetTypeName(field.TypeName)
	}

	if field.Index[0] == '_' {
		serdesAnnotation = fmt.Sprintf("\t#[serde(rename = \"%s\")]\n", field.Index)
		index = field.Index[1:]
	} else {
		index = field.Index
	}

	return serdesAnnotation + "\tpub " + index + ": " + typeName + ",\n"
}

func (rw *RustWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int16":
		return "i16"
	case "int32":
		return "i32"
	case "int64":
		return "i64"
	case "float32":
		return "f32"
	case "float64":
		return "f64"
	case "string":
		return "String"
	case "bool":
		return "bool"
	default:
		return typeName
	}
}