package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"


	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/raedahgroup/webtail/controllers"
	"github.com/raedahgroup/webtail/util"
	"github.com/raedahgroup/webtail/config"
)

var (

)

func main() {

	err := config.Parse()

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/ws/{b64file}", Use(controllers.WSHandler, controllers.AuthHandler, controllers.GetContext)).Methods("GET")
	router.HandleFunc("/", Use(controllers.RootHandler, controllers.AuthHandler, controllers.GetContext)).Methods("GET")
	/*router.HandleFunc("/login", Use(controllers.LoginHandler, controllers.GetContext)).Methods("POST")
	router.HandleFunc("/login", Use(controllers.LoginPageHandler, controllers.GetContext)).Methods("GET")*/
	router.HandleFunc("/logout", Use(controllers.LogoutHandler, controllers.AuthHandler, controllers.GetContext)).Methods("POST")
	router.HandleFunc("/user", Use(controllers.UserDetails, controllers.AuthHandler, controllers.GetContext))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	csrfHandler := csrf.Protect([]byte(util.GenerateSecureKey()),
		csrf.Secure(false), csrf.CookieName("X-CSRF-Token"))

	csrfRouter := Use(csrfHandler(router).ServeHTTP, controllers.CSRFExceptions)

	if *config.Secure == false {
		server := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%d", *config.Port), Handler: handlers.CombinedLoggingHandler(os.Stdout, csrfRouter)}
		panic(server.ListenAndServe())
	} else {
		serverCert, err := tls.LoadX509KeyPair(*config.Cert, *config.Key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read the certificates %s", err)
			panic(err)
		}
		tlsConfig := &tls.Config{
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS11,
			Certificates:             []tls.Certificate{serverCert},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		server := &http.Server{
			Addr:      fmt.Sprintf("0.0.0.0:%d", *config.Port),
			Handler:   handlers.CombinedLoggingHandler(os.Stdout, csrfRouter),
			TLSConfig: tlsConfig,
		}

		panic(server.ListenAndServeTLS(*config.Cert, *config.Key))
	}
}

// Use - Stacking middlewares
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
