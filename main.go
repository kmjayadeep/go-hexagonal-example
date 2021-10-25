package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kmjayadeep/go-hexagonal-example/api"
	"github.com/kmjayadeep/go-hexagonal-example/repository/memory"
	"github.com/kmjayadeep/go-hexagonal-example/shortner"
)

func main() {
  repo := chooseRepo()
  service := shortner.NewRedirectService(repo)
  handler := api.NewHandler(service)

  r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  r.Get("/{code}", handler.Get)
  r.Post("/", handler.Post)


  errs := make(chan error, 2)

  go func() {
    port := httpPort()
    fmt.Println("listening on port", port)
    errs <- http.ListenAndServe(port, r)
  }()

  go func() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT)
    errs <- fmt.Errorf("%s", <- c)
  }()

  fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
  port := "8000"
  if os.Getenv("PORT") != "" {
    port = os.Getenv("PORT")
  }
  return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortner.RedirectRepository {
  repo, err := memory.NewMemoryRepository()

  if err != nil {
    log.Fatal(err)
  }

  return repo
}
