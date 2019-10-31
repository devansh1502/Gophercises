package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/debug/", sourceCodeHandler).Methods("GET")
	router.HandleFunc("/panic", panicDemo).Methods("GET")
	router.HandleFunc("/", hello)
	return router
}
func TestSourceCodeHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/debug/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	if status := response.Code; status == http.StatusInternalServerError {
		t.Logf("%v", status)
	}
	request, _ = http.NewRequest("GET", "/debug/?line=24&path=/home/gs-1547/go/src/Gophercises/Exercise15", nil)
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	if status1 := response.Code; status1 == http.StatusInternalServerError {
		t.Logf("%v", status1)
	}
	request, _ = http.NewRequest("GET", "/debug/?line=24&path=/usr/local/go/src/runtime/debug/stack.go", nil)
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	if status2 := response.Code; status2 != http.StatusOK {
		t.Logf("%v", status2)
	}
}

func TestHello(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
}

func HandlerRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	response := httptest.NewRecorder()
	// response.Result()
	handler.ServeHTTP(response, req)
	return response, err
}

func TestPanicDemo(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	HandlerRequest("GET", "/panic", devMw(handler))
}

func TestPanicAfterDemo(t *testing.T) {
	handler := http.HandlerFunc(panicAfterDemo)
	HandlerRequest("GET", "/panic-after", devMw(handler))
}
func TestMakeLinks(t *testing.T) {
	output := "goroutine 35 [running]:\nruntime/debug.Stack(0xc0001419c8, 0x1, 0x1)\n\t/usr/local/go/src/runtime/debug/stack.go:24 +0x9d"
	result := makeLinks(output)
	if result == " " {
		t.Logf("no output:%v", result)
	}
}
func TestMain(t *testing.T) {
	go main()
}
