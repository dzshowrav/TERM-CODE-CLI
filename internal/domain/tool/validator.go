package tool

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ValidationResult struct {
	Valid  bool
	Errors []ValidationError
}

func (r ValidationResult) Error() string {
	var parts []string
	for _, e := range r.Errors {
		parts = append(parts, e.Error())
	}
	return strings.Join(parts, "; ")
}

func ValidateArgs(schema any, args map[string]any) ValidationResult {
	schemaMap, ok := schema.(map[string]any)
	if !ok {
		return ValidationResult{Valid: true}
	}

	var errors []ValidationError

	props, _ := schemaMap["properties"].(map[string]any)
	required, _ := schemaMap["required"].([]any)

	// Check required fields
	requiredSet := make(map[string]bool)
	for _, r := range required {
		reqName := fmt.Sprintf("%v", r)
		requiredSet[reqName] = true
		if _, exists := args[reqName]; !exists {
			errors = append(errors, ValidationError{
				Field:   reqName,
				Message: "required field missing",
			})
		}
	}

	// Validate types for provided args
	for key, val := range args {
		propDef, ok := props[key].(map[string]any)
		if !ok {
			continue
		}
		propType, _ := propDef["type"].(string)
		if err := validateType(key, propType, val); err != nil {
			errors = append(errors, *err)
		}
	}

	if len(errors) > 0 {
		return ValidationResult{Valid: false, Errors: errors}
	}
	return ValidationResult{Valid: true}
}

func validateType(field, expectedType string, val any) *ValidationError {
	if val == nil {
		return nil
	}

	switch expectedType {
	case "string":
		if _, ok := val.(string); !ok {
			return &ValidationError{Field: field, Message: "expected string"}
		}
	case "integer":
		switch val.(type) {
		case int, int32, int64, float64:
			return nil
		default:
			return &ValidationError{Field: field, Message: "expected integer"}
		}
	case "number":
		switch val.(type) {
		case int, int32, int64, float64, float32:
			return nil
		default:
			return &ValidationError{Field: field, Message: "expected number"}
		}
	case "boolean":
		if _, ok := val.(bool); !ok {
			return &ValidationError{Field: field, Message: "expected boolean"}
		}
	case "array":
		if _, ok := val.([]any); !ok {
			return &ValidationError{Field: field, Message: "expected array"}
		}
	case "object":
		if _, ok := val.(map[string]any); !ok {
			return &ValidationError{Field: field, Message: "expected object"}
		}
	}

	if enumVals, ok := getEnums(expectedType); ok {
		strVal := fmt.Sprintf("%v", val)
		found := false
		for _, e := range enumVals {
			if strVal == fmt.Sprintf("%v", e) {
				found = true
				break
			}
		}
		if !found {
			return &ValidationError{Field: field, Message: fmt.Sprintf("must be one of: %v", enumVals)}
		}
	}

	return nil
}

func getEnums(typeDef string) ([]any, bool) {
	return nil, false
}
