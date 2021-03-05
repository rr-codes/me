package tests

import (
	"encoding/json"
	"me/handlers"
	"me/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAboutInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleAboutInfo)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("got: %d; want: http.StatusOK", rr.Code)
	}

	want := models.GetAboutInfo()
	got := models.AboutInfo{}

	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	isEqual := reflect.DeepEqual(want, got)
	if !isEqual {
		t.Fatalf("got: %v; want: %v", got, want)
	}
}

type ResponseWriterMock struct {
	statusCode int
}

func (rw *ResponseWriterMock) Header() http.Header {
	return http.Header{}
}

func (rw *ResponseWriterMock) Error() string {
	return ""
}

func (rw *ResponseWriterMock) Write([]byte) (int, error) {
	return 500, rw
}

func (rw *ResponseWriterMock) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

func TestAboutInfoHandlerFail(t *testing.T) {
	writer := ResponseWriterMock{}
	handlers.HandleAboutInfo(&writer, nil)

	if writer.statusCode != 500 {
		t.Fatal("Expected status code to be 500")
	}
}
