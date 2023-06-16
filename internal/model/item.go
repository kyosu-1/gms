package model

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	ID              uuid.UUID
	Name            string
	Description     string
	AcquisitionDate time.Time
	Location        string
	Categories      []Category
}
