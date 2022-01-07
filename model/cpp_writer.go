package model

import "os"

type CppWriter struct {
	outputPath string
}

func (cw *CppWriter) SetOutputPath(outputPath string) {
	cw.outputPath = outputPath
}

func (cw *CppWriter) SetNested(nested bool) {}

func (cw *CppWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create(cw.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("#include <iostream>\n#include <vector>\n#include <string>\n\n")
	if err != nil {
		return err
	}

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(cw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (cw *CppWriter) GetStruct(abstractStruct Struct) string {
	return "class " + abstractStruct.Name + " {\npublic:\n" + cw.GetFields(abstractStruct.Fields) + "};\n\n"
}

func (cw *CppWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += cw.GetField(i, field)
	}

	return result
}

func (cw *CppWriter) GetField(_ int, field Field) string {
	typeName := ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "std::vector<" + cw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ">"
	} else {
		typeName = cw.GetTypeName(field.TypeName)
	}

	return "\t" + typeName + " " + field.Index + ";\n"
}

func (cw *CppWriter) GetTypeName(typeName string) string {
	// https://en.cppreference.com/w/cpp/language/types
	switch typeName {
	case "int16":
		return "short int"
	case "int32":
		return "int"
	case "int64":
		return "long long int"
	case "string":
		return "std::string"
	case "bool":
		return "bool"
	case "float32":
		return "float"
	case "float64":
		return "double"
	default:
		return typeName
	}
}
