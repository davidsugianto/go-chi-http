package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Info("CreateUser")
}

func (h *Handler) SearchUser(w http.ResponseWriter, r *http.Request) {
	log.Info("SearchUser")
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Info("GetUser")
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Info("UpdateUser")
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Info("DeleteUser")
}
