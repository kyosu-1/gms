package model

import "time"

type Item struct {
	ID              int
	Name            string
	Description     string
	AcquisitionDate time.Time
	Location        string
	Categories      []Category
}
