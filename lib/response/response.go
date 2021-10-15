package response

import (
	"WeddingDressManage/lib/wdmError"
)

type ResBody struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}

// 非正常响应状态码定义规则:
// 100XX:校验参数错误
// 101XX:数据库错误
// 102XX:接收文件错误
// 103XX:业务逻辑错误(业务逻辑错误:即在请求参数的类型和值均能通过校验的前提下,仍旧导致系统无法描述业务逻辑的错误)
const (
	// Success 正常响应
	Success = 200

	// DBError 数据库操作错误
	DBError = 10100

	// TransactionError 事务错误
	TransactionError = 10101

	// BindingValidatorError 将绑定错误转化为校验器错误时失败产生的错误
	BindingValidatorError = 10101

	// NumericStringError 一个字符串类型的字段无法被转化为整型
	// 形如"001"的字段 传值时传了个"AAA"
	NumericStringError = 10102

	// ParamValueError 校验请求参数错误
	ParamValueError = 10103

	// ParamTypeError 参数类型错误
	ParamTypeError = 10104

	// ReceiveFileError 接收文件错误
	ReceiveFileError = 10200

	// SaveFileError 保存文件错误
	SaveFileError = 10201

	// FileTypeError 文件类型错误
	FileTypeError = 10001

	// KindIsNotExist 品类名称与编码信息不存在
	KindIsNotExist = 10301

	// CategoryHasExisted 品类信息已存在
	CategoryHasExisted = 10302
)

var Message = map[int]string {
	Success:            "success",
	FileTypeError:      "file type must be jpg, jpeg or png",
	NumericStringError: "following fields are not numeric data in string objects",
	ParamValueError:    "param value error",
	ParamTypeError:     "param type error",
	KindIsNotExist:     "kind is not exist",
	CategoryHasExisted: "category has existed",
}

func (r *ResBody) DBError(err error, data map[string]interface{}) {
	r.Code = DBError
	r.Message = err.Error()
	r.Data = data
}

func (r *ResBody) TransactionError(err error, data map[string]interface{}) {
	r.Code = TransactionError
	r.Message = err.Error()
	r.Data = data
}

func (r *ResBody) Success(data map[string]interface{})  {
	r.Code = Success
	r.Message = Message[Success]
	r.Data = data
}

func (r *ResBody) ReceiveFileError(err error,data map[string]interface{})  {
	r.Code = ReceiveFileError
	r.Message = err.Error()
	r.Data = data
}

func (r *ResBody) FileTypeError(err error, data map[string]interface{}) {
	r.Code = FileTypeError
	r.Message = err.Error()
	r.Data = data
}

func (r *ResBody) SaveFileError(err error, data map[string]interface{})  {
	r.Code = SaveFileError
	r.Message = err.Error()
	r.Data = data
}

func (r *ResBody) BindingValidatorError(err error, data map[string]interface{}) {
	r.Code = BindingValidatorError
	r.Message = err.Error()
	r.Data = data
}

func (r *ResBody) NumericStringError(err wdmError.NumericStringError) {
	r.Code = NumericStringError
	r.Message = Message[NumericStringError]
	r.Data = map[string]interface{}{
		"notNumericStringFields": err.NotNumericFields,
	}
}

func (r *ResBody) ParamValueError(err wdmError.ParamValueError) {
	r.Code = ParamValueError
	r.Message = Message[ParamValueError]
	r.Data = map[string]interface{}{
		"infos":err.Details,
	}
}

func (r *ResBody) ParamTypeError(err *wdmError.ParamTypeError) {
	r.Code = ParamTypeError
	r.Message = Message[ParamTypeError]
	r.Data = map[string]interface{}{
		"info":err.FormFieldName + " should be a " + err.ShouldType,
	}
}

func (r *ResBody) KindNotExistError(data map[string]interface{})  {
	r.Code = KindIsNotExist
	r.Message = Message[KindIsNotExist]
	r.Data = data
}

func (r *ResBody) CategoryHasExistedError(data map[string]interface{}) {
	r.Code = CategoryHasExisted
	r.Message = Message[CategoryHasExisted]
	r.Data = data
}