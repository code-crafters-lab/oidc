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

type AuthorizationServer interface {
	op.Server
	RegisterRouter() chi.Router
	Run() error
}

type authorizationServer struct {
	op.LegacyServer
	logger *slog.Logger
}

func (a *authorizationServer) RegisterRouter() chi.Router {
	router := chi.NewRouter()

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
			//a.logger.Error("failed to stop web server: %v", err.Error())
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
