package viewmodels

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgUrl      string `json:"imgUrl"`
	Price       int    `json:"price"`
	Likes       uint   `json:"likes"`
}
