package model

import "os"

type KotlinWriter struct {
	nester bool
}

func (kw *KotlinWriter) SetNested(nested bool) {
	kw.nester = nested
}

func (kw *KotlinWriter) Write(abstractStructs []Struct) error {
	file, err := os.Create("out/kotlin_result.kt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("import com.fasterxml.jackson.annotation.JsonProperty\n\n")
	if err != nil {
		return err
	}

	for _, abstractStruct := range abstractStructs {
		if abstractStruct.Name == "Nested" {
			_, err := file.WriteString("typealias Nested = List<AutoGenerated>;\n\n")
			if err != nil {
				return err
			}
			continue
		}

		_, err := file.WriteString(kw.GetStruct(abstractStruct))
		if err != nil {
			return err
		}
	}

	return nil
}

func (kw *KotlinWriter) GetStruct(abstractStruct Struct) string {
	return "data class " + abstractStruct.Name + "(\n" + kw.GetFields(abstractStruct.Fields) + ")\n\n"
}

func (kw *KotlinWriter) GetFields(fields []Field) string {
	var result string = ""

	for i, field := range fields {
		result += kw.GetField(i, field)
	}

	return result
}

func (kw *KotlinWriter) GetField(_ int, field Field) string {
	typeName := ""
	jsonAnnotation := ""
	index := ""

	if field.TypeName[len(field.TypeName)-2:] == "[]" {
		typeName = "List<" + kw.GetTypeName(field.TypeName[:len(field.TypeName)-2]) + ">"
	} else {
		typeName = kw.GetTypeName(field.TypeName)
	}

	if field.Index[0] == '_' {
		jsonAnnotation = "\t@JsonProperty(\"" + field.Index + "\")\n"
		index = field.Index[1:]
	} else {
		index = field.Index
	}

	return jsonAnnotation + "\tval " + index + ": " + typeName + ",\n"
}

func (kw *KotlinWriter) GetTypeName(typeName string) string {
	switch typeName {
	case "int":
		return "Int"
	case "int32":
		return "Int"
	case "int64":
		return "Long"
	case "string":
		return "String"
	case "bool":
		return "Boolean"
	case "float":
		return "Float"
	case "double":
		return "Double"
	default:
		return typeName
	}
}