package model

import "os"

type CSharpWriter struct {
	outputPath string
}

func (csw *CSharpWriter) SetOutputPath(outputPath string) {
	csw.outputPath = outputPath
}

func (csw *CSharpWriter) SetNested(nested bool) {}

func (csw *CSharpWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(csw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(csw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (csw *CSharpWriter) GetStruct(abstractStruct Struct) string {
	return "public class " + abstractStruct.Name + " {\n" + csw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (csw *CSharpWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += csw.GetField(i, field)
	}

	return result
}

func (csw *CSharpWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "List<" + csw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ">"
	} else {
		typeName = csw.GetTypeName(field.TypeName)
	}

	return "\tpublic " + typeName + " " + field.Index + ";\n"
}

func (csw *CSharpWriter) GetTypeName(typeName string) string {
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
	case "string":
		return "String"
	case "bool":
		return "boolean"
	default:
		return typeName
	}
}
