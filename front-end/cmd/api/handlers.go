package main

import (
	"context"
	"fmt"
	"front-end/protobuf/short"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

// @Tags URL
// @Summary URL shortener
// @Description save URL To DB and get short URL back
// @Accept json
// @Produce json
// @Param LongURL query string true "Long URL"
// @Success 200 {object} jsonResponse "Successful shorten"
// @Router /shorten [post]
func shortenerURLHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	longURL := queryParams.Get("LongURL")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", shortenerHost, shortenerGRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		_ = ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	c := short.NewUrlShortenerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &short.URLRequest{URL: longURL}
	shortURL, err := c.ShortenURL(ctx, req)
	if err != nil {
		_ = ErrorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "shorted URL successfully",
		Data:    shortURL,
	}
	_ = WriteJSON(w, http.StatusOK, payload)
}

// @Tags URL
// @Summary get URL back
// @Description get URL back from short URL
// @Accept json
// @Produce json
// @Param ShortURL query string true "Short URL"
// @Success 200 {object} jsonResponse "Successful returned"
// @Router /original [get]
func getOriginalURLHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	shortURL := queryParams.Get("ShortURL")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", shortenerHost, shortenerGRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		_ = ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	c := short.NewUrlShortenerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &short.URLRequest{URL: shortURL}
	longURL, err := c.GetOriginalURL(ctx, req)
	if err != nil {
		_ = ErrorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "got original URL successfully",
		Data:    longURL,
	}
	_ = WriteJSON(w, http.StatusOK, payload)
}
