package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/upDJ/Image-Processing-Service/models"
	"github.com/upDJ/Image-Processing-Service/routers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := routers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestRegisterRoute(t *testing.T) {
	router := routers.SetupRouter()
	router = routers.SetupUserRoutes(router)

	user := models.User{ID: "1", Username: "testing", Password: "test"}
	userJson, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), string(userJson))
}

func TestRegisterRouteUserNotInBody(t *testing.T) {
	router := routers.SetupRouter()
	router = routers.SetupUserRoutes(router)

	user := models.User{ID: "1", Username: "testing", Password: "test"}
	userJson, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	fakeUser := models.User{
		ID: "not a real ID", Username: "not a real uname", Password: "not a real password",
	}
	fakeUserJson, _ := json.Marshal(fakeUser)

	assert.Equal(t, 200, w.Code)
	assert.NotContains(t, w.Body.String(), string(fakeUserJson))
}
