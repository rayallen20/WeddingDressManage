package sysError

import (
	"WeddingDressManage/lib/helper/paramHelper"
	"reflect"
)

// UnmarshalTypeError 校验请求参数时 JSON反序列化无效错误 属于程序外部错误
// 错误成因:前端传过来的JSON中 字段类型与后端定义的不匹配
type UnmarshalTypeError struct {
	// Value 前端传过来的字段类型
	Value  string
	// Type 后端定义的类型
	Type   reflect.Type
	// Struct 类型不匹配的字段所在的结构体
	Struct string
	// Field 类型不匹配的字段
	Field  string
	// 类型不匹配的字段对应的errField tag值
	errField string
}

func (u UnmarshalTypeError) Error() string {
	return u.errField + " should be a " + u.Type.String() + ", not a " + u.Value
}

func (u *UnmarshalTypeError) SetMsg(objs []interface{}) {
	u.errField = paramHelper.GetErrFieldValue(objs, u.Field)
}
