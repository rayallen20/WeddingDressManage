package response

type ResBody struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}

// 非正常响应状态码定义规则:
// 100XX:校验参数错误
// 101XX:数据库错误
// 102XX:接收文件错误
const (
	// Success 正常响应
	Success = 200

	// DBError 数据库操作错误
	DBError = 10100

	// ReceiveFileError 接收文件错误
	ReceiveFileError = 10200

	// SaveFileError 保存文件错误
	SaveFileError = 10201

	// FileTypeError 文件类型错误
	FileTypeError = 10001


)

var Message = map[int]string {
	Success: "success",
	FileTypeError: "file type must be jpg, jpeg or png",
}

func (r *ResBody) DBError(err error, data map[string]interface{}) {
	r.Code = DBError
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

func (r ResBody) SaveFileError(err error, data map[string]interface{})  {
	r.Code = SaveFileError
	r.Message = err.Error()
	r.Data = data
}