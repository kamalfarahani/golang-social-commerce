package viewmodels

type Collection struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	ImgUrl   string    `json:"imgUrl"`
	Products []Product `json:"products"`
}
