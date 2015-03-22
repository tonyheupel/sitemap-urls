// sitemap-urls is a tool for retrieving information from a domain's sitemap
// and outputting the urls and their information in a tab-delimed
// output.
// For me, this is useful for comparing sitemap data with what
// is in an Elasticsearch index or as seed urls for crawling.
package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	domain := flag.String("d", "", "the domain to get the sitemap for")
	flag.Parse()
	validateArgs(*domain)

	urls, errors := RetrieveURLsForDomain(*domain) // TODO: Provide channel to recieve urls

	if errors != nil && len(errors) > 0 {
		log.Fatal("Errors:", errors)
	}

	for _, url := range urls {
		fmt.Printf("%s\t%s\n", url.Location, url.LastModified)
	}
}

func validateArgs(domain string) {
	if domain == "" {
		log.Fatal("domain (-d=<domain>) is required")
	}
}
