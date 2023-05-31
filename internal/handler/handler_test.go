package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"generals", "/generals", "GET", []postData{}, http.StatusOK},
	{"majors", "/majors", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"rs", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-avaliable", "GET", []postData{}, http.StatusOK},
	{"psa", "/search-avaliable", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2021-01-01"},
	}, http.StatusOK},
	{"psaj", "/search-avaliable-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2021-01-01"},
	}, http.StatusOK},
	{"pmr", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Ishan"},
		{key: "last_name", value: "Shrestha"},
		{key: "email", value: "a@a.com"},
		{key: "phone", value: "12334576"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
