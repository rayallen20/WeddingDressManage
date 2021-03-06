package response

import (
	"WeddingDressManage/lib/sysError"
	"time"
)

type RespBody struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
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
	DbError = 504

	// FieldTypeError 字段类型错误
	FieldTypeError = 10201

	// ValidateError 参数校验错误
	ValidateError = 10202

	// NilJsonError 请求参数中的JSON为空错误
	NilJsonError = 10203

	// TimeParseError 无法将请求参数中的字符串解析为时间错误
	TimeParseError = 10204

	// KindNotExist 大类信息不存在
	KindNotExist = 10301

	// CategoryHasExist 品类信息已存在
	CategoryHasExist = 10302

	// CategoryNotExist 品类信息不存在
	CategoryNotExist = 10303

	// DressHasGifted 礼服状态已经为赠与
	DressHasGifted = 10304

	// DressHasDiscarded 礼服状态已经为销库
	DressHasDiscarded = 10305

	// DressNotExist 礼服不存在
	DressNotExist = 10306

	// CustomerNotExist 客户不存在
	CustomerNotExist = 10307

	// LaundryStatusError 礼服状态不符合送洗条件错误
	LaundryStatusError = 10308

	// MaintainStatusError 礼服状态不符合维护条件错误
	MaintainStatusError = 10309

	// LaundryRecordNotExist 送洗记录不存在错误
	LaundryRecordNotExist = 10310

	// DressIsNotLaundering 礼服不处于送洗状态错误
	DressIsNotLaundering = 10311

	// MaintainRecordNotExist 维护记录不存在错误
	MaintainRecordNotExist = 10312

	// DressIsNotMaintaining 礼服不处于维护状态错误
	DressIsNotMaintaining = 10313

	// WeddingDateBeforeToday 创建订单操作中搜索礼服步骤时 选择的婚期早于当天错误
	WeddingDateBeforeToday = 10314

	// CustomerBeBanned 客户被封禁错误
	CustomerBeBanned = 10315

	// DiscountInvalid 折扣无效错误
	DiscountInvalid = 10316

	// StrategyNotExist 订单优惠策略不存在
	StrategyNotExist = 10317

	// CustomPriceTooFew 自定义租金过低
	CustomPriceTooFew = 10318

	// DeliveryOrderNotExist 待出件订单不存在
	DeliveryOrderNotExist = 10319

	// 20XXX 前端需要渲染的
)

var message = map[int]string{
	ValidateError: "validate error",
	Success:       "success",
}

func (r *RespBody) Success(data map[string]interface{}) {
	r.Code = Success
	r.Message = message[Success]
	r.Data = data
}

func (r *RespBody) RequestNilJsonError(err *sysError.RequestNilJsonError) {
	r.Code = NilJsonError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) TimeParseError(err *time.ParseError) {
	r.Code = TimeParseError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) InvalidUnmarshalError(err *sysError.InvalidUnmarshalError) {
	r.Code = InvalidUnmarshalError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) FieldTypeError(err *sysError.UnmarshalTypeError) {
	r.Code = FieldTypeError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) ValidateError(errs []*sysError.ValidateError) {
	validateFailInfos := make([]string, 0, len(errs))
	for _, validateError := range errs {
		validateFailInfo := validateError.Error()
		validateFailInfos = append(validateFailInfos, validateFailInfo)
	}
	r.Code = ValidateError
	r.Message = message[ValidateError]
	r.Data = map[string]interface{}{
		"validateFailInfos": validateFailInfos,
	}
}

func (r *RespBody) DbError(err *sysError.DbError) {
	r.Code = DbError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) KindNotExistError(err *sysError.KindNotExistError) {
	r.Code = KindNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CategoryHasExistError(err *sysError.CategoryHasExistError) {
	r.Code = CategoryHasExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CategoryNotExistError(err *sysError.CategoryNotExistError) {
	r.Code = CategoryNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) ReceiveFileError(err *sysError.ReceiveFileError) {
	r.Code = ReceiveFileError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) SaveFileError(err *sysError.SaveFileError) {
	r.Code = SaveFileError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DressHasGiftedError(err *sysError.DressHasGiftedError) {
	r.Code = DressHasGifted
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DressHasDiscardedError(err *sysError.DressHasDiscardedError) {
	r.Code = DressHasDiscarded
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DressNotExistError(err *sysError.DressNotExistError) {
	r.Code = DressNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CustomerNotExistError(err *sysError.CustomerNotExistError) {
	r.Code = CustomerNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) LaundryStatusError(err *sysError.LaundryStatusError) {
	r.Code = LaundryStatusError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) MaintainStatusError(err *sysError.MaintainStatusError) {
	r.Code = MaintainStatusError
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) LaundryRecordNotExistError(err *sysError.LaundryRecordNotExistError) {
	r.Code = LaundryRecordNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DressIsNotLaunderingError(err *sysError.DressIsNotLaunderingError) {
	r.Code = DressIsNotLaundering
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) MaintainRecordNotExistError(err *sysError.MaintainRecordNotExistError) {
	r.Code = MaintainRecordNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DressIsNotMaintainingError(err *sysError.DressIsNotMaintainingError) {
	r.Code = DressIsNotMaintaining
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) WeddingDateBeforeTodayError(err *sysError.DateBeforeTodayError) {
	r.Code = WeddingDateBeforeToday
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CustomerBeBannedError(err *sysError.CustomerBeBannedError) {
	r.Code = CustomerBeBanned
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DiscountInvalidError(err *sysError.DiscountInvalidError) {
	r.Code = DiscountInvalid
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) StrategyNotExistError(err *sysError.StrategyNotExistError) {
	r.Code = StrategyNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) CustomPriceTooFewError(err *sysError.CustomPriceTooFewError) {
	r.Code = CustomPriceTooFew
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}

func (r *RespBody) DeliveryOrderNotExistError(err *sysError.DeliveryOrderNotExist) {
	r.Code = DeliveryOrderNotExist
	r.Message = err.Error()
	r.Data = map[string]interface{}{}
}
