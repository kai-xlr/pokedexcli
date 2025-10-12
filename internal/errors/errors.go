package errors

import "fmt"

// CommandError represents errors that occur during command execution
type CommandError struct {
	Command string
	Message string
	Err     error
}

func (e *CommandError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s command error: %s: %v", e.Command, e.Message, e.Err)
	}
	return fmt.Sprintf("%s command error: %s", e.Command, e.Message)
}

func (e *CommandError) Unwrap() error {
	return e.Err
}

// NewCommandError creates a new command error
func NewCommandError(command, message string, err error) *CommandError {
	return &CommandError{
		Command: command,
		Message: message,
		Err:     err,
	}
}

// ValidationError represents input validation errors
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for %s: %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// APIError represents errors from the Pokemon API
type APIError struct {
	Operation  string
	StatusCode int
	Message    string
	Err        error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("API error during %s (status %d): %s: %v", e.Operation, e.StatusCode, e.Message, e.Err)
	}
	return fmt.Sprintf("API error during %s (status %d): %s", e.Operation, e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new API error
func NewAPIError(operation string, statusCode int, message string, err error) *APIError {
	return &APIError{
		Operation:  operation,
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}
