package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henrieto/account/handlers"
	"github.com/henrieto/account/models/repository"
	mock_storage "github.com/henrieto/account/storage/mock"
)

func TestSignupHandler(t *testing.T) {
	// initialize the user repository for mocking data base access
	repository.UserRepository = mock_storage.NewUserStorage()
	// set up the request data
	data := map[string]any{
		"first_name":       "Henry",
		"last_name":        "kalu-kennedy",
		"email":            "kalukennedyh@gmail.com",
		"password":         "09037873790*Henro",
		"confirm_password": "09037873790*Henro",
		"terms":            "true",
		"gender":           "male",
	}
	// convert the data to json byte
	json_data, _ := json.Marshal(data)
	// generate an io reader from the json byte
	byte_reader := bytes.NewReader(json_data)
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/signup", byte_reader)
	if err != nil {
		t.Fatal(err)
	}
	//create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// generate a http handler from the handler func
	handler := http.HandlerFunc(handlers.Signup)
	handler.ServeHTTP(rr, req)
	// Check if the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// initialize a map object for umarshaling the response data
	response_data := map[string]any{}
	// umarshal the response data in to the map object
	json.Unmarshal(rr.Body.Bytes(), &response_data)
	// check if the data in the response is null
	if response_data["status"] == "failed" {
		t.Error(" data is null")
	}
	// check if the data in the response
	switch data := response_data["data"].(type) {
	case map[string]any:
		// check if the body is what we expected
		if data["email"] != "kalukennedyh@gmail.com" {
			t.Error(" test failed ")
		}
	default:
		t.Error(" test failed ")
	}
}
