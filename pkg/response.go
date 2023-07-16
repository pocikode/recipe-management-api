package pkg

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	customError "recipe-management/pkg/error"
)

type ResponseError struct {
	Message string `json:"message"`
}

func NewHttpErrorHandler(err error, c echo.Context) {
	//log.Println(reflect.TypeOf(err))
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if ce, ok := err.(*customError.Error); ok {
		he.Code = ce.Code()
		he.Message = ce.Error()
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		he.Code = http.StatusBadRequest
		switch ve[0].Tag() {
		case "required":
			he.Message = fmt.Sprintf("%s is required", ve[0].Field())
		default:
			he.Message = ve.Error()
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = map[string]interface{}{"message": message}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
