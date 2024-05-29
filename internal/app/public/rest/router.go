package public_router

import (
	"eccomerce_apis/internal/app/public/rest/auth"
	"eccomerce_apis/internal/app/public/rest/cart"
	"eccomerce_apis/internal/app/public/rest/customers"
	"eccomerce_apis/internal/app/public/rest/orders"
	"eccomerce_apis/internal/app/public/rest/products"
	stores "eccomerce_apis/internal/app/public/rest/store"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"os"

	"github.com/kataras/iris/v12/middleware/jwt"
)

func ConfigPublicAPIS(c *config.Config) {
	publicApiNoAuth := c.App.Party("/api")
	{

		publicApiNoAuth.UseRouter(middlewares.SessionMiddleware())
		publicApiNoAuth.RegisterDependency(c)
		publicApiNoAuth.PartyConfigure("/auth", new(auth.PublicApiNoAuth))
	}

	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")
	verifier := jwt.NewVerifier(jwt.HS256, API_SECRET_KEY)
	{
		verifier.WithDefaultBlocklist()
	}

	publicApi := c.App.Party("/api/v1")
	{
		verifyMiddleware := verifier.Verify(func() interface{} {
			return new(middlewares.ApisClaims)
		})

		publicApi.UseRouter(verifyMiddleware)
		publicApi.RegisterDependency(c)

		publicApi.PartyConfigure("/auth", new(auth.PublicAuthApi))
		publicApi.PartyConfigure("/products", new(products.ProductPublicApi))
		publicApi.PartyConfigure("/carts", new(cart.PublicCartApis))
		publicApi.PartyConfigure("/customer", new(customers.CustomerPublicApi))
		publicApi.PartyConfigure("/store", new(stores.StorePublicApi))
		publicApi.PartyConfigure("/orders", new(orders.OrderPublicApi))
	}
}
