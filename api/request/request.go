package request

// User represents the json object for creating a User
type User struct {
	Username   string     `json:"username" validate:"required"`
	Email      string     `json:"email" validate:"required"`
	Permission Permission `json:"permission" validate:"required"`
}

// Permission represents the json object for creating a Permission
type Permission struct {
	Name  string   `json:"name" validate:"required"`
	Roles []string `json:"roles" validate:"required"`
}

// LoginUser represents the json object for login a User
type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
