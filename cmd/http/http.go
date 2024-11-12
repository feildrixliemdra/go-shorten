package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"go-shorten/internal/appcontext"
	"go-shorten/internal/bootstrap"
	"go-shorten/internal/handler"
	"go-shorten/internal/repository"
	"go-shorten/internal/server"
	"go-shorten/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// Start function handler starting http listener
func Start() {
	var (
		cfg          = bootstrap.NewConfig()
		err          error
		postgreConn  *sqlx.DB
		mongoDBConn  *mongo.Client
		appCtxOption = appcontext.Option{
			Config: cfg,
		}
		repo   *repository.Repository
		hndler *handler.Handler
		svc    *service.Service
		router *gin.Engine
	)

	// bootstrap dependency
	bootstrap.SetJSONFormatter()

	if cfg.Postgre.IsEnabled {
		postgreConn, err = bootstrap.InitiatePostgreSQL(cfg)
		if err != nil {
			log.Fatalf("error connect to PostgreSQL | %v", err)
		}

		//make sure connected
		err = postgreConn.Ping()
		if err != nil {
			log.Fatalf("failed to ping PostgreSQL | %v", err)
		}
	}

	if cfg.MongoDB.IsEnabled {
		mongoDBConn, err = bootstrap.InitiateMongoDB(cfg)
		if err != nil {
			log.Fatalf("error connect to MongoDB | %v", err)
		}

		//make sure connected
		err = mongoDBConn.Ping(context.Background(), nil)
		if err != nil {
			log.Fatalf("failed to ping MongoDB | %v", err)
		}
	}

	repo = repository.InitiateRepository(repository.Option{
		Option: appCtxOption,
		DB:     postgreConn,
	})
	svc = service.InitiateService(service.Option{
		Option:     appCtxOption,
		Repository: repo,
	})
	hndler = handler.InitiateHandler(handler.Option{
		Option:  appCtxOption,
		Service: svc,
	})

	router = bootstrap.InitiateGinRouter(cfg, hndler)

	serve := server.NewHTTPServer(cfg, router.Handler())
	defer serve.Done()

	if err := serve.Run(); err != nil {
		log.Fatalf("error running http server %v", err.Error())
	}

	return
}
