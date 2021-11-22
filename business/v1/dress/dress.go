package dress

// Dress 礼服类 即具体的每一件礼服
type Dress struct {
	Id int
	Category Category
	SerialNumber int
	RentCounter int
	LaundryCounter int
	MaintainCounter int
	CoverImg string
	SecondaryImg []string
	Status string
}