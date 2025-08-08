package structs

type ErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"` // Gunakan map untuk menyimpan pesan error spesifik
}
