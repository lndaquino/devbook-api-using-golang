package models

// Password type models the password change DTO
type Password struct {
	New     string `json:"newPassword"`
	Current string `json:"password"`
}
