package sysError

import "reflect"

// InvalidUnmarshalError 校验请求参数时 JSON反序列化无效错误 属于程序内部错误
// 错误成因:controller层实例化参数时 实例化了一个nil指针
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (i InvalidUnmarshalError) Error() string {
	return "json.Unmarshal:(nil " + i.Type.String() + ")"
}

