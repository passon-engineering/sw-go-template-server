package webserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/tpasson/sw-go-logger-lib/logger"
	"github.com/tpasson/sw-go-template-server/modules/application"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtKey = []byte("your_secret_key")

// Credentials struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct to be encoded in the JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// customLogWriter redirects log messages to the application's custom logger.
type customLogWriter struct {
	app *application.Application
}

func (w *customLogWriter) Write(p []byte) (n int, err error) {
	// Here, you might want to parse the log message or directly log it as is.
	// This example logs the message directly as an informational entry.
	w.app.Logger.Entry(logger.Container{
		Status: logger.STATUS_INFO,
		Info:   string(p),
	})
	return len(p), nil
}

func Init(app *application.Application) {
	// Redirect standard log output to the custom logger.
	log.SetOutput(&customLogWriter{app: app})

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(handleRoot(app))

	// Authentication routes
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/content", authenticateMiddleware(handleContent(app))).Methods("GET")

	server := http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	// Check if SSL certificates are provided
	if app.Config.WebTlsCert != "" && app.Config.WebTlsKey != "" {
		// HTTPS configuration
		server.Addr = app.Config.HttpsAddress
		handleRedirectToHttps(app) // Assuming this sets up HTTP to HTTPS redirection

		// Start HTTPS server
		err := server.ListenAndServeTLS(app.Config.WebTlsCert, app.Config.WebTlsKey)
		if err != nil {
			logFatal(app, "Could not initialize HTTPS server: "+err.Error())
		}
	} else {
		// HTTP configuration
		server.Addr = app.Config.HttpAddress

		// Start HTTP server
		err := server.ListenAndServe()
		if err != nil {
			logFatal(app, "Could not initialize HTTP server: "+err.Error())
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate credentials (this is just a simple example, in production you should use a database)
	if creds.Username != "your_username" || creds.Password != "your_password" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func authenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Create a new token with updated expiration time
		expirationTime := time.Now().Add(1 * time.Minute)
		claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		newTokenString, err := newToken.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as a cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   newTokenString,
			Expires: expirationTime,
		})

		next.ServeHTTP(w, r)
	})
}

func handleContent(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := app.MongoClient.Database("your_database").Collection("your_collection")

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			http.Error(w, "Failed to fetch data from MongoDB", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.TODO())

		var parts []bson.M
		if err := cursor.All(context.TODO(), &parts); err != nil {
			http.Error(w, "Failed to parse MongoDB data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(parts)
	}
}

func logFatal(app *application.Application, message string) {
	app.Logger.Entry(logger.Container{
		Status: logger.STATUS_ERROR,
		Error:  message,
	})
	time.Sleep(2 * time.Second)
	os.Exit(1)
}
