package handlers

import(
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIndex)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	expectedStatus := http.StatusOK
	if status := recorder.Code; status != expectedStatus {
		t.Errorf("GetIndex returned wrong status code: got %v want %v", status, expectedStatus)
	}

	expectedBody := `I am Groot!`
	if recorder.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expectedBody)
	}
}

func TestGetHello(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetHello)

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	expectedStatus := http.StatusOK
	if status := recorder.Code; status != expectedStatus {
		t.Errorf("GetIndex returned wrong status code: got %v want %v", status, expectedStatus)
	}

	expectedBody := `Hello!`
	if recorder.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expectedBody)
	}
}
