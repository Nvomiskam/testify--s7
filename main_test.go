package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Проверяет, что запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerCheckStatusandBody(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	body := responseRecorder.Body.String()

	require.Equal(t, status, http.StatusOK)
	require.NotEmpty(t, body)

}

// Проверяет, что город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenCityIsNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=invalidcity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	body := responseRecorder.Body.String()

	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, "wrong city value", body)
}

// Проверяет, что в случае если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)

}
