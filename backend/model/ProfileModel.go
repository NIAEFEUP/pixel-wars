package model

// Profile struct keeps the user's profile
type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}
