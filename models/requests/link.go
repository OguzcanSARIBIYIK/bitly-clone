package requests

type LinkStore struct {
	Url string `validate:"required" json:"url"`
}

type LinkDelete struct {
	ID uint `validate:"required" json:"id"`
}
