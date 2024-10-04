package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/saaste/pastebin/pkg/auth"
	"github.com/saaste/pastebin/pkg/config"
	"github.com/saaste/pastebin/pkg/documents"
)

type Handler struct {
	appConfig       *config.AppConfig
	jwtParser       *auth.JwtParser
	documentStorage *documents.Storage
}

type PublicData struct {
	Title             string
	IsAuthenticated   bool
	Error             string
	CurrentDocument   *documents.Document
	Documents         []documents.Document
	SupportedSyntaxes []string
	DefaultSyntax     string
	BaseURL           string
}

func NewHandler(appConfig *config.AppConfig) *Handler {
	return &Handler{
		appConfig:       appConfig,
		jwtParser:       auth.NewJwtParser(appConfig),
		documentStorage: documents.NewStorage(),
	}
}

func (h *Handler) loadPublicTemplate(w http.ResponseWriter, templateFile string, data *PublicData) {
	funcMap := template.FuncMap{}

	t, err := template.New("").Funcs(funcMap).ParseFiles(fmt.Sprintf("ui/%s/base.html", h.appConfig.Theme), fmt.Sprintf("ui/%s/%s", h.appConfig.Theme, templateFile))
	if err != nil {
		h.internalServerError(w, err, "failed to parse diary templates")
		return
	}

	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		h.internalServerError(w, err, "failed to execute diary templates")
		return
	}
}

func (h *Handler) internalServerError(w http.ResponseWriter, err error, message string) {
	fmt.Printf("ERROR: %s: %v\n", message, err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
