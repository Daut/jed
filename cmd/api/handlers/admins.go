package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	db "github.com/daut/jed/sqlc"
)

func (h *Handler) AdminRead(w http.ResponseWriter, r *http.Request) {
	// Needs admin authentication

	username := r.PathValue("username")
	// TODO: validate username

	admin, err := h.Queries.GetAdmin(r.Context(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.NotFound(w)
		} else {
			h.ServerError(w, err)
		}
		return
	}

	h.WriteJSON(w, http.StatusOK, admin, nil)
}

func (h *Handler) AdminList(w http.ResponseWriter, r *http.Request) {
	// Needs admin authentication

	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "1"
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		h.ClientError(w, http.StatusBadRequest)
		return
	}

	args := &db.ListAdminsParams{
		Limit:  10,
		Offset: (int32(page) - 1) * 10,
	}
	admins, err := h.Queries.ListAdmins(r.Context(), *args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.NotFound(w)
		} else {
			h.ServerError(w, err)
		}
		return
	}

	if len(admins) == 0 {
		h.NotFound(w)
		return
	}

	h.WriteJSON(w, http.StatusOK, admins, nil)
}
