package endpoints

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"disk"
	"strings"
	"os"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io"
)

func TestUp(t *testing.T) {
	req, err := http.NewRequest("HEAD", "/sample/up/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Up)

	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)

	if r.Code != 200 {
		t.Errorf("Wrong status code %d", r.Code)
	}
}

func BenchmarkUp(b *testing.B) {
	req, err := http.NewRequest("HEAD", "/sample/up/", nil)
	if err != nil {
		b.Fatal(err)
	}

	handler := http.HandlerFunc(Up)

	r := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		handler.ServeHTTP(r, req)

		if r.Code != 200 {
			b.Errorf("Wrong status code %d", r.Code)
		}
	}
}

func TestExistsNotFound(t *testing.T) {
	req, err := http.NewRequest("HEAD", "/sample/exists/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Exists)

	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)

	if r.Code != 404 {
		t.Errorf("Wrong status code %d", r.Code)
	}
}

func TestExistsFound(t *testing.T) {
	err := disk.WriteToStorage(strings.NewReader("test"), "123", 4)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("HEAD", "/sample/exists/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Exists)

	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)

	if r.Code != 200 {
		t.Errorf("Wrong status code %d", r.Code)
	}

	fPath, err := disk.ConvertToStoragePath("123")
	if err != nil {
		t.Error(err)
	}

	os.Remove(fPath)
}

func BenchmarkExistsNotFound(b *testing.B) {
	req, err := http.NewRequest("HEAD", "/sample/exists/123", nil)
	if err != nil {
		b.Fatal(err)
	}

	handler := http.HandlerFunc(Exists)

	r := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		handler.ServeHTTP(r, req)

		if r.Code != 404 {
			b.Errorf("Wrong status code %d", r.Code)
		}
	}
}

func BenchmarkExistsFound(b *testing.B) {
	err := disk.WriteToStorage(strings.NewReader("test"), "12344", 4)
	if err != nil {
		b.Error(err)
	}

	req, err := http.NewRequest("HEAD", "/sample/exists/12344", nil)
	if err != nil {
		b.Fatal(err)
	}

	handler := http.HandlerFunc(Exists)

	r := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		handler.ServeHTTP(r, req)

		if r.Code != 200 {
			b.Errorf("Wrong status code %d", r.Code)
		}
	}

	fPath, err := disk.ConvertToStoragePath("12344")
	if err != nil {
		b.Error(err)
	}

	os.Remove(fPath)
}


func TestDownloadNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/sample/download/123123", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Download)

	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)

	if r.Code != 404 {
		t.Errorf("Wrong status code %d", r.Code)
	}
}


func TestDownloadFound(t *testing.T) {
	err := disk.WriteToStorage(strings.NewReader("test"), "123", 4)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", "/sample/download/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Download)

	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)

	if r.Code != 200 {
		t.Errorf("Wrong status code %d", r.Code)
	}

	fPath, err := disk.ConvertToStoragePath("123")
	if err != nil {
		t.Error(err)
	}

	os.Remove(fPath)
}


func BenchmarkDownloadNotFound(b *testing.B) {
	req, err := http.NewRequest("GET", "/sample/download/1231233", nil)
	if err != nil {
		b.Fatal(err)
	}

	handler := http.HandlerFunc(Download)

	r := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		handler.ServeHTTP(r, req)

		if r.Code != 404 {
			b.Errorf("Wrong status code %d", r.Code)
		}
	}
}

func BenchmarkDownloadFound(b *testing.B) {
	err := disk.WriteToStorage(strings.NewReader("test"), "123", 4)
	if err != nil {
		b.Error(err)
	}

	req, err := http.NewRequest("GET", "/sample/download/123", nil)
	if err != nil {
		b.Fatal(err)
	}

	handler := http.HandlerFunc(Download)

	r := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		handler.ServeHTTP(r, req)

		if r.Code != 200 {
			b.Errorf("Wrong status code %d", r.Code)
		}
	}

	fPath, err := disk.ConvertToStoragePath("123")
	if err != nil {
		b.Error(err)
	}

	os.Remove(fPath)
}


func TestUpload(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base("123"))
	if err != nil {
		t.Error(err)
	}
	_, err = io.Copy(part, strings.NewReader("test"))

	err = writer.Close()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/sample/upload/1234", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	handler := http.HandlerFunc(Upload)

	r := httptest.NewRecorder()
	handler.ServeHTTP(r, req)

	if r.Code != 201 {
		t.Errorf("Wrong status code %d", r.Code)
	}

	fPath, err := disk.ConvertToStoragePath("1234")
	if err != nil {
		t.Error(err)
	}

	os.Remove(fPath)
}

func BenchmarkUpload(b *testing.B) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base("123"))
	if err != nil {
		b.Error(err)
	}
	_, err = io.Copy(part, strings.NewReader("test"))

	err = writer.Close()
	if err != nil {
		b.Error(err)
	}

	req, err := http.NewRequest("POST", "/sample/upload/1234", body)
	if err != nil {
		b.Fatal(err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	handler := http.HandlerFunc(Upload)

	r := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		handler.ServeHTTP(r, req)

		if r.Code != 201 {
			b.Errorf("Wrong status code %d", r.Code)
		}
	}

	fPath, err := disk.ConvertToStoragePath("1234")
	if err != nil {
		b.Error(err)
	}

	os.Remove(fPath)
}