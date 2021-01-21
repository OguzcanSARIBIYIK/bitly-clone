package requests

type LinkStore struct {
	Url string `validate:"required" json:"url"`
}
