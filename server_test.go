package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func TestHelloHandler(t *testing.T) {
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", Hello)

	//e.Logger.Fatal(e.Start(":8080"))

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("failed test")
	}
	if rec.Body.String() != "お仕事おつかれ様" {
		t.Fatal("failed test")
	}
	t.Log("HTTP response Ok")
}
