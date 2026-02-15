package server

import (
	"net/http"
	"encoding/json"
)

func (s *Server) GetV1GetUser(w http.ResponseWriter, r *http.Request, name string) {
	user, err := s.db.GetUserByName(name)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) PostV1CreateUser(w http.ResponseWriter, r *http.Request, name string) {
	err := s.db.AddUser(name)
	// TODO: check whether it is something else that causes error
	if err != nil {
		http.Error(w, "user already exists", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) DeleteV1DeleteUser(w http.ResponseWriter, r *http.Request, name string) {
	err := s.db.DeleteUserByName(name)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetV1GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

