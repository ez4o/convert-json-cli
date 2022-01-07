package model

import "os"

type JavaWriter struct {
	outputPath string
}

func (jw *JavaWriter) SetOutputPath(outputPath string) {
	jw.outputPath = outputPath
}

func (jw *JavaWriter) SetNested(nested bool) {}

func (jw *JavaWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(jw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("import java.util.ArrayList;\n\n")
	if err != nil {
		return err
	}

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(jw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (jw *JavaWriter) GetStruct(abstractStruct Struct) string {
	return "public class " + abstractStruct.Name + " {\n" + jw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (jw *JavaWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += jw.GetField(i, field)
	}

	return result
}

func (jw *JavaWriter) GetField(_ int, field Field) string {
	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		return "\tpublic ArrayList<" + jw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + "> " + field.Index + " = new ArrayList<" + jw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ">();\n"
	} else {
		return "\tpublic " + jw.GetTypeName(field.TypeName) + " " + field.Index + ";\n"
	}
}

func (jw *JavaWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int16":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "long"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "bool":
		return "boolean"
	case "string":
		return "String"
	default:
		return typeName
	}
}
