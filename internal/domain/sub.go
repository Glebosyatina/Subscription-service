package domain

type Sub struct {
	Id          uint64 `json:"id"`
	UserId      uint64 `json:"user_id"`
	NameService string `json:"service_name"`
	Price       uint64 `json:"price"`
	Start       string `json:"start_date"`
	End         string `json:"end_date"`
}
