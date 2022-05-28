package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/thehaung/cloudflare-worker/pkg/common/constant"
	"net/http"
)

type ErrorDetail struct {
	Namespace       string      `json:"namespace,omitempty"`
	Field           string      `json:"field,omitempty"`
	StructNamespace string      `json:"struct_namespace,omitempty"`
	StructField     string      `json:"struct_field,omitempty"`
	Tag             string      `json:"tag,omitempty"`
	ActualTag       string      `json:"actual_tag,omitempty"`
	Kind            string      `json:"kind,omitempty"`
	Type            string      `json:"type,omitempty"`
	Param           string      `json:"param,omitempty"`
	Value           interface{} `json:"value,omitempty"`
}

type MessageValidateError struct {
	StatusCode  int            `json:"statusCode,omitempty"`
	Message     string         `json:"message,omitempty"`
	ErrorDetail *[]ErrorDetail `json:"errorDetail,omitempty"`
}

func HandleErrorInValid(err error) *MessageValidateError {
	result := make([]ErrorDetail, 0)
	errors := err.(validator.ValidationErrors)

	for _, e := range errors {
		result = append(result, ErrorDetail{
			Namespace:       e.Namespace(),
			Field:           e.Field(),
			StructNamespace: e.StructNamespace(),
			StructField:     e.StructField(),
			Tag:             e.Tag(),
			ActualTag:       e.ActualTag(),
			Kind:            e.Kind().String(),
			Type:            e.Type().Name(),
			Value:           e.Value(),
		})
	}

	return &MessageValidateError{
		Message:     constant.MessageBadRequest,
		StatusCode:  http.StatusBadRequest,
		ErrorDetail: &result,
	}
}
