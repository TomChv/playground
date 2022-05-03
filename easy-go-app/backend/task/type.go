package task

type CreateData struct {
	Title       string
	Description string
}

type UpdateData struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
