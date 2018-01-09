package errors

import "fmt"

// NewModelsError returns a new specific Error for Models errors.
// It return only one InnerError in the `Errors` silce.
func NewModelsError(k string, metadata M) error {
	key := fmt.Sprintf("models-%s", k)
	return &Error{
		Status:     status(key),
		StatusText: statusText(key),
		Errors: []InnerError{{
			Code:     code(key),
			Kind:     "models",
			Metadata: appendReasonTo(key, metadata),
		}},
	}
}
