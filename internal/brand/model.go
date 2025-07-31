package brand

type Model struct {
	ID   string
	Name string
}

type Request struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
