package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/fizzbuzz"
	"hello/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloGinHandler(c *gin.Context) {
	numberStr := c.Param("number")
	n, err := strconv.Atoi(numberStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fb := fizzbuzz.Count(n)
	c.String(http.StatusOK, fb)
}
func helloEchoHandler(c echo.Context) error {
	numberStr := c.Param("number")
	n, err := strconv.Atoi(numberStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	fb := fizzbuzz.Count(n)
	return c.String(http.StatusOK, fb)
}

func main() {
	viper.SetDefault("app.addr", ":8000")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	db, err := gorm.Open(sqlite.Open("./users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/fizzbuzz/:number", helloEchoHandler)
	e.GET("/auth", credentialEchoHandler)
	e.POST("/users", user.NewEchoHandler(db))
	h := user.EchoHandle{DB: db}
	e.POST("/users", h.Handler)
	fmt.Println(viper.GetString("app.addr"))

	go func() {
		if err := e.Start(viper.GetString("app.addr")); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func ginRouter() {
	router := gin.Default()

	router.GET("/fizzbuzz/:number", helloGinHandler)

	router.Run(":8000")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) < 7 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	numberStr := vars["number"]
	n, err := strconv.Atoi(numberStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	fb := fizzbuzz.Count(n)
	fmt.Fprint(w, fb)
}

// Handler
// HandlerFunc
// Handle
// HandleFunc

func gorillamux() {
	db, err := gorm.Open(sqlite.Open("./users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", helloHandler).Methods(http.MethodGet)
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler).Methods(http.MethodGet)
	r.HandleFunc("/auth", credentialHandler).Methods(http.MethodPost)
	r.Handle("/users", user.NewHandler(db)).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func primes(n int) {
	for i := 1; i <= n; i++ {
		count := 0
		for j := i; j > 0; j-- {
			if (i % j) == 0 {
				count++
			}
		}
		if count == 2 {
			println(i)
		}
	}
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func credentialHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var cred Credential
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		Issuer:    cred.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": ss,
	})
}

func credentialEchoHandler(c echo.Context) error {
	var cred Credential

	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		Issuer:    cred.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": ss,
	})
}
