package structHelper

import "reflect"

// StructAssign 使用反射机制 将结构体value的字段值 复制给结构体binding
// 前提:两个结构体中的字段名相同
func StructAssign(binding interface{}, value interface{}) {
	// 获取reflect.Type类型
	bVal := reflect.ValueOf(binding).Elem()
	vVal := reflect.ValueOf(value).Elem()

	vTypeOfT := vVal.Type()

	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段 有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}
