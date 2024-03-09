package web

type TodoCreateRequest struct {
	ActivityGroupId int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Priority        string `json:"priority" validate:"required"`
}
