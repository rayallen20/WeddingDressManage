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
	// Success 成功响应
	Success = 200

	// InvalidUnmarshalError JSON反序列化错误
	InvalidUnmarshalError = 501

	// ReceiveFileError 接收文件错误
	ReceiveFileError = 502

	// SaveFileError 保存文件错误
	SaveFileError = 503

	// DbError 数据库错误
	DbError = 502

	// FieldTypeError 字段类型错误
	FieldTypeError = 10201

	// ValidateError 参数校验错误
	ValidateError = 10202

	// KindNotExist 大类信息不存在
	KindNotExist = 10301

	// CategoryHasExist 品类信息已存在
	CategoryHasExist = 10302

	// CategoryNotExist 品类信息不存在
	CategoryNotExist = 10303
)

var message = map[int]string{
	ValidateError: "validate error",
	Success: "success",
}

func (r *RespBody) Success(data map[string]interface{}) {
	r.Code = Success
	r.Message = message[Success]
	r.Data = data
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

func (r *RespBody) DbError(err *sysError.DbError)  {
	r.Code = DbError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) KindNotExistError(err *sysError.KindNotExistError) {
	r.Code = KindNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CategoryHasExistError(err *sysError.CategoryHasExistError)  {
	r.Code = CategoryHasExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CategoryNotExistError(err *sysError.CategoryNotExistError)  {
	r.Code = CategoryNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) ReceiveFileError(err *sysError.ReceiveFileError)  {
	r.Code = ReceiveFileError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) SaveFileError(err *sysError.SaveFileError)  {
	r.Code = SaveFileError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}