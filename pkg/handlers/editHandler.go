package handlers

import (
	"fmt"
	"html"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saaste/pastebin/pkg/auth"
	"github.com/saaste/pastebin/pkg/syntax"
)

func (h *Handler) EditHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	doc, err := h.documentStorage.GetById(id)
	if err != nil {
		h.notFound(w)
		return
	}

	data := &PublicData{
		Title:             h.appConfig.Title,
		CurrentDocument:   doc,
		SupportedSyntaxes: syntax.SupportedSyntaxes,
		DefaultSyntax:     h.appConfig.DefaultSyntax,
		IsAuthenticated:   ctx.Value(auth.AuthContextKeyIsAuthenticated) == true,
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			h.internalServerError(w, err, "failed to parse login form")
			return
		}

		action := r.Form.Get("action")

		if action == "delete" {
			err = h.documentStorage.Delete(id)
			if err != nil {
				data.Error = err.Error()
			} else {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		} else if action == "save" {
			data.CurrentDocument.Name = r.Form.Get("name")
			data.CurrentDocument.Syntax = r.Form.Get("syntax")
			data.CurrentDocument.IsPublic = r.Form.Get("is_public") == "on"
			data.CurrentDocument.PublicPath = r.Form.Get("public_path")
			data.CurrentDocument.Content = r.Form.Get("content")

			err = h.documentStorage.Update(id, data.CurrentDocument, "fofo")
			if err != nil {
				data.Error = err.Error()
			} else {
				http.Redirect(w, r, fmt.Sprintf("/edit/%s", id), http.StatusFound)
				return
			}
		}

	}

	data.CurrentDocument.Content = html.EscapeString(data.CurrentDocument.Content)
	h.loadPublicTemplate(w, "edit.html", data)
}
