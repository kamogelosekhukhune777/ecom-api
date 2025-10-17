package model

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	ModifiedAt time.Time
}
