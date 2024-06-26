package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/greeflas/racing-engine-backend/pkg/apperror"
)

type APIError struct {
	Error string `json:"error"`
}

func NewAPIError(err error) *APIError {
	return &APIError{
		Error: err.Error(),
	}
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationError(field string, message string) *ValidationError {
	return &ValidationError{Field: field, Message: message}
}

func HandleError(w http.ResponseWriter, err error) {
	log.Println(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	var entityNotFoundError *apperror.EntityNotFoundError

	if errors.As(err, &entityNotFoundError) {
		apiError := NewAPIError(err)

		if err := json.NewEncoder(w).Encode(apiError); err != nil {
			log.Println(err)
		}

		return
	}

	var validationErrors []*ValidationError

	for _, validationErr := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, NewValidationError(
			validationErr.Field(),
			validationErr.Error(),
		))
	}

	if err := json.NewEncoder(w).Encode(validationErrors); err != nil {
		log.Println(err)
	}
}
