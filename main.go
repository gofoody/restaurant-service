package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofoody/restaurant-service/pkg/config"
	"github.com/gofoody/restaurant-service/pkg/ctrl"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.New()

	initLogger(config.GetLogLevel())

	router := mountEndpoints()
	startService(config.GetHttpPort(), router)
}

func initLogger(logLevel string) {
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func mountEndpoints() *mux.Router {
	r := mux.NewRouter()

	statusCtrl := ctrl.NewStatusCtrl()
	r.HandleFunc("/api/status", statusCtrl.Show)

	restaurantCtrl := ctrl.NewRestaurantCtrl()
	r.HandleFunc("/api/restaurants/{restaurantId}", restaurantCtrl.Show)
	r.HandleFunc("/api/restaurants", restaurantCtrl.Create)

	return r
}

func startService(port int, router *mux.Router) {
	addr := fmt.Sprintf("localhost:%d", port)
	log.Infof("restaurant service running at:%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("failed to start restaurant service, error:%v", err)
	}
}
