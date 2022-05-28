package pkg

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/thehaung/cloudflare-worker/pkg/config"
	"github.com/thehaung/cloudflare-worker/pkg/database"
	"github.com/thehaung/cloudflare-worker/pkg/modules/worker"
	"github.com/thehaung/cloudflare-worker/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"sync"
	"time"
)

type HttpServer struct {
	doOnce sync.Once
	Server *http.Server
	Router *chi.Mux
	DB     *mongo.Database
}

func NewHttpServer() {
	httpServer := &HttpServer{}

	httpServer.doOnce.Do(func() {
		httpServer.InitEnv()
		httpServer.InitDB()
		httpServer.InitRouter()
		httpServer.InitModule()
		httpServer.InitServer()
	})
}

func (s *HttpServer) InitEnv() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error while loading Env: %v", err)
	}
}

func (s *HttpServer) InitDB() {
	conn := config.NewConfig()
	s.DB = database.InitializeConnection("medh", conn.MongoURI())
}

func (s *HttpServer) InitRouter() {
	s.Router = chi.NewRouter()

	// A good base middleware stack
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.CleanPath)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(15 * time.Second))

	// Global prefix
	utils.SetGlobalPrefix("/api/v1", s.Router)
}

func (s *HttpServer) InitModule() {
	timeoutContextAPI := time.Duration(config.GetTimeDeadlineAPI()) * time.Millisecond

	worker.InitModule(timeoutContextAPI, s.Router, s.DB)
}

func (s *HttpServer) InitServer() {
	log.Println("Server Initialize....")
	port := config.GetServerPort()

	server := &http.Server{
		Handler:      s.Router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(fmt.Sprintf("Server listening on %s", config.GetAPIUrl()))

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
