// handlers_test.go
package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    //"fmt"
    //"github.com/gorilla/sessions"
)

var (
    password = "Testing123"
)

func TestHashAndSalt(t *testing.T) {
    hashAndSalt([]byte(password))
    //fmt.Printf(salt)
}

func TestComparePasswords(t *testing.T) {
    var salt = hashAndSalt([]byte(password))
    passwordBytes := []byte(password)
    var result = comparePasswords(salt, passwordBytes)
    //fmt.Printf("%t", result)
    if result != true {
        t.Errorf("Password Test Failed")
    }
}

func TestGetMongoClient(t *testing.T) {
    client, err := getMongoClient(mongoHost, mongoUser, mongoPass, mongoTLS)
    if err != nil && client != nil {
		t.Errorf("Mongo Client Failed") 
	}

}

func TestRegisterHandler(t *testing.T) {
    t.Run("Returns Registration Page", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/register", nil)
        response := httptest.NewRecorder()

        registerHandler(response, request)
        got := response.Result()

        if got.StatusCode != http.StatusOK {
            t.Errorf("Unexpected status code %d", got.StatusCode)
        }

    })
}

func TestLoginHandler(t *testing.T) {
    t.Run("Returns Login Page", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/login", nil)
        response := httptest.NewRecorder()

        loginHandler(response, request)
        got := response.Result()

        if got.StatusCode != http.StatusOK {
            t.Errorf("Unexpected status code %d", got.StatusCode)
        }

    })
}

func TestHealthHandler(t *testing.T) {
    t.Run("Returns Health Status", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/healthz", nil)
        response := httptest.NewRecorder()

        healthHandler(response, request)

        got := response.Body.String()
        want := "Ready"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}

func TestLogoutHandler(t *testing.T) {
    t.Run("Returns Logout Page", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/logout", nil)
        response := httptest.NewRecorder()
        http.SetCookie(response, &http.Cookie{Name: "authenticated", Value: "true"})
        //session.Values["authenticated"] = true
        //session.Values["email"] = "test@vmware.com"
        //session.Save(request, response)
        request.Header.Set("cookie-name", "authenticated=true")

        logoutHandler(response, request)
        got := response.Result()

        if got.StatusCode != 302 {
            t.Errorf("Unexpected status code %d", got.StatusCode)
        }

    })
}