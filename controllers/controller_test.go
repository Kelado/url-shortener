package controller

import (
	"testing"

	"github.com/Kelado/url-shortener/models"
	"github.com/Kelado/url-shortener/repositories"
	"github.com/stretchr/testify/assert"
)

var (
	testCodeSize = 6
	testHostname = "http://localhost:8000/"
)

func newController() *Controller {
	repo := repositories.NewMockDB()
	return NewController(testHostname, testCodeSize, repo)
}

func TestSameGeneratedCode(t *testing.T) {
	repeatedCode := "AGAINandAGAIN"
	finalCode := "___Finally___"
	c := newController()
	c.WithCodeGenerator(func() func() string {
		var timesInvoked int
		limit := 2
		return func() string {
			if timesInvoked < limit {
				timesInvoked++
				return repeatedCode
			}
			return finalCode
		}
	}())

	url := models.URL("https://example.com")
	linkReq := models.LinkRequest{
		OriginalURL: &url,
	}
	// Insert it once, so the code is already used
	c.CreateLink(linkReq)

	linkResp, err := c.CreateLink(linkReq)
	assert.Nil(t, err)
	codeGenerated := string(linkResp)[len(linkResp)-len(finalCode):]
	assert.Equal(t, finalCode, codeGenerated)
}

func TestCreateWithEmptyURL(t *testing.T) {
	c := newController()
	url := models.URL("")
	linkReq := models.LinkRequest{
		OriginalURL: &url,
	}
	_, err := c.CreateLink(linkReq)
	assert.NotNil(t, err)
}

func TestCreateWithWrongURL(t *testing.T) {
	c := newController()
	url := models.URL("https://example.gr")
	linkReq := models.LinkRequest{
		OriginalURL: &url,
	}
	_, err := c.CreateLink(linkReq)
	assert.NotNil(t, err)
}

func TestGetKownURL(t *testing.T) {
	expectedURL := models.URL("https://example.com")
	c := newController()
	linkReq := models.LinkRequest{
		OriginalURL: &expectedURL,
	}
	shortURL, _ := c.CreateLink(linkReq)
	code := getCode(string(shortURL))
	gotURL, err := c.GetLink(code)
	assert.Nil(t, err)
	assert.Equal(t, expectedURL, gotURL)
}

func TestGetUnkownURL(t *testing.T) {
	c := newController()
	_, err := c.GetLink("RandomCode")
	assert.NotNil(t, err)
}

func getCode(url string) string {
	return url[len(url)-testCodeSize:]
}
