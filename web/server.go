package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	v1 "github.com/mattcardoso/foxbit-tech-test/api/v1"
)

type Server struct {
	port   string
	mux    *http.ServeMux
	logger *slog.Logger
}

func (s *Server) Run() {
	s.logger.Info("starting server on port " + s.port + "....")
	s.mux.HandleFunc("GET /api/sum", v1.Sum)
	s.mux.HandleFunc("GET /api/sub", v1.Sub)
	s.mux.HandleFunc("GET /api/mul", v1.Mul)
	s.mux.HandleFunc("GET /api/div", v1.Div)

	s.mux.HandleFunc("GET /api/healthcheck", v1.Healthcheck)
	serverLog := slog.NewLogLogger(s.logger.Handler(), slog.LevelError)
	server := http.Server{
		Addr:     ":" + s.port,
		Handler:  s.mux,
		ErrorLog: serverLog,
	}

	// Channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Goroutine to listen for termination signal
	go func() {
		<-sigChan
		fmt.Println("Shutting down server gracefully...")
		server.Close()
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("error starting the server")
	}
}

func NewServer() *Server {
	var logLevel slog.Level
	var appEnv = os.Getenv("ENV")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	logLevelEnv := strings.ToLower(os.Getenv("APP_LOG_LEVEL"))
	switch logLevelEnv {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "error":
		logLevel = slog.LevelError
	}

	logOpts := &slog.HandlerOptions{
		Level: logLevel,
	}
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, logOpts)
	if appEnv == "production" || appEnv == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, logOpts)
	}
	logger := slog.New(handler)
	return &Server{
		port:   port,
		mux:    http.NewServeMux(),
		logger: logger,
	}
}
