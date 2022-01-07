package model

import "os"

type PHPWriter struct {
	outputPath string
}

func (pw *PHPWriter) SetOutputPath(outputPath string) {
	pw.outputPath = outputPath
}

func (pw *PHPWriter) SetNested(nested bool) {}

func (pw *PHPWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(pw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("<?php\n\n")
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

func (pw *PHPWriter) GetStruct(abstractStruct Struct) string {
	return "class " + abstractStruct.Name + " {\n" + pw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (pw *PHPWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += pw.GetField(i, field)
	}

	return result
}

func (pw *PHPWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "array(" + pw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ")"
	} else {
		typeName = pw.GetTypeName(field.TypeName)
	}

	return "\tpublic $" + field.Index + "; // " + typeName + "\n"
}

func (pw *PHPWriter) GetTypeName(typeName string) string {
	// https://www.php.net/manual/en/language.types.intro.php
	switch typeName {
	case "int16":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "float32":
		return "float"
	case "float64":
		return "float"
	case "string":
		return "string"
	case "bool":
		return "bool"
	default:
		return typeName
	}
}
