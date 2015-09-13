package silscraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/kennygrant/sanitize"
	"io"
	"net/http"
	"strings"
)

func docQueryFromUrl(url string) (html *goquery.Document, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	utfile, err := utfBody(res.Body)
	if err != nil {
		return
	}
	html, err = goquery.NewDocumentFromReader(utfile)
	return
}

func docQueryFromFile(file io.Reader) (html *goquery.Document, err error) {
	utfile, err := utfBody(file)
	if err != nil {
		return
	}
	html, err = goquery.NewDocumentFromReader(utfile)
	return
}

func utfBody(file io.Reader) (utfile io.Reader, err error) {
	utfile, err = iconv.NewReader(file, "windows-1252", "utf-8")
	return
}

func normalize(text string) string {
	text = toLower(text)
	text = trimSpace(text)
	text = noPunctuation(text)
	return text
}

func toLower(text string) string {
	return strings.ToLower(text)
}

func trimSpace(text string) string {
	return strings.TrimSpace(text)
}

func noPunctuation(text string) string {
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ":", "", -1)
	text = strings.Replace(text, ";", "", -1)
	text = strings.Replace(text, "(", "", -1)
	text = strings.Replace(text, ")", "", -1)
	return sanitize.Accents(text)
}
