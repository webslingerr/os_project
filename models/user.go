package models

type UserType string

const (
	Client   UserType = "Client"
	Owner    UserType = "Owner"
	SysAdmin UserType = "Sys_admin"
)

type User struct {
	ID        string   `json:"id" db:"id"`
	Email     string   `json:"email" db:"email"`
	Fullname  string   `json:"fullname" db:"fullname"`
	Password  string   `json:"password" db:"password"`
	Address   string   `json:"address" db:"address"`
	Type      UserType `json:"type" db:"type"`
	CreatedAt string   `json:"created_at" db:"created_at"`
	UpdatedAt string   `json:"updated_at" db:"updated_at"`
}

type UserPrimaryKey struct {
	ID string `json:"id" db:"id"`
}

type CreateUser struct {
	Email    string   `json:"email" db:"email" validate:"required,email"`
	Fullname string   `json:"fullname" db:"fullname" validate:"required"`
	Password string   `json:"password" db:"password" validate:"required,min=8"`
	Address  string   `json:"address" db:"address" validate:"required"`
	Type     UserType `json:"type" db:"type" validate:"required,oneof=Client Owner Sys_admin"`
}

type UpdateUser struct {
	ID        string   `json:"id" db:"id" validate:"required,uuid"`
	Email     string   `json:"email" db:"email" validate:"email"`
	Fullname  string   `json:"fullname" db:"fullname"`
	Password  string   `json:"password" db:"password" validate:"min=8"`
	Address   string   `json:"address" db:"address"`
	Type      UserType `json:"type" db:"type" validate:"oneof=Client Owner Sys_admin"`
	UpdatedAt string   `json:"updated_at" db:"updated_at"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset" db:"offset"`
	Limit  int    `json:"limit" db:"limit"`
	Search string `json:"search" db:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
