package model

import "github.com/javiorfo/go-microservice-lib/auditory"

// Dummy represents a dada structure
type Dummy struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Info string `json:"info"`
	auditory.Auditable
}

// User represents a dada structure
type User struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Permission Permission `json:"permission"`
	Password   string     `json:"-"`
	Salt       string     `json:"-"`
	Temporary  bool       `json:"-"`
	auditory.Auditable
}

// Permission represents a dada structure
type Permission struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Roles []Role `json:"roles"`
}

// Role represents a dada structure
type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}
