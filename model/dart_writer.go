package model

import "os"

type DartWriter struct{}

func (dw *DartWriter) SetNested(nested bool) {}

func (dw *DartWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create("out/dart_result.dart")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			continue
		}

		_, err := file.WriteString(dw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (dw *DartWriter) GetStruct(abstractStruct Struct) string {
	return "class " + abstractStruct.Name + " {\n" + dw.GetFields(abstractStruct.Fields) + "}\n\n"
}

func (dw *DartWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += dw.GetField(i, field)
	}

	return result
}

func (dw *DartWriter) GetField(_ int, field Field) string {
	var typeName string = ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "List<" + dw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ">?"
	} else {
		typeName = dw.GetTypeName(field.TypeName) + "?"
	}

	return "\t" + typeName + " " + field.Index + ";\n"
}

func (dw *DartWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "string":
		return "String"
	case "bool":
		return "bool"
	case "float32":
		return "double"
	case "float64":
		return "double"
	}

	return typeName
}
