package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	config Config
}

func New(config Config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		config: config,
	}

	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	fmt.Println("Connecting to redis")
	err := a.rdb.Ping(ctx).Err()

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to connect to redis: %w", err)
		}
	}()

	fmt.Println("Starting server on port ", a.config.ServerPort)
	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("server error: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
