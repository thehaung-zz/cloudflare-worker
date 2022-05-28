package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type HttpUtil struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type MessageError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type HttpError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewHttpUtil(response http.ResponseWriter, request *http.Request) *HttpUtil {
	return &HttpUtil{
		Response: response,
		Request:  request,
	}
}

func (h *HttpUtil) WriteJson(status int, data interface{}) {
	h.Response.Header().Set("Content-Type", "application/json")
	h.Response.WriteHeader(status)

	res, err := json.Marshal(data)

	if err != nil {
		log.Println()
	}

	if _, err := h.Response.Write(res); err != nil {
		log.Println(err.Error())
	}
}

func (h *HttpUtil) WriteError(status int, message string) {
	h.Response.Header().Set("Content-Type", "application/json")
	h.Response.WriteHeader(status)
	data := &MessageError{
		StatusCode: status,
		Message:    message,
	}
	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	if _, err := h.Response.Write(res); err != nil {
		log.Println(err.Error())
	}
}

func (h *HttpUtil) WriteHttpError(httpError HttpError) {
	h.Response.Header().Set("Content-Type", "application/json")
	h.Response.WriteHeader(httpError.StatusCode)

	res, err := json.Marshal(httpError)

	if err != nil {
		log.Println()
	}

	if _, err := h.Response.Write(res); err != nil {
		log.Println(err.Error())
	}
}

func (h *HttpUtil) GetBody(data interface{}) {
	if err := json.NewDecoder(h.Request.Body).Decode(data); err != nil {
		log.Println(err.Error())
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	return json.NewDecoder(resp.Body).Decode(target)
}

func (h *HttpUtil) QueryParamsString(params ...string) (results map[string]string) {
	results = make(map[string]string)

	for _, param := range params {
		query := h.Request.URL.Query().Get(param)
		results[param] = query
	}

	return results
}

func (h *HttpUtil) QueryParamsInt(params ...string) (results map[string]int) {
	results = make(map[string]int)

	for _, param := range params {
		query := h.Request.URL.Query().Get(param)
		queryInt, err := strconv.Atoi(query)
		if err != nil {
			log.Println(err.Error())
		}

		results[param] = queryInt
	}

	return results
}
