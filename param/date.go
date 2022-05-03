package param

import (
	"time"
)

const DateFormat = "2006-01-02"

// Date 接收请求和发送响应时 对yyyy-mm-dd格式的时间进行转化用
// 接收请求时 将一个"yyyy-mm-dd"格式的时间转化为一个time.Time对象
// 发送响应时 将一个time.Time对象转化为"yyyy-mm-dd"的时间
type Date time.Time

// UnmarshalJSON gin框架的c.ShouldBindJSON会调用field.UnmarshalJSON
// 故该方法在接收请求时会被调用 其功能为 将一个"yyyy-mm-dd"格式的时间转化为一个time.Time对象
func (d *Date) UnmarshalJSON(data []byte) (err error) {
	// 空值则不解析
	if len(data) == 2 {
		*d = Date(time.Time{})
		return nil
	}

	// 指定解析格式
	paramDate, err := time.Parse(`"`+DateFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*d = Date(paramDate)
	return nil
}

// MarshalJSON 在c.JSON时会被调用
// 故该方法在发送响应时会被调用 其功能为 将一个time.Time对象转化为"yyyy-mm-dd"的时间
func (d *Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateFormat)+2)
	b = time.Time(*d).AppendFormat(b, DateFormat)
	return b, nil
}

func (d *Date) IsBefore(target time.Time) bool {
	return time.Time(*d).Before(target)
}
