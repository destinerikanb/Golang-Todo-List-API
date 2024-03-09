package web

type ActivityCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required"`
}
