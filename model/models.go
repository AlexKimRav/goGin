package model

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"tittle"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}
