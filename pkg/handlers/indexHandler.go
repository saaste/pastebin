package handlers

import (
	"net/http"

	"github.com/saaste/pastebin/pkg/auth"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	docList, err := h.documentStorage.List()
	if err != nil {
		h.internalServerError(w, err, "failed to load the document list")
	}
	data := &PublicData{
		Title:           h.appConfig.Title,
		Documents:       docList.Documents,
		IsAuthenticated: ctx.Value(auth.AuthContextKeyIsAuthenticated) == true,
		BaseURL:         h.appConfig.BaseURL,
	}
	h.loadPublicTemplate(w, "index.html", data)
}
