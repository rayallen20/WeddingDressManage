package wdmError

import "WeddingDressManage/lib/structHelper"

// DBError 数据库错误
type DBError struct {
	Message string
}

func (d DBError) Error() string {
	return d.Message
}

// FileTypeError 上传文件类型错误
type FileTypeError struct {
	Message string
}

func (f FileTypeError) Error() string {
	return f.Message
}

// SaveFileError 保存文件错误
type SaveFileError struct {
	Message string
}

func (s SaveFileError) Error() string {
	return s.Message
}

// BindingValidatorError 将绑定错误转化为校验器错误时失败产生的错误
type BindingValidatorError struct {
	Message string
}

func (b BindingValidatorError) Error() string {
	return b.Message
}

// ParamValueError 参数值错误
type ParamValueError struct {
	Message string
	Details []string
}

func (p ParamValueError) Error() string {
	return p.Message
}

type NumericStringError struct {
	Message string
	NotNumericFields []string
}

func (n NumericStringError) Error() string {
	return n.Message
}

type ParamTypeError struct {
	Message string
	StructFieldName string
	FormFieldName string
	ShouldType string
}

func (j ParamTypeError) Error() string {
	return j.Message
}

func (j *ParamTypeError) GetFormFieldAndShouldType(obj interface{})  {
	// TODO:此处未做错误处理
	fieldInfo, _ := structHelper.GetFieldAndTag(obj, j.StructFieldName, "form")
	j.FormFieldName = fieldInfo["tagContent"]
	j.ShouldType = fieldInfo["fieldType"]
}


