package auth

import (
	public_auth_service "eccomerce_apis/internal/app/public/services/auth"

	"github.com/kataras/iris/v12"
)

func (api *PublicApiNoAuth) generatePublicToken(ctx iris.Context) {
	token, err := public_auth_service.GenerateToken(api.Config, ctx.Values().GetString("key"), nil)
	if err != nil {
		ctx.StopWithJSON(iris.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"role":        "guest",
		"session_uid": token.SessionId,
		"token":       token.TokenPair,
	})
}

func (api *PublicApiNoAuth) refreshToken(ctx iris.Context) {
	token, err := public_auth_service.RefreshToken(
		api.Config,
		ctx.URLParam("refresh_token"),
		ctx.URLParam("session_uid"),
		ctx.Values().GetString("key"),
	)

	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"session_uid": token.SessionId,
		"token":       token.TokenPair,
	})
}

func (api *PublicApiNoAuth) checkAuth(ctx iris.Context) {
	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"session": ctx.GetCookie("_fmSession"),
	})

}
