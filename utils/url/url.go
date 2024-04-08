package url

import (
	"net/url"
)

func GetDomainFromURL(urlStr string) string {
	parsedURL, _ := url.Parse(urlStr)
	host := parsedURL.Hostname()
	scheme := parsedURL.Scheme

	u := scheme + "://" + host + "/"
	return u
}

func GetPathFromURL(urlStr string) string {
	parsedURL, _ := url.Parse(urlStr)
	return parsedURL.Path
}
