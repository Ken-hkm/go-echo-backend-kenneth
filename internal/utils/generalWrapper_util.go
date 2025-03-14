package utils

import (
	"time"
)

type Response[T any] struct {
	Status     string    `json:"status"`
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Data       T         `json:"data"`
	Timestamp  time.Time `json:"timestamp"`
}
