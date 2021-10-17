package unit

type Unit struct {
	// 礼服ID
	Id int	`json:"id"`

	// 礼服品类ID
	CategoryId int `json:"categoryId"`

	// 礼服序号
	SerialNumber int `json:"serialNumber"`

	// 尺码
	Size string `json:"size"`

	// 出租次数
	RentNumber int `json:"rentNumber"`

	// 送洗次数
	LaundryNumber int `json:"laundryNumber"`

	// 封面图
	CoverImg string `json:"coverImg"`

	// 副图
	SecondaryImg []string `json:"secondaryImg,omitempty"`

	// 状态
	Status string `json:"status,omitempty"`
}