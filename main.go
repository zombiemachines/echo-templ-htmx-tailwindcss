package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	//
	"github.com/zombiemachines/echo-templ-htmx-tailwindcss/view"
	//

	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
)

func main() {
	port := flag.String("port", "4000", "HTTP network port for http://localhost:")
	flag.Parse()

	// helloComponent := view.Hello("LYOKO")

	e := echo.New()
	e.Logger.SetLevel(echoLog.INFO)
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {

		context := c.Request().Context()
		writer := c.Response().Writer

		return view.Hello("Lyoko").Render(context, writer)

	})

	go func() {

		if err := e.StartTLS(":"+*port, "tls/cert.pem", "tls/key.pem"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("•• Shutting down server ◤•_•_•〓〓")
		}
	}()
	e.Logger.Infof("〓〓•_•_•◥ ••• Server Listening on https://localhost:" + *port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	e.Logger.Printf("〓〓•_•_•◥ •• Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	e.Logger.Infof("• Server gracefully stopped ◤•_•_•〓〓")

}
