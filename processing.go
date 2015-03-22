package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// RetrieveSitemapForDomain will retrieve a sitemap for the given domain, assuming sitemap.xml is at the root.
// It supports auto-detecting whether to use a multi-page sitemap (using sitemapindex) or if it
// is a single-page sitemap using urlsets directly.
func RetrieveURLsForDomain(domain string) ([]URL, []error) {
	return requestSitemap(fmt.Sprintf("http://%s/sitemap.xml", domain))
}

func requestSitemap(url string) ([]URL, []error) {

	body, err := getURLContents(url)

	if err != nil {
		return nil, []error{err}
	}

	return processSitemapBody(body)
}

// getURLContents retrieves the page contents of a url and
// returns it as a string.
func getURLContents(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Golang Spider Bot v. 3.0")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}

func processSitemapBody(body string) ([]URL, []error) {
	siteindex, err := parseSitemapIndexFromString(body)

	if err == nil {
		// Got a sitemapindex, so process it
		return processSitemapSiteindex(siteindex)
	} else {
		// Probably got a regular sitemap with a urlset back
		urlset, err := parseURLSetFromString(body)

		if err != nil {
			return nil, []error{err}
		} else {
			return processSitemapURLSet(urlset), nil
		}
	}
}

func processSitemapSiteindex(sitemapindex SitemapIndex) ([]URL, []error) {
	sitemapc, errorc := make(chan []URL), make(chan []error)

	sitemaps := sitemapindex.Sitemaps
	for _, sitemap := range sitemaps {
		go func(location string) {
			urls, errors := requestSitemap(location)

			if errors != nil {
				errorc <- errors
			} else {
				sitemapc <- urls
			}
		}(sitemap.Location)
	}

	allURLs := make([]URL, 0)
	errors := make([]error, 0)

	for i := 0; i < len(sitemaps); i++ {
		select {
		case urls := <-sitemapc:
			allURLs = append(allURLs, urls...)
		case err := <-errorc:
			errors = append(errors, err...)
		}
	}

	return allURLs, errors
}

func processSitemapURLSet(urlset URLSet) []URL {
	return urlset.URLs
}

// parseURLSetFromString takes a single urlset XML element as a string
// and returns a single URLSet.
func parseURLSetFromString(urlsetString string) (URLSet, error) {
	urlsetReader := strings.NewReader(urlsetString)

	var urlset URLSet
	err := xml.NewDecoder(urlsetReader).Decode(&urlset)

	return urlset, err
}

// parseSitemapIndexFromString takes the content of a sitemapindex page as
// a string and returns a SitemapIndex.
func parseSitemapIndexFromString(sitemapString string) (SitemapIndex, error) {
	indexReader := strings.NewReader(sitemapString)

	var index SitemapIndex
	err := xml.NewDecoder(indexReader).Decode(&index)

	return index, err
}
