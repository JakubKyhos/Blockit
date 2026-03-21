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
			host := strings.ToLower(r.URL.Hostname())

			if host == "" {
				return r, nil
			}

			// Extract TLD safely
			parts := strings.Split(host, ".")
			tld := parts[len(parts)-1]

			// Strip www
			web, _ := strings.CutPrefix(host, "www.")
			web = strings.TrimSpace(web)

			fmt.Printf("host: %s\n", host)
			fmt.Printf("web: %s\n", web)

			// 1. Whitelist → allow immediately
			if allowed, err := dbQueries.GetWhitelistDom(context.Background(), web); err == nil {
				log.Printf("Access granted to domain: %s", allowed.Name)
				return r, nil
			}

			// 2. Blacklist → block immediately
			if blocked, err := dbQueries.GetBlacklistDom(context.Background(), web); err == nil {
				log.Printf("Blocked request (blacklist): %s", blocked.Name)
				return r, goproxy.NewResponse(
					r,
					goproxy.ContentTypeText,
					http.StatusForbidden,
					fmt.Sprintf("Domain %s is blacklisted.", blocked.Name),
				)
			}

			// 3. TLD block → block
			if domain, err := dbQueries.GetDomain(context.Background(), tld); err == nil && domain.IsBlocked {
				log.Printf("Blocked request (TLD): %s", domain.Name)
				return r, goproxy.NewResponse(
					r,
					goproxy.ContentTypeText,
					http.StatusForbidden,
					fmt.Sprintf("Access to TLD %s is blocked.", domain.Name),
				)
			}

			// 4. Default → allow
			return r, nil
		},
	)

	log.Println("Proxy listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
