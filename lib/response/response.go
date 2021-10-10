package response

import (
	"github.com/gin-gonic/gin"
)

type ResBody struct {
	Code int
	// TODO: 此处数据校验有可能出现多种错误 但除数据校验之外 可能不会再有多种错误的情况
	// TODO: 为了满足数据校验错误 所以才把Message字段设计为[]string 有没有什么办法解决?
	Message string
	Data map[string]interface{}
}

// 非正常响应状态码定义规则:
// 100XX:校验类错误
// 101XX:数据库错误
// 102XX:字段错误
const (
	// Success 正常响应
	Success = 200

	// ConvertBindErrFailed 将binding的err转化为validator.ValidationErrors失败
	ConvertBindErrFailed = 10000

	// RequestContentTypeErr 请求类型错误
	RequestContentTypeErr = 10001

	// FieldNotInt 字段值非int
	FieldNotInt = 10001

	// DBError 数据库操作错误
	DBError = 10100

	// ParamsInvalid 表示参数校验错误
	ParamsInvalid = 10201

	// SerialNumberInvalid 无效的礼服编号
	SerialNumberInvalid = 10203

	// FileNumZero 上传文件的字段没有文件
	FileNumZero = 10204

	// UploadFileFailed 上传文件失败
	UploadFileFailed = 10205

	// FileIsNotImg 文件非图片
	FileIsNotImg = 10206

	// KindAndCodeInvalid 无效的品类名称和编码
	KindAndCodeInvalid = 10207
)

var Message = map[int]string {
	Success: "success",
	ConvertBindErrFailed: "convert binding err to validator err failed",
	SerialNumberInvalid: "invalid serial number",
	FieldNotInt: "field is not int",
	FileNumZero: "field is not have file",
	FileIsNotImg: "file is not img",
	KindAndCodeInvalid: "kind and code invalid",
}

//func (r *ResBody) ConvertBindErrFailed(data []interface{}) {
//	r.Code = ConvertBindErrFailed
//	r.Message = []string{Message[ConvertBindErrFailed]}
//	r.Data = data
//}

// TODO:此处从err中获取错误信息 这一步属于"封装信息"的功能 应该由ResBody来完成
//func (r *ResBody) ParamsInvalid(errsInfo []string, data []interface{}) {
//	r.Code = ParamsInvalid
//	r.Message = errsInfo
//	r.Data = data
//}

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

func SuccessResp(data []interface{}) gin.H {
	return gin.H {
		"code": Success,
		"message": Message[Success],
		"data": data,
	}
}

func RequestContentTypeErrResp(message string, data []interface{}) gin.H {
	return gin.H {
		"code": RequestContentTypeErr,
		"message": message,
		"data": data,
	}
}

func FieldNotIntResp(field string, data []interface{}) gin.H {
	return gin.H {
		"code": FieldNotInt,
		"message": field + " " + Message[FieldNotInt],
		"data": data,
	}
}

func DBErrorResp(err error, data []interface{}) gin.H {
	return gin.H {
		"code": DBError,
		"message": err.Error(),
		"data": data,
	}
}

func ParamsInvalidResp(errsInfo []string, data []interface{}) gin.H {
	return gin.H {
		"code":    ParamsInvalid,
		"message": errsInfo,
		"data":    data,
	}
}

func SerialNumberInvalidResp(data []interface{}) gin.H {
	return gin.H {
		"code":    SerialNumberInvalid,
		"message": Message[SerialNumberInvalid],
		"data":    data,
	}
}

func FileNumZeroResp(field string, data []interface{}) gin.H {
	return gin.H {
		"code": FileNumZero,
		"message": field + " " + Message[FileNumZero],
		"data": data,
	}
}

func UploadFileFailedResp(err error, data[]interface{}) gin.H {
	return gin.H {
		"code": UploadFileFailed,
		"message": err.Error(),
		"data": data,
	}
}

func FileIsNotImgResp(data []interface{}) gin.H {
	return gin.H {
		"code": FileIsNotImg,
		"message": Message[FileIsNotImg],
		"data": data,
	}
}

func KindAndCodeInvalidResp(data []interface{}) gin.H {
	return gin.H {
		"code": KindAndCodeInvalid,
		"message": Message[KindAndCodeInvalid],
		"data": data,
	}
}
