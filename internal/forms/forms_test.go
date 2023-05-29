package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// test form valid post
func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("form invalud when should have been valid")
	}
}

// test for required field function
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		// this generates error when nil data are in post form required field
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{} // creating new post data
	postedData.Add("a", "a")   // append post data with field and value
	postedData.Add("b", "b")
	postedData.Add("c", "c")
	r, _ = http.NewRequest("POST", "/whatever", nil) // new post request
	r.PostForm = postedData                          // add post data to request
	form = New(r.PostForm)                           // create new post form with posted data
	form.Required("a", "b", "c")                     // determine required field
	if !form.Valid() {
		t.Error("shows does not have required field when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}
	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("shows form does have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non existant field")
	}
	is_error := form.Errors.Get("x")
	if is_error == "" {
		t.Error("should have an error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("some_field", "some_field")
	form = New(postedData)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 met when data is shorter")
	}
	postedData = url.Values{}
	postedData.Add("another_field", "adb123")
	form = New(postedData)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("show min length of 1 not met when data met")
	}
	is_error = form.Errors.Get("another_field")
	if is_error != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_IsValidEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("show valid email for non existant field")
	}
	postedData = url.Values{}
	postedData.Add("email", "a@a.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got invaliud error for valid email")
	}
}

func TestForm_IsInValidEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("show valid email for non existant field")
	}
	postedData = url.Values{}
	postedData.Add("email", "am")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid email for invalid email")
	}
}
