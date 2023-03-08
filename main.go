package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/brotherlogic/aoctracker/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	port        = flag.Int("port", 8080, "The server port.")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
	globalPort  = flag.Int("global_port", 8082, "The server port.")
)

var (
	year = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "aoctracker_year",
		Help: "The year",
	})
)

type Server struct {
}

func (s *Server) Track(_ context.Context, req *pb.TrackRequest) (*pb.TrackResponse, error) {
	year.Set(float64(req.GetCurrentYear()))
	return &pb.TrackResponse{}, nil
}

func main() {
	flag.Parse()

	s := &Server{}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("aoctracker failed to listen on the serving port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterAOCTrackerServiceServer(gs, s)

	// Setup prometheus export
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
	}()

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("aoctracker failed to serve: %v", err)
	}
}
