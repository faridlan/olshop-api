package productweb

type ProductRequest struct {
	Id         string
	Name       string
	Price      int
	Quantity   int
	CategoryId string
	Desription string
	ImageUrl   string
	CreatedAt  int64
	UpdatedAt  int64
}
