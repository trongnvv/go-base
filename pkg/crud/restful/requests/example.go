package requests

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type EditRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type DeleteRequest struct {
	Id string `json:"id" validate:"required"`
}
