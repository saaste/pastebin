package handlers

import (
	"fmt"
	"html"
	"net/http"

	"github.com/saaste/pastebin/pkg/auth"
	"github.com/saaste/pastebin/pkg/documents"
	"github.com/saaste/pastebin/pkg/syntax"
)

func (h *Handler) NewHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := &PublicData{
		Title: h.appConfig.Title,
		CurrentDocument: &documents.Document{
			Syntax: h.appConfig.DefaultSyntax,
		},
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

		data.CurrentDocument.Name = r.Form.Get("name")
		data.CurrentDocument.Syntax = r.Form.Get("syntax")
		data.CurrentDocument.IsPublic = r.Form.Get("is_public") == "on"
		data.CurrentDocument.PublicPath = r.Form.Get("public_path")
		data.CurrentDocument.Content = r.Form.Get("content")

		err = h.documentStorage.Create(data.CurrentDocument)
		if err != nil {
			data.Error = err.Error()
		} else {
			http.Redirect(w, r, fmt.Sprintf("/edit/%s", data.CurrentDocument.Id), http.StatusFound)
			return
		}

	}

	data.CurrentDocument.Content = html.EscapeString(data.CurrentDocument.Content)
	h.loadPublicTemplate(w, "new.html", data)
}
