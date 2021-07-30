package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid content type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("invalid user agent header")
	}

	if ContentTypeJson != "application/json;charset=utf-8" {
		t.Error("invalid content type json header")
	}
	if ContentTypeXml != "application/xml" {
		t.Error("invalid content type xml header")
	}
	if ContentTypeSoap != "application/soap+xml" {
		t.Error("invalid content type soap header")
	}
	if ContentTypeCsv != "text/csv" {
		t.Error("invalid content type csv header")
	}
}
