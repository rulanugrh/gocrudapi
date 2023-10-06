package entity

type ResponseBook struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	Price      int    `json:"price"`
	CategoryID uint   `json:"category_id"`
}

type ResponseCategory struct {
	Name        string `json:"name"`
	Description int    `json:"description"`
}

type ResponseAPI struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
