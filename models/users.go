package models

import "time"

type (
	//User represents a Single User in our Collection
	User struct {
		ID          string    `json:"user_id"`
		FirtName    string    `json:"first_name"`
		MiddleName  string    `json:"middle_name"`
		LastName    string    `json:"last_name"`
		Username    string    `json:"username"`
		Password    string    `json:"password"`
		Email       string    `json:"email"`
		IsAdmin     bool      `json:"is_admin"`
		DateCreated time.Time `json:"date_created"`
		DateUpdated time.Time `json:"date_updated"`
	}
)
