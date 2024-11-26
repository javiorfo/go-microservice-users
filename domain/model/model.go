package model

import (
	"github.com/javiorfo/go-microservice-lib/auditory"
	"github.com/javiorfo/go-microservice-users/security/pwd"
)

// User represents a dada structure
type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PermissionID uint       `json:"-"`
	Permission   Permission `json:"permission" gorm:"column:permission_id;not null"`
	Password     string     `json:"-"`
	Salt         string     `json:"-"`
	Temporary    bool       `json:"-"`
	auditory.Auditable
}

func (u User) VerifyPassword(password string) bool {
	hashedInputPassword := pwd.Hash(password, u.Salt)
	return hashedInputPassword == u.Password
}

// Permission represents a data structure
type Permission struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"unique"`
	Roles []Role `json:"roles" gorm:"many2many:permissions_roles;"`
}

// Role represents a data structure
type Role struct {
	ID          uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string       `json:"name" gorm:"unique"`
	Permissions []Permission `json:"permissions" gorm:"many2many:permissions_roles;"`
}

type PermissionRole struct {
	PermissionID uint `gorm:"primaryKey"`
	RoleID       uint `gorm:"primaryKey"`
}

func (PermissionRole) TableName() string {
	return "permissions_roles"
}
