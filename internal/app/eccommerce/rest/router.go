package private_router

import (
	"eccomerce_apis/internal/app/eccommerce/rest/analytics"
	"eccomerce_apis/internal/app/eccommerce/rest/auth"
	authconfigs "eccomerce_apis/internal/app/eccommerce/rest/authConfigs"
	"eccomerce_apis/internal/app/eccommerce/rest/customers"
	"eccomerce_apis/internal/app/eccommerce/rest/delivery"
	"eccomerce_apis/internal/app/eccommerce/rest/orders"
	payments "eccomerce_apis/internal/app/eccommerce/rest/payment"
	"eccomerce_apis/internal/app/eccommerce/rest/products"
	"eccomerce_apis/internal/app/eccommerce/rest/storage"
	"eccomerce_apis/internal/app/eccommerce/rest/stores"
	"eccomerce_apis/internal/app/eccommerce/rest/subscriptions"
	"eccomerce_apis/internal/app/eccommerce/rest/users"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"os"

	"github.com/kataras/iris/v12/middleware/jwt"
)

func ConfigAPIS(c *config.Config) {

	//web

	publicApi := c.App.Party("api/")
	{
		publicApi.RegisterDependency(c)
		publicApi.PartyConfigure("/auth", new(auth.AuthApi))
		publicApi.PartyConfigure("/user", new(users.UserPublicApi))
	}

	verifier := jwt.NewVerifier(jwt.HS256, os.Getenv("SIGNING_KEY"))
	{
		verifier.WithDefaultBlocklist()
	}

	protectedApi := c.App.Party("api/v1")
	{
		verifyMiddleware := verifier.Verify(func() interface{} {
			return new(middlewares.JWTClaims)
		})
		protectedApi.UseRouter(middlewares.PrivateApiAuthorization, verifyMiddleware)
		protectedApi.RegisterDependency(c)

		protectedApi.PartyConfigure("/auth", new(auth.ProtectedAuthApi))
		protectedApi.PartyConfigure("/analytics", new(analytics.AnalyticsApi))
		protectedApi.PartyConfigure("/users", new(users.UserProtectApi))
		protectedApi.PartyConfigure("/customers", new(customers.CustomerApi))
		protectedApi.PartyConfigure("/stores", new(stores.StoreApi))
		protectedApi.PartyConfigure("/products", new(products.ProductApi))
		protectedApi.PartyConfigure("/storage", new(storage.StorageApi))
		protectedApi.PartyConfigure("/orders", new(orders.OrderApi))
		protectedApi.PartyConfigure("/subscriptions", new(subscriptions.SubscriptionApi))
		protectedApi.PartyConfigure("/config", new(authconfigs.AuthConfigApi))
		protectedApi.PartyConfigure("/payments", new(payments.PaymentApi))
		protectedApi.PartyConfigure("/delivery", new(delivery.DeliveryApi))

	}

}
