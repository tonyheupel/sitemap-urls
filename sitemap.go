// The sitemap package provides convenient processing of
// web site sitemap files (assuming domain/sitemap.xml).
// It supports both single list sitemaps and sitemaps
// spread across multiple siteindex pages.
package main

import (
	"encoding/xml"
)

// Sitemap represents a single sitemap entry within a sitemapindex element.
type Sitemap struct {
	XMLName      xml.Name `xml:"sitemap"`
	Location     string   `xml:"loc"`
	LastModified string   `xml:"lastmod"`
}

// SitemapIndex represents the root of a multi-page sitemap.
type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

// URL represents the url element of a sitemap urlset.
type URL struct {
	XMLName         xml.Name `xml:"url"`
	Location        string   `xml:"loc"`
	LastModified    string   `xml:"lastmod"`
	ChangeFrequency string   `xml:"changefreq"`
	Priority        float64  `xml:"priority"`
}

// URLSet represents a sitemap urlset element that contains urls.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}
