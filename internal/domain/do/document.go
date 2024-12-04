package do

import (
	"time"
)

type Document struct {
	ID     uint
	TaskID uint

	Content     string
	I18nContent string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
