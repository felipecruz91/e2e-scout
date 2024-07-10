package main

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/labstack/echo"
	"github.com/tektoncd/pipeline/pkg/names"
)

func main() {

	// Package 'github.com/hashicorp/go-retryablehttp' with version <0.7.7 is affected by CVE-2024-6104 (MEDIUM)
	// https://github.com/advisories/GHSA-v6v8-xj6m-xwqh
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	_ = retryClient.StandardClient()

	// Package 'github.com/tektoncd/pipeline' with version 0.52.0 is affected by CVE-2023-37264 (LOW)
	// https://github.com/advisories/GHSA-w2h3-vvvq-3m53
	_ = names.SimpleNameGenerator.RestrictLength("foo")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
