package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	controller "github.com/Kelado/url-shortener/controllers"
	"github.com/Kelado/url-shortener/repositories"
	"github.com/go-chi/chi/v5"

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

func initHandler() *Handler {
	linkRepo := repositories.NewMockDB()
	controller := controller.NewController(testHostname, testCodeSize, linkRepo)
	return NewHandler(controller)
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

	req := httptest.NewRequest(http.MethodPost, "/shorten", reqBody)
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

	req := httptest.NewRequest(http.MethodPost, "/shorten", reqBody)
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

	req := httptest.NewRequest(http.MethodPost, "/shorten", reqBody)
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

	req := httptest.NewRequest(http.MethodPost, "/shorten", reqBody)
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

func TestGetExistingURL(t *testing.T) {
	initEnv()
	h := initHandler()

	expectedURL := "https://example.com"
	data := map[string]interface{}{
		"url": expectedURL,
	}
	shortURL := insertExampleURL(h, data)
	code := getCode(shortURL)

	req := httptest.NewRequest(http.MethodGet, "/{code}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("code", code)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	h.HandleGetURL(w, req)
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusSeeOther, res.StatusCode)

	redirectToURL, _ := res.Location()
	assert.Equal(t, expectedURL, redirectToURL.String())
}

func insertExampleURL(h *Handler, data map[string]interface{}) string {
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(data)

	req := httptest.NewRequest(http.MethodPost, "/shorten", reqBody)
	w := httptest.NewRecorder()
	h.HandlePostURL(w, req)
	res := w.Result()
	defer res.Body.Close()

	respData := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&respData)

	return respData["shortUrl"].(string)
}

func getCode(url string) string {
	return url[len(url)-testCodeSize:]
}
