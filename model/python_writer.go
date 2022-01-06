package model

import "os"

type PythonWriter struct{}

func (pw *PythonWriter) SetNested(nested bool) {}

func (pw *PythonWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create("out/python_result.py")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("from typing import List\n\n")
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

func (pw *PythonWriter) GetStruct(abstractStruct Struct) string {
	return "class " + abstractStruct.Name + ":\n" + pw.GetFields(abstractStruct.Fields) + "\n" + pw.getConstructor(abstractStruct) + "\n\n"
}

func (pw *PythonWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += pw.GetField(i, field)
	}

	return result
}

func (pw *PythonWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "List[" + pw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + "]"
	} else {
		typeName = pw.GetTypeName(field.TypeName)
	}

	return "\t" + field.Index + ": " + typeName + "\n"
}

func (pw *PythonWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "string":
		return "str"
	case "bool":
		return "bool"
	case "float32":
		return "float"
	case "float64":
		return "float"
	default:
		return typeName
	}
}

func (pw *PythonWriter) getConstructor(abstractStruct Struct) string {
	var result string = ""

	result += "\tdef __init__(self, "

	for _, field := range abstractStruct.Fields {
		result += pw.getConstructorParams(field)
	}

	result += ") -> None:\n"

	for _, field := range abstractStruct.Fields {
		result += pw.getConstructorAssignments(field)
	}

	return result
}

func (pw *PythonWriter) getConstructorParams(field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "List[" + pw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + "]"
	} else {
		typeName = pw.GetTypeName(field.TypeName)
	}

	return " " + field.Index + ": " + typeName + ","
}

func (pw *PythonWriter) getConstructorAssignments(field Field) string {
	return "\t\tself." + field.Index + " = " + field.Index + "\n"
}
