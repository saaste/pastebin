package handlers

import (
	"errors"
	"html"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saaste/pastebin/pkg/auth"
	"github.com/saaste/pastebin/pkg/documents"
)

func (h *Handler) PasteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	publicPath := chi.URLParam(r, "public_path")
	if publicPath == "" {
		h.notFound(w)
		return
	}
	doc, err := h.documentStorage.GetByPublicPath(publicPath)
	if err != nil {
		if errors.Is(err, documents.ErrNotFound) {
			h.notFound(w)
			return
		}
		h.internalServerError(w, err, "Internal Server Error")
		return
	}

	data := &PublicData{
		Title:           h.appConfig.Title,
		IsAuthenticated: ctx.Value(auth.AuthContextKeyIsAuthenticated) == true,
		BaseURL:         h.appConfig.BaseURL,
		CurrentDocument: doc,
	}
	data.CurrentDocument.Content = html.EscapeString(data.CurrentDocument.Content)
	h.loadPublicTemplate(w, "paste.html", data)
}
