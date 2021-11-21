package response

import "WeddingDressManage/lib/sysError"

type RespBody struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}

// 错误码规则:
// 5XX:程序内部错误
// 101XX:数据库错误
// 102XX:参数校验错误
// 103XX:业务逻辑错误
const (
	// InvalidUnmarshalError JSON反序列化错误
	InvalidUnmarshalError = 501

	// FieldTypeError 字段类型错误
	FieldTypeError = 10201

	// ValidateError 参数校验错误
	ValidateError = 10202
)

var message = map[int]string{
	ValidateError: "validate error",
}

func (r *RespBody) InvalidUnmarshalError(err *sysError.InvalidUnmarshalError)  {
	r.Code = InvalidUnmarshalError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) FieldTypeError(err *sysError.UnmarshalTypeError) {
	r.Code = FieldTypeError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) ValidateError(errs []*sysError.ValidateError)  {
	validateFailInfos := make([]string, 0, len(errs))
	for _, validateError := range errs {
		validateFailInfo := validateError.Error()
		validateFailInfos = append(validateFailInfos, validateFailInfo)
	}
	r.Code = ValidateError
	r.Message = message[ValidateError]
	r.Data = map[string]interface{}{
		"validateFailInfos":validateFailInfos,
	}
}