package middlewares

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/mdouchement/risuto/errors"
)

// HTTPErrorHandler is a middleware that formats rendered errors.
func HTTPErrorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		c.Set("error", err.Error())

		switch err := err.(type) {
		case *echo.HTTPError:
			c.HTML(err.Code, fmt.Sprintf("%s", err.Message))
		case *errors.Error:
			c.JSON(errors.StatusCode(err), err)
		default:
			err = errors.NewControllersError("unexpected", errors.M{
				"reason":  "Calm down, yo.",
				"details": err.Error(),
			})
			c.JSON(errors.StatusCode(err), err)
		}
	}
}
