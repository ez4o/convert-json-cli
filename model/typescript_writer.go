package model

import "os"

type TypeScriptWriter struct {
	outputPath string
	nested     bool
}

func (tsw *TypeScriptWriter) SetOutputPath(outputPath string) {
	tsw.outputPath = outputPath
}

func (tsw *TypeScriptWriter) SetNested(nested bool) {
	tsw.nested = nested
}

func (tsw *TypeScriptWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(tsw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			_, err := file.WriteString("export type Nested = AutoGenerated[]\n\n")
			if err != nil {
				return err
			}
			continue
		}

		_, err := file.WriteString(tsw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (tsw *TypeScriptWriter) GetStruct(abstractStruct Struct) string {
	return "export interface " + abstractStruct.Name + " {\n" + tsw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (tsw *TypeScriptWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += tsw.GetField(i, field)
	}

	return result
}

func (tsw *TypeScriptWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = tsw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + "[]"
	} else {
		typeName = tsw.GetTypeName(field.TypeName)
	}

	return "\t" + field.Index + ": " + typeName + ";\n"
}

func (tsw *TypeScriptWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int16":
		return "number"
	case "int32":
		return "number"
	case "int64":
		return "number"
	case "float32":
		return "number"
	case "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	default:
		return typeName
	}
}
