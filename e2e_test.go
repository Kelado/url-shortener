package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	controller "github.com/Kelado/url-shortener/controllers"
	handler "github.com/Kelado/url-shortener/handlers"
	"github.com/Kelado/url-shortener/repositories"

	"github.com/stretchr/testify/assert"
)

var (
	testCodeSize = 6
	testHostname = "http://localhost:8000/"
)

var (
	shortURLpattern = `^baseURL/codeRunes{codeSize}$`
)

func initEnv() {
	shortURLpattern = fmt.Sprintf("^%s/[a-zA-Z]{%d}$", testHostname, testCodeSize)
}

func initHandler() *handler.Handler {
	linkRepo := repositories.NewMockDB()
	controller := controller.NewController(testHostname, testCodeSize, linkRepo)
	return handler.NewHandler(controller)
}

func TestPostReturnShortURL(t *testing.T) {
	initEnv()
	h := initHandler()
	expected := fmt.Sprintf("^%s[a-zA-Z]{%d}$", testHostname, testCodeSize)

	data := map[string]interface{}{
		"url": "https://example.com",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(data)

	req := httptest.NewRequest(http.MethodGet, "/shorten", reqBody)
	w := httptest.NewRecorder()
	h.HandlePostURL(w, req)
	res := w.Result()
	defer res.Body.Close()

	respData := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&respData)

	assert.NotNil(t, respData["shortUrl"], "shortUrl field not found")

	shortURL := respData["shortUrl"].(string)
	re := regexp.MustCompile(expected)
	errMesg := fmt.Sprintf("expected: %v, got: %v", expected, shortURL)
	assert.Equal(t, true, re.MatchString(shortURL), errMesg)
}

func TestPostWrongFieldName(t *testing.T) {
	initEnv()
	h := initHandler()

	data := map[string]interface{}{
		"wrongFieldName": "https://example.com",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(data)

	req := httptest.NewRequest(http.MethodGet, "/shorten", reqBody)
	w := httptest.NewRecorder()
	h.HandlePostURL(w, req)
	res := w.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	msg := string(bodyBytes)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assert.Equal(t, "missing required field", msg)
}

func TestPostEmptyURL(t *testing.T) {
	initEnv()
	h := initHandler()

	data := map[string]interface{}{
		"url": "",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(data)

	req := httptest.NewRequest(http.MethodGet, "/shorten", reqBody)
	w := httptest.NewRecorder()
	h.HandlePostURL(w, req)
	res := w.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	msg := string(bodyBytes)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assert.Equal(t, "empty url", msg)
}

func TestPostWrongURL(t *testing.T) {
	initEnv()
	h := initHandler()

	data := map[string]interface{}{
		"url": "https://example.gr",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(data)

	req := httptest.NewRequest(http.MethodGet, "/shorten", reqBody)
	w := httptest.NewRecorder()
	h.HandlePostURL(w, req)
	res := w.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	msg := string(bodyBytes)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assert.Equal(t, "url does not exist", msg)
}
