package controllers

import (
	"encoding/json"

	"github.com/mdouchement/risuto/errors"
	"github.com/mdouchement/risuto/util"
	govalidator "gopkg.in/asaskevich/govalidator.v5"
)

type (
	ParamsValidator struct{}
	// A ParamsField can store details of a field from HTTP params.
	ParamsField map[string]interface{}
	resource    interface {
		RName() string
	}
)

func (pv *ParamsValidator) Validate(i interface{}) error {
	if ok, err := govalidator.ValidateStruct(i); !ok {
		name := "unknown"
		if r, ok := i.(resource); ok {
			name = r.RName()
		}
		return errors.NewControllersError("params_binding", errors.M{
			"resource": name,
			"fields":   paramsErrorsFormatter(err),
		})
	}

	return nil
}

func paramsErrorsFormatter(err error) []ParamsField {
	currentFields := make([]ParamsField, 0)

	switch v := err.(type) {
	case govalidator.Errors:
		// Get nested fields errors
		for _, err2 := range v {
			// There is only one error until no embedded fields
			fields := paramsErrorsFormatter(err2)
			currentFields = append(currentFields, fields...)
		}
	case govalidator.Error:
		return []ParamsField{{
			"parameter": util.ToSnake(v.Name),
			"type":      v.Err.Error(),
		}}

	}

	return currentFields
}

// MergeParams adds params to the resource.
func MergeParams(resource, params interface{}) error {
	raw, err := json.Marshal(params)
	if err != nil {
		return errors.NewControllersError("merge_params", errors.M{
			"reason": err.Error(),
		})
	}

	if err := json.Unmarshal(raw, resource); err != nil {
		return errors.NewControllersError("merge_params", errors.M{
			"reason": err.Error(),
		})
	}

	return nil
}
