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
