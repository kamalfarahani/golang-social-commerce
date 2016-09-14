package viewmodels

type Category struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	ImgUrl   string    `json:"imgUrl"`
	Products []Product `json:"products"`
}
