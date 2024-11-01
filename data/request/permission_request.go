package request

type PermissionRequest struct {
	Name        string  `json:"name" form:"name" validate:"required,max=255"`
	Description *string `json:"description" form:"description"`
}

type PermissionParam struct {
	Name *string `form:"name"`
	PaginationParam
}
