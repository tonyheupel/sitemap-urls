# sitemap-urls

## Overview

```sitemap-urls``` is a  sitemap tool, pass it a domain and it will spit out sitemap info, specifically all urls in the sitemap, even if they are spread across sitemap index pages.

The urls and publish dates are output as tab-delimited lines to stdout so that you may pipe the output anywhere you would like.

## Building the tool

```
$ ./build
```

## Examples of Running the tool
```
$ sitemap-urls -d mydomain.com > urls.txt      # put all urls to a text file
$ sitemap-urls -d mydomain.com | grep "xml$"   # show all files with "xml" at the end
```