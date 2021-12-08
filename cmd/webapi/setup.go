package main

import (
	author_application "bookstore/internal/author/application"
	author_inmem "bookstore/internal/author/infrastructure/inmem"
	author_http "bookstore/internal/author/ui/http"
	platform_inmem "bookstore/internal/platform/db/inmem"
	platform_http "bookstore/internal/platform/http"
	"context"
	"fmt"
	"os"
	"os/signal"
)

type WebApiMain struct {
	Config     Config
	Repository *platform_inmem.InMemRepository
	HTTPServer *platform_http.Server
}

func NewMain() *WebApiMain {
	config := DefaultConfig()
	repository := platform_inmem.NewInMemRepository(config.DB.DSN)

	return &WebApiMain{
		Config:     config,
		Repository: repository,
		HTTPServer: platform_http.NewServer(),
	}
}

func (m *WebApiMain) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}
	if m.Repository != nil {
		if err := m.Repository.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (m *WebApiMain) Run(ctx context.Context) (err error) {
	m.Repository.AddMigration(author_inmem.Migration)

	if err := m.Repository.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	authorApplication := author_application.NewAuthorService(m.AuthorRepository)
	authorServer := author_http.NewAuthorServer(authorApplication)

	m.HTTPServer.RouteRegisters = []platform_http.RouteRegister{authorServer}
	m.HTTPServer.Addr = m.Config.HTTP.Addr
	m.HTTPServer.Domain = m.Config.HTTP.Domain

	if err := m.HTTPServer.Open(); err != nil {
		return err
	}

	return nil
}

func getContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	return ctx
}
