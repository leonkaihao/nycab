// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	client "github.com/leonkaihao/nycab/pkg/api"
	"github.com/leonkaihao/nycab/pkg/cache"
	"github.com/leonkaihao/nycab/pkg/config"
	log "github.com/sirupsen/logrus"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/leonkaihao/nycab/api/swagger/handler"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
)

//go:generate swagger generate server --target ../../swagger --name Nycab --spec ../nycab.yaml --principal models.Principal --exclude-main

func getConfigs() (apiConfig *config.API, err error) {
	v, err := config.LoadConfig("api.nycab")
	if err != nil {
		return
	}
	apiConfig, err = config.GetAPIConfig(v)
	if err != nil {
		return
	}
	return
}

func configureFlags(api *operations.NycabAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NycabAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf
	apiConfig, err := getConfigs()
	if err != nil {
		log.Fatalln("Cannot read config: ", err)
	}
	client, err := client.NewClient(apiConfig)
	if err != nil {
		log.Fatalln(err)
	}
	c, err := cache.NewRedisCache(apiConfig)
	if err != nil {
		log.Errorf("failed to connect redis, %v, use mem instead", err)
		c, err = cache.NewMemCache()
		if err != nil {
			log.Errorf("failed to connect mem, %v, cache disabled", err)
		}
	}
	hdlr := handler.NewHandler(client, c)
	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DeleteCabsPickupsCountCacheHandler = operations.DeleteCabsPickupsCountCacheHandlerFunc(hdlr.DeleteCache)

	api.GetCabsPickupsCountHandler = operations.GetCabsPickupsCountHandlerFunc(hdlr.GetPickupCount)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
