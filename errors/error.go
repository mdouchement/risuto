package errors

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/mdouchement/risuto/util"
)

// HTTPCoder interface is implemented by application errors
type HTTPCoder interface {
	// HTTPCode return the HTTP status code for the given error
	HTTPCode() int
}

// M is metadata structure
type M map[string]interface{}

// Error describe all errors occured when a license is asked.
type Error struct {
	Status     int          `json:"status"`
	StatusText string       `json:"status_text"`
	Errors     []InnerError `json:"errors"`
}

type InnerErrors []InnerError

type InnerError struct {
	Code     string `json:"code"`
	Kind     string `json:"type"`
	Metadata M      `json:"metadata,omitempty"`
}

var codeList = map[string]map[string]string{
	"controllers-unexpected": {
		"code":        "RISUTO-500-000",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"controllers-merge_params": {
		"code":        "RISUTO-500-001",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"controllers-params_binding": {
		"code":        "RISUTO-422-000",
		"status":      "422",
		"status_text": "Unprocessable Entity",
		"reason":      "Validation of <no value> failed",
	},
	"models-not_found": {
		"code":        "RISUTO-404-000",
		"status":      "404",
		"status_text": "Not Found",
		"reason":      "Could not find {{.resource}} with id `{{.id}}`",
	},
	"models-unexpected": {
		"code":        "RISUTO-500-002",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
}

func StatusCode(err error) int {
	if hc, ok := err.(HTTPCoder); ok {
		return hc.HTTPCode()
	}
	return http.StatusInternalServerError
}

func (e *InnerError) Error() string {
	return fmt.Sprintf("%s-%s: %s", e.Kind, e.Code, e.Metadata["reason"])
}

// Error contacts all InnerError in a single string.
func (e *Error) Error() string {
	var errf bytes.Buffer
	errf.WriteString("[")
	for i, err := range e.Errors {
		errf.WriteString("\"")
		errf.WriteString(err.Error())
		if i < len(e.Errors)-1 {
			errf.WriteString("\",")
		} else {
			errf.WriteString("\"]")
		}
	}
	return errf.String()
}

func (e *Error) HTTPCode() int {
	return e.Status
}

func code(key string) string {
	return codeList[key]["code"]
}

func statusText(key string) string {
	return codeList[key]["status_text"]
}

func status(key string) int {
	return util.MustAtoi(codeList[key]["status"])
}

func appendReasonTo(key string, metadata M) M {
	if reasonTemplate, ok := codeList[key]["reason"]; ok {
		t := template.Must(template.New("reason").Parse(reasonTemplate))
		reason := &bytes.Buffer{}
		t.Execute(reason, metadata)
		metadata["reason"] = reason.String()
	}
	return metadata
}
