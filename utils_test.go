package silscraper

import (
	"os"
	"testing"
)

func TestUtfBody(t *testing.T) {
	file, err := os.Open("./datatest/test.html")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	_, err = utfBody(file)
	if err != nil {
		t.Error(err)
	}
}
