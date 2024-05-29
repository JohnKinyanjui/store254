package auth

import (
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
)

func setAuthCookies(ctx iris.Context, refreshToken string, session string) {
	ctx.SetCookie(&http.Cookie{
		Name:     "_fmSession",
		Value:    refreshToken,
		Expires:  time.Now().UTC().Add(time.Duration(730 * time.Hour)),
		Secure:   true,
		SameSite: iris.SameSiteNoneMode,
	}, iris.CookieSecure)

	ctx.SetCookie(&http.Cookie{
		Name:     "_fmSessionIndentification",
		Value:    session,
		Expires:  time.Now().UTC().Add(time.Duration(730 * time.Hour)),
		Secure:   true,
		SameSite: iris.SameSiteNoneMode,
	}, iris.CookieSecure)

}
