package model

import "os"

type CWriter struct{}

func (cw *CWriter) SetNested(nested bool) {}

func (cw *CWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create("out/c_result.c")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("#include <stdio.h>\n\n")
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

func (cw *CWriter) GetStruct(abstractStruct Struct) string {
	return "struct " + abstractStruct.Name + " {\n" + cw.GetFields(abstractStruct.Fields) + "};\n\n"
}

func (cw *CWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += cw.GetField(i, field)
	}

	return result
}

func (cw *CWriter) GetField(_ int, field Field) string {
	typeName := ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = cw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + "*"
	} else {
		typeName = cw.GetTypeName(field.TypeName)
	}

	return "\t" + typeName + " " + field.Index + ";\n"
}

func (cw *CWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int":
		return "short int"
	case "int32":
		return "int"
	case "int64":
		return "long long"
	case "string":
		return "char *"
	case "bool":
		return "int"
	case "float32":
		return "float"
	case "float64":
		return "double"
	}

	return "struct " + typeName
}
