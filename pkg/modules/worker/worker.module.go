package worker

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func InitModule(timeoutContext time.Duration, Router *chi.Mux, db *mongo.Database) {
	workerService := NewService(db, timeoutContext)

	NewWorkerController(Router, workerService)
}