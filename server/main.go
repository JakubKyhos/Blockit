package main

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/elazarl/goproxy"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	caCert, err := tls.LoadX509KeyPair("./goproxy/ca.pem", "./goproxy/key.pem")
	if err != nil {
		log.Fatalf("failed to load cert: %v", err)
	}
	goproxy.GoproxyCa = caCert
	goproxy.GoproxyCa.Leaf, _ = x509.ParseCertificate(caCert.Certificate[0])
	log.Printf("MITM CA Subject: %s", goproxy.GoproxyCa.Leaf.Subject)
	log.Printf("MITM CA Thumbprint: %X", sha1.Sum(goproxy.GoproxyCa.Leaf.Raw))

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	// Enable HTTPS interception (MITM)
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			host := r.URL.Hostname() // strips port
			slice := strings.Split(host, ".")
			tld := slice[len(slice)-1]

			domain, err := dbQueries.GetDomain(context.Background(), tld)
			if err != nil {
				log.Printf("Blocked request to webpage: %s", host)
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					fmt.Sprintf("Access to webpage %s is blocked. Domain is either malformed or isn't present in database. Try updating DB and loading the webpage again.", host))
			} else if domain.IsBlocked == true {
				log.Printf("Blocked request to domain: %s", domain.Name)
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					fmt.Sprintf("Access to domain %s is blocked.", domain.Name))
			}

			return r, nil
		},
	)

	log.Println("Proxy listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
