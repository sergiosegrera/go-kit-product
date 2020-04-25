package models

type Thumbnail struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Price     int64  `json:"price"`
}

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Thumbnail   string    `json:"thumbnail"`
	Images      []string  `json:"images"`
	Description string    `json:"description"`
	Options     []*Option `json:"options"`
	Price       int64     `json:"price"`
	Public      bool      `json:"public"`
}

type Option struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
}
