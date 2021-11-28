package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

type WebApiMain struct {
	Config     Config
	DB         *sqlite.DB
	HTTPServer *http.Server
}

func NewMain() *WebApiMain {
	config := DefaultConfig()
	sqliteDB := sqlite.NewDB(config.DB)

	return &WebApiMain{
		Config:     config,
		DB:         sqliteDB,
		HTTPServer: http.NewServer(),
	}
}

func (m *WebApiMain) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}
	if *m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (m *WebApiMain) Run(ctx context.Context) (err error) {

	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	m.HTTPServer.BookService = application.NewBookService(m.DB)
	m.HTTPServer.AuthorService = application.NewAuthorService(m.DB)

	m.HTTPServer.Addr = m.Config.HTTP.Addr
	m.HTTPServer.Domain = m.Config.HTTP.Domain
	m.HTTPServer.HashKey = m.Config.HTTP.HashKey
	m.HTTPServer.BlockKey = m.Config.HTTP.BlockKey
	m.HTTPServer.GitHubClientID = m.Config.GitHub.ClientID
	m.HTTPServer.GitHubClientSecret = m.Config.GitHub.ClientSecret

	if err := m.HTTPServer.Open(); err != nil {
		return err
	}

	if m.HTTPServer.UseTLS() {
		go func() {
			log.Fatal(http.ListenAndServeTLSRedirect(m.Config.HTTP.Domain))
		}()
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