package handlers

import (
	"net/http"

	"github.com/saaste/pastebin/pkg/auth"
)

func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     auth.AuthCookieKey,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
	http.Redirect(w, r, "/login/", http.StatusFound)
}
