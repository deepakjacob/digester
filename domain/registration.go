package domain

import "time"

// Registration captures a file registration.
// Please note that we are not capturing / persisting file contents

type FileIDType string

type Registration struct {
	FileID     FileIDType
	FileName   string
	FileDate   time.Time
	TowerID    string
	LocationID string
	PostalCode string
	AreaCode   string
}

// RegistrationStatus the status of registation
type RegistrationStatus struct {
	FileID     FileIDType `json:"file_id"`
	StatusCD   string     `json:"status_cd"`
	StatusDesc string     `json:"status_desc"`
}
