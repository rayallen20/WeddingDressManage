package paramHelper

import (
	"reflect"
	"strings"
)

// GetErrFieldValue 根据类型不匹配的字段名 获取该字段的自定义tag:errField的值
func GetErrFieldValue(objs []interface{}, exceptionFieldName string) string {
	var res string
	var found bool
	fieldNameAndErrField := extractFieldNameAndErrField(objs)
	exceptionFieldNames := strings.Split(exceptionFieldName, ".")

	for index, exceptionFieldName := range exceptionFieldNames {
		found = false
		for _, v := range fieldNameAndErrField {
			for k1, v1 := range v {
				if k1 == exceptionFieldName {
					found = true
					if index == len(exceptionFieldNames) - 1 {
						res += v1
					} else {
						res += v1 + "."
					}
					break
				}
			}

			if found {
				break
			}
		}
	}
	return res
}

// extractFieldNameAndErrField 使用反射机制 提取结构体中 字段名和自定义tag:errField的映射关系
func extractFieldNameAndErrField(params []interface{}) map[string]map[string]string {
	res := map[string]map[string]string{}

	for i := 0; i < len(params); i++ {
		structReflectInfo := reflect.TypeOf(params[i])
		structName := strings.Split(structReflectInfo.String(), ".")[1]

		tmpRes := map[string]string{}
		for j := 0; j < structReflectInfo.NumField(); j++ {
			fieldName := structReflectInfo.Field(j).Name
			errFieldName := structReflectInfo.Field(j).Tag.Get("errField")
			tmpRes[fieldName] = errFieldName
		}
		res[structName] = tmpRes
	}

	return res
}
