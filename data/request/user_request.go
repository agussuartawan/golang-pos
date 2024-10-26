package request

type UserRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=255"`
	Email string `json:"email" form:"email" validate:"required,max=255,email"`
	Phone *string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password" validate:"required,max=255"`
}

type AppendRoleRequest struct {
	UserId uint `json:"userId" form:"userId" validate:"required"`
	RoleIds []uint `json:"roleIds" form:"roleIds" validate:"required,gt=0"`
}