package response

import "time"

type ErrorDto struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Field     any       `json:"field,omitempty"`
}
