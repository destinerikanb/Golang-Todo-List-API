package web

type ActivityUpdateRequest struct {
	ID    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}
