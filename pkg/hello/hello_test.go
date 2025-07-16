// hello_test.go
package hello_test

import (
	"fmt"
	"testing"

	"github.com/yuyan42517/hello/pkg/hello"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := hello.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}

}
