package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Teapot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

func TestTeapotHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got status %d, want %d", res.Code, http.StatusTeapot)
	}
}

type MockUserService struct {
	RegisterFunc   func(user User) (string, error)
	UserRegistered []User
}

func (m *MockUserService) Register(user User) (id string, err error) {
	m.UserRegistered = append(m.UserRegistered, user)
	return m.RegisterFunc(user)
}

func TestUserServer_RegisterUser(t *testing.T) {
	t.Run("can register valid users", func(t *testing.T) {
		user := User{Name: "Tom"}
		wantId := "whatever"

		service := &MockUserService{
			RegisterFunc: func(user User) (string, error) {
				return wantId, nil
			},
		}
		server := NewUserServer(service)

		req := httptest.NewRequest(http.MethodPost, "/", user.toJSON())
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res, http.StatusCreated)

		if res.Body.String() != wantId {
			t.Errorf("got id %q, want %q", res.Body.String(), wantId)
		}

		if len(service.UserRegistered) != 1 {
			t.Errorf("got registered user %d, want %d", len(service.UserRegistered), 1)
		}

		if !reflect.DeepEqual(service.UserRegistered[0], user) {
			t.Errorf("got user %+v, want %+v", service.UserRegistered[0], user)
		}
	})

	t.Run("return 400 bad request if body is not valid", func(t *testing.T) {
		server := NewUserServer(nil)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("invalid user"))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res, http.StatusBadRequest)
	})

	t.Run("return 500 internal server error if the service fails", func(t *testing.T) {
		user := User{Name: "Tom"}

		service := &MockUserService{
			RegisterFunc: func(user User) (string, error) {
				return "", errors.New("register error")
			},
		}

		server := NewUserServer(service)

		req := httptest.NewRequest(http.MethodPost, "/", user.toJSON())
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res, http.StatusInternalServerError)
	})
}

func assertStatus(t *testing.T, res *httptest.ResponseRecorder, want int) {
	if res.Code != want {
		t.Errorf("got statue %d, want %d", res.Code, want)
	}
}

func (u *User) toJSON() io.Reader {
	b, _ := json.Marshal(u)
	return bytes.NewReader(b)
}
