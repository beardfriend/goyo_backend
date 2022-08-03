package yoga

type YogaDAO struct {
	Id   uint
	Name string
}

type Response struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
