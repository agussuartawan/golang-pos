package request

type RoleRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=255"`
	Description *string `json:"description" form:"description"`
}

type AppendPermissionRequest struct {
	RoleId uint `json:"roleId" form:"roleId" validate:"required"`
	PermissionIds []uint `json:"permissionIds" form:"permissionIds" validate:"required,gt=0"`
}