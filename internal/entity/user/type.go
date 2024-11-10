package user

import "time"

const (
	UserTable = "ms_user"
)

type Pagination struct {
	Page      int64 `json:"page,omitempty"`
	PageSize  int64 `json:"page_size,omitempty" `
	TotalData int64 `json:"total_data,omitempty"`
	TotalPage int64 `json:"total_page,omitempty"`
}

type User struct {
	Name      string    `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type SearchUserResponse struct {
	Users      []User     `json:"users,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}
