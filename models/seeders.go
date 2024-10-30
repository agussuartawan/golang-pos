package models

import (
	"github.com/agussuartawan/golang-pos/core/config"
	helper "github.com/agussuartawan/golang-pos/core/helpers"
	"log"
)

func RunSeeders() {
	SeedUser()
}

func SeedUser() {
	users := []User{
		{Name: "Super Admin", Email: "superadmin@mail.com", Password: "superadmin"},
	}

	roles := []Role{
		{Name: "super_admin"},
		{Name: "admin"},
		{Name: "admin_warehouse"},
	}

	var superAdminUser User
	var superAdminRole []Role

	for _, role := range roles {
		if err := config.DB.FirstOrCreate(&role, Role{Name: role.Name}).Error; err != nil {
			helper.LogError(err)
		}
		if role.Name == "super_admin" {
			superAdminRole = append(superAdminRole, role)
		}
		log.Printf("Membuat role dengan nama: %v", role.Name)
	}

	for _, user := range users {
		err := config.DB.FirstOrCreate(&user, User{Email: user.Email}).Error
		if err != nil {
			helper.LogError(err)
		}
		if user.Name == "Super Admin" {
			superAdminUser = user
		}
		log.Printf("Membuat user dengan nama: %v", user.Name)
	}

	// assign super admin role
	if err := config.DB.Model(&superAdminUser).Association("Roles").Append(superAdminRole); err != nil {
		helper.LogError(err)
	}

}
