package handlers

import (
	"net/http"
	"strings"

	"github.com/saaste/pastebin/pkg/auth"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.URL.Query().Get("return")
	if redirectUrl == "" || !strings.HasPrefix(redirectUrl, "/") {
		redirectUrl = "/"
	}

	data := &PublicData{
		Title: h.appConfig.Title,
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			h.internalServerError(w, err, "failed to parse login form")
			return
		}

		password := r.Form.Get("password")

		if password == h.appConfig.Password {
			jwtToken, err := h.jwtParser.CreateJWT()
			if err != nil {
				h.internalServerError(w, err, "failed to create JWT token")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     auth.AuthCookieKey,
				Value:    jwtToken,
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
				MaxAge:   60 * 60 * 24 * 7,
			})
			http.Redirect(w, r, redirectUrl, http.StatusFound)
			return
		}

		data.Error = "Invalid password"
	}

	h.loadPublicTemplate(w, "login.html", data)
}
