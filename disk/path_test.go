package disk

import (
	"os"
	"strings"
	"testing"
)

func TestRename(t *testing.T) {
	err := Rename("123", "1234")
	if err == nil {
		t.Error("The non found file has been renamed")
	}

	err = WriteToStorage(strings.NewReader("test"), "test123", 4)
	if err != nil {
		t.Error(err)
	}

	err = Rename("test", "1234")
	if err == nil {
		t.Error("The non found file has been renamed")
	}

	fPath, err := ConvertToStoragePath("test")
	if err != nil {
		t.Error(err)
	}

	fPath2, err := ConvertToStoragePath("1234")
	if err != nil {
		t.Error(err)
	}

	os.Remove(fPath)
	os.Remove(fPath2)
}

func TestWriteToStorage(t *testing.T) {
	err := WriteToStorage(nil, "", 0)
	if err == nil {
		t.Error("The null input has been written to empty filename")
	}

	err = WriteToStorage(strings.NewReader("test"), "", 0)
	if err == nil {
		t.Error("The proper input has been written to empty filename")
	}

	err = WriteToStorage(strings.NewReader("test"), "test123", 30)
	if err == nil {
		t.Error("The proper input has been written to proper filename with wrong size")
	}

	err = WriteToStorage(strings.NewReader("test"), "test123", 4)
	if err != nil {
		t.Error(err)
	}

	fPath, err := ConvertToStoragePath("test123")
	if err != nil {
		t.Error(err)
	}

	os.Remove(fPath)
}

func TestExists(t *testing.T) {
	b, err := Exists("123test")
	if b {
		t.Error(err)
	}

	err = WriteToStorage(strings.NewReader("test"), "123test", 4)
	if err != nil {
		t.Error(err)
	}

	b, err = Exists("123test")
	if !b {
		t.Error(err)
	}

	fPath, err := ConvertToStoragePath("123test")
	if err != nil {
		t.Error(err)
	}

	os.Remove(fPath)
}

func TestStorageMap(t *testing.T) {
	counter := make(map[string]int)
	for _, v := range storageMap {
		val, ok := counter[v]
		if !ok {
			counter[v] = 1
		} else {
			counter[v] = val + 1
		}
	}

	for k, v := range counter {
		if v != 2 {
			t.Errorf("%s has wrong size of element %d", k, v)
		}
	}
}

func BenchmarkGetRandomExt(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		string1 := GetRandomExt()
		string2 := GetRandomExt()
		if string1 == string2 {
			b.Errorf("The random strings are the same 1st: %s 2nd: %s", string1, string2)
		}
	}
}
