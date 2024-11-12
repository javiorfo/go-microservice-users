package model

import "github.com/javiorfo/go-microservice-lib/auditory"

// User represents a dada structure
type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
	Salt           string `json:"-"`
	auditory.Auditable
}
