package do

import (
	"time"

	"github.com/kackerx/interview/internal/domain/enum"
)

type User struct {
	UserName  string          `json:"user_name"`
	Password  string          `json:"password"`
	Nickname  string          `json:"nickname"`
	Email     string          `json:"email"`
	Avatar    string          `json:"avatar"`
	Status    enum.UserStatus `json:"status"`
	Gender    enum.UserGender `json:"gender"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
