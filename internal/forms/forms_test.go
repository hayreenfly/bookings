package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
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
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")

	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")

	form = New(postedValues)

	form.MinLength("some_field", 100)

	if form.Valid() {
		t.Error("form shows min length is valid when form value is less than the set 100 min length")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")

	form = New(postedValues)

	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("shows invalid minlength when form value does meet minLength")
	}

	isError = form.Errors.Get("another_field")

	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValue := url.Values{}
	form := New(postedValue)

	form.IsEmail("x")

	if form.Valid() {
		t.Error("shows valid email when field does not even exists")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "jane@smith.com")
	form = New(postedValue)

	form.IsEmail("email")

	if !form.Valid() {
		t.Error("shows invalid when it is an email")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "x")
	form = New(postedValue)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("shows valid email when it is not")
	}
}
