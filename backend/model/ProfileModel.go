package model

// Profile struct keeps the user's profile
type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

// Client represents the full internal description of a client (including remainingpixels and timestamps)
type Client struct {
	Profile         *Profile `json:"profile"`
	LastTimestamp   uint64   `json:"lastTimeStamp"`
	RemainingPixels uint64   `json:"remainingPixels"`
}
