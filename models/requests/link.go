package requests

type LinkStore struct {
	Url string `validate:"required" json:"url"`
}

type LinkDelete struct {
	ID int `validate:"required" json:"id"`
}
