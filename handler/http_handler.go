package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserService interface {
	Register(user User) (id string, err error)
}

type UserServer struct {
	service UserService
}

func NewUserServer(service UserService) *UserServer {
	return &UserServer{service}
}

func (u *UserServer) Register(user User) (id string, err error) {
	return user.Id, nil
}

func (u *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode user payload: %v", err), http.StatusBadRequest)
		return
	}

	id, err := u.service.Register(user)

	if err != nil {
		http.Error(w, fmt.Sprintf("problem register user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, id)
}
