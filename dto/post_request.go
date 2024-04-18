package dto

type PostRequest struct {
	Content  string `json:"content"`
	ImageUrl string `json:"image_url" validate:"required,url"`
}
