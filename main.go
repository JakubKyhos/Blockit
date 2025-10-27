package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	scraper "github.com/JakubKyhos/Blockit.git/domainScraper"
	"github.com/elazarl/goproxy"
	_ "github.com/lib/pq"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	// Enable HTTPS interception (MITM)
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	// Block .cz domains
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			host := r.URL.Hostname() // strips port
			if strings.HasSuffix(host, ".cz") {
				log.Printf("Blocked request to .cz domain: %s", host)
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					fmt.Sprintf("Access to .cz domains is blocked: %s", host))
			}

			return r, nil
		},
	)

	HTMLdata, err := scraper.CrawlPage("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")

	log.Println("Proxy listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
