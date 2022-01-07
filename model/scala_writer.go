package model

import "os"

type ScalaWriter struct {
	outputPath string
}

func (sw *ScalaWriter) SetOutputPath(outputPath string) {
	sw.outputPath = outputPath
}

func (sw *ScalaWriter) SetNested(nested bool) {}

func (sw *ScalaWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(sw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(sw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (sw *ScalaWriter) GetStruct(abstractStruct Struct) string {
	return "case class " + abstractStruct.Name + "(\n" + sw.GetFields(abstractStruct.Fields) + ")\n\n"
}

func (sw *ScalaWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += sw.GetField(i, field)
	}

	return result
}

func (sw *ScalaWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "Seq[" + sw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + "]"
	} else {
		typeName = sw.GetTypeName(field.TypeName)
	}

	return "\t" + field.Index + ": " + typeName + ",\n"
}

func (sw *ScalaWriter) GetTypeName(typeName string) string {
	// https://docs.scala-lang.org/overviews/scala-book/built-in-types.html
	switch typeName {
	case "int":
		return "Short"
	case "int32":
		return "Int"
	case "int64":
		return "Long"
	case "float32":
		return "Float"
	case "float64":
		return "Double"
	case "string":
		return "String"
	case "bool":
		return "Boolean"
	default:
		return typeName
	}
}
