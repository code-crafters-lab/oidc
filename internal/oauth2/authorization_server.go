package oauth2

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/run"
	"github.com/zitadel/oidc/v3/pkg/op"
)

type ProtocolEndpoints interface {
	OAuth2AuthorizationEndpoint() *op.Endpoint
	OAuth2DeviceAuthorization() *op.Endpoint
	OAuth2DeviceVerification() *op.Endpoint
	OAuth2Token() *op.Endpoint
	OAuth2TokenIntrospection() *op.Endpoint
	OAuth2TokenRevocationEndpoint() *op.Endpoint
	OAuth2AuthorizationServerMetadataEndpoint() *op.Endpoint
	JWKSetEndpoint() *op.Endpoint

	OpenIDConnect10ProviderConfigurationEndpoint() *op.Endpoint
	OpenIDConnect10LogoutEndpoint() *op.Endpoint
	OpenIDConnect10UserInfoEndpoint() *op.Endpoint
	OpenIDConnect10ClientRegistrationEndpoint() *op.Endpoint
}

type AuthorizationServer interface {
	GetMetadata() *AuthorizationServerMetadata
	RegisterRouter() chi.Router
	Run() error
}

type authorizationServer struct {
	logger *slog.Logger
}

func (a *authorizationServer) GetMetadata() *AuthorizationServerMetadata {
	return nil
}

func (a *authorizationServer) RegisterRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		a.logger.Info("OpenID Connect Discovery Document")
		writer.Header().Set("Location", "https://example.com/.well-known/openid-configuration")
		writer.WriteHeader(http.StatusFound)
	})

	router.Get(AuthorizationServerMetadataEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		a.logger.Info("OpenID Connect Discovery Document")
	})

	//router.Mount("/", nil)
	//router.Mount("/oauth2", func()
	return router
}

func (a *authorizationServer) Run() error {
	g := &run.Group{}
	addHttp(g, a)
	addExtra(g)
	return g.Run()
}

func addHttp(g *run.Group, a *authorizationServer) {
	var httpSrv *http.Server
	g.Add(func() error {
		httpSrv = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", "", 80),
			Handler: a.RegisterRouter(),
		}
		return httpSrv.ListenAndServe()
	}, func(err error) {
		if err := httpSrv.Close(); err != nil {
			a.logger.Error("failed to stop web server: %v", err)
		}
	})
}

func addExtra(g *run.Group) {
	g.Add(run.SignalHandler(context.TODO(), syscall.SIGINT, syscall.SIGTERM))
}

func NewAuthorizationServer() AuthorizationServer {
	server := &authorizationServer{
		logger: slog.Default(),
	}
	return server
}
