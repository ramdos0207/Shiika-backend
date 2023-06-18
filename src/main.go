package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
)

var (
	db *sqlx.DB
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	conf := mysql.Config{
		User:      os.Getenv("NS_MARIADB_DATABASE"),
		Passwd:    os.Getenv("NS_MARIADB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("NS_MARIADB_HOSTNAME") + ":" + os.Getenv("NS_MARIADB_PORT"),
		DBName:    os.Getenv("NS_MARIADB_DATABASE"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	_db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conntected")
	db = _db

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}

	allowOrigin := "https://servisenmae"
	if os.Getenv("DEVELOPMENT") == "true" {
		allowOrigin = "http://localhost:5173"
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{allowOrigin},
		AllowCredentials: true,
	}))
	e.Use(session.Middleware(store))

	e.POST("/login", postLoginHandler)
	e.POST("/register", postRegisterHandler)

	withLogin := e.Group("")
	withLogin.Use(checkLogin)

	e.GET("/kaminoku", getKaminokuHandler)
	withLogin.POST("/kaminoku", postKaminokuHandler)
	e.GET("/kaminoku/:kaminoku_id", getKaminokuDetailHandler)

	e.GET("/kaminoku/:kaminoku_id/simonoku", getSimonokuHandler)
	withLogin.POST("/kaminoku/:kaminoku_id/simonoku", postSimonokuHandler)
	e.GET("/simonoku", getAllSimonokuHandler)

	withLogin.GET("/user/kaminoku", getUserKaminokuHandler)
	withLogin.GET("/user/simonoku", getUserSimonokuHandler)
	e.Start(":8080")
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("sessions", c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}
		if sess.Values["userName"] == nil {
			return c.String(http.StatusForbidden, "please login")
		}
		c.Set("userName", sess.Values["userName"].(string))

		return next(c)
	}
}
