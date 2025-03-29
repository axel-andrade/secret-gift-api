package common_adapters

type OutputPort struct {
	StatusCode int
	Data       any
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorMessageWithDetails struct {
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details"`
}
