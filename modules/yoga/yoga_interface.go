package yoga

// --------- Response -----------

type GetResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

// ---------- DTO ---------------

type YogaSorts struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
