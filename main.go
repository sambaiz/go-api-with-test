package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/sambaiz/go-api-with-test/handler"
    "github.com/gocraft/dbr"
    "github.com/facebookgo/grace/gracehttp"
    "github.com/labstack/echo/middleware"
    _ "github.com/go-sql-driver/mysql"
    "github.com/sambaiz/go-api-with-test/conf"
)

func main() {
    conf, _ := conf.Parse()
    conn, err := dbr.Open("mysql", conf.Db, nil)
    if err != nil {
        panic(err)
    }
    conn.SetMaxIdleConns(200)
    conn.SetMaxOpenConns(200)

    e := echo.New()

    // middlewares
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
    }))

    // endpoints
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.GET("/messages", handler.NewMessageWithSession(conn.NewSession(nil)).GetMessages)
    e.POST("/messages", handler.NewMessageWithSession(conn.NewSession(nil)).CreateMessage)
    std := standard.New(":1323")
    std.SetHandler(e)
    gracehttp.Serve(std.Server)
}
