package server

import (
	"board/internal/db"
	"net/http"
	"encoding/json"
)

func (s *Server) GetV1GetUser(w http.ResponseWriter, r *http.Request, name string) {
	user, err := s.db.GetUser(name)

	if err == db.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) PutV1CreateUser(w http.ResponseWriter, r *http.Request, name string) {
	err := s.db.AddUser(name)

	if err == db.ErrUserExists {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) DeleteV1DeleteUser(w http.ResponseWriter, r *http.Request, name string) {
	err := s.db.DeleteUser(name)

	if err == db.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

