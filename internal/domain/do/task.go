package do

import (
	"time"

	"github.com/kackerx/interview/internal/domain/enum"
)

type Task struct {
	ID     uint
	Status enum.TaskStatus

	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
