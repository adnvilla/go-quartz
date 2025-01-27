package quartz_test

import (
	"reflect"
	"testing"

	"github.com/adnvilla/go-quartz/quartz"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v != %v", a, b)
	}
}

func assertEqualInt64(t *testing.T, a int64, b int64) {
	if a != b {
		t.Fatalf("%d != %d", a, b)
	}
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		t.Fatalf("%v == %v", a, b)
	}
}

func TestUtils(t *testing.T) {
	hash := quartz.HashCode("foo")
	assertEqual(t, hash, 2851307223)
}
