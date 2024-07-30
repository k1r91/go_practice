package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// начало решения

// Handy предоставляет удобный интерфейс
// для выполнения HTTP-запросов
type Handy struct{
	client *http.Client
	params url.Values
	headers map[string]string
	formData url.Values
	jsonData []byte
	url string
	err error
}

// NewHandy создает новый экземпляр Handy
func NewHandy() *Handy {
	client := http.Client{}
	return &Handy{&client, url.Values{}, make(map[string]string), url.Values{}, make([]byte, 0), "", nil}
}

// URL устанавливает URL, на который пойдет запрос
func (h *Handy) URL(uri string) *Handy {
	h.url = uri
	return h
}

// Client устанавливает HTTP-клиента
// вместо умолчательного http.DefaultClient
func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

// Header устанавливает значение заголовка
func (h *Handy) Header(key, value string) *Handy {
	h.headers[key] = value
	return h
}

// Param устанавливает значение URL-параметра
func (h *Handy) Param(key, value string) *Handy {
	h.params.Add(key, value)
	return h
}

// Form устанавливает данные, которые будут закодированы
// как application/x-www-form-urlencoded и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) Form(form map[string]string) *Handy {
	h.jsonData = make([]byte, 0)
	for key, value := range form {
		h.formData.Add(key, value)
	}
	return h
}

// JSON устанавливает данные, которые будут закодированы
// как application/json и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) JSON(v any) *Handy {
	var err error
	h.formData = url.Values{}
	h.jsonData, err = json.Marshal(v)
	if err != nil {
		h.err = err
	}
	return h
}


// Get выполняет GET-запрос с настроенными ранее параметрами
func (h *Handy) Get() *HandyResponse {
	if h.err != nil {
		return &HandyResponse{nil, h.err, 0}
	}
	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		return &HandyResponse{nil, err, 0}
	}
	if len(h.params) > 0 {
		req.URL.RawQuery = h.params.Encode()
	}
	for k, v := range h.headers {
		req.Header.Add(k, v)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return &HandyResponse{nil, err, 0}
	}
	defer resp.Body.Close()
	return &HandyResponse{resp, nil, resp.StatusCode}
}

// Post выполняет POST-запрос с настроенными ранее параметрами
func (h *Handy) Post() *HandyResponse {
	var resp *http.Response
	var err error
	if h.err != nil {
		return &HandyResponse{nil, h.err, 0}
	}
	if len(h.formData) > 0 {
		resp, err = h.client.PostForm(h.url, h.formData)
		if err != nil {
			return &HandyResponse{nil, err, 0}
		}
	} else if len(h.jsonData) > 0 {
		req, err := http.NewRequest(http.MethodPost, h.url, bytes.NewReader(h.jsonData))
		if err != nil {
			return &HandyResponse{nil, err, 0}
		}
		for k, v := range h.headers {
			req.Header.Add(k, v)
		}
		req.Header.Add("Content-Type", "application/json")  // (2)
		req.Header.Add("Accept", "application/json")
		resp, err = h.client.Do(req)
		if err != nil {
			return &HandyResponse{resp, nil, resp.StatusCode}
		}
	} else {
		req, err := http.NewRequest(http.MethodPost, h.url, nil)
		if err != nil {
			return &HandyResponse{resp, nil, resp.StatusCode}
		}
		for k, v := range h.headers {
			req.Header.Add(k, v)
		}
		if len(h.params) > 0 {
			req.URL.RawQuery = h.params.Encode()
		}
		resp, err = h.client.Do(req)
		if err != nil {
			return &HandyResponse{resp, nil, resp.StatusCode}
		}

	}
	defer resp.Body.Close()
	return &HandyResponse{resp, nil, resp.StatusCode}
}

// HandyResponse представляет ответ на HTTP-запрос
type HandyResponse struct {
	response *http.Response
	err error
	StatusCode int
	// ...
}

// OK возвращает true, если во время выполнения запроса
// не произошло ошибок, а код HTTP-статуса ответа равен 200
func (r *HandyResponse) OK() bool {
	if r.err == nil && r.StatusCode == 200 {
		return true
	}
	return false
}

// Bytes возвращает тело ответа как срез байт
func (r *HandyResponse) Bytes() []byte {
	data, err := io.ReadAll(r.response.Body)
	if err != nil {
		r.err = err
	}
	return data
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	return string(r.Bytes())
}

// JSON декодирует тело ответа из JSON и сохраняет
// результат по адресу, на который указывает v
func (r *HandyResponse) JSON(v any) {
	err := json.Unmarshal(r.Bytes(), &v)
	if err != nil {
		r.err = err
	}
	// работает аналогично json.Unmarshal()
	// если при декодировании произошла ошибка,
	// она должна быть доступна через r.Err()
}

// Err возвращает ошибку, которая возникла при выполнении запроса
// или обработке ответа
func (r *HandyResponse) Err() error {
	return r.err
}

// конец решения

func main() {
	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}