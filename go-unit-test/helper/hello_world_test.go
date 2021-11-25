package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//only run on current package
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("TestHelloWorldSubTest1", func(t *testing.T) {
		result := HelloWorld("Agus")

		if result != "Hello Agus" {
			panic("Result is not Hello Agus")
		}
		fmt.Println("TestHelloWorld done")
	})

	t.Run("TestHelloWorldSubTest2", func(t *testing.T) {
		result := HelloWorld("Agus")

		if result != "Hello Agus" {
			panic("Result is not Hello Agus")
		}
		fmt.Println("TestHelloWorld done")
	})
}

func TestTableHelloWorld(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "TestHelloWorld(Agus)",
			request:  "Agus",
			expected: "Hello Agus",
		},
		{
			name:     "TestHelloWorld(Megha)",
			request:  "Megha",
			expected: "Hello Megha",
		},
		{
			name:     "TestHelloWorld(agus)",
			request:  "agus",
			expected: "Hello agus",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Agus")

	if result != "Hello Agus" {
		panic("Result is not Hello Agus")
	}
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("agus")

	if result != "hello agus" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldFail done")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("agus")

	if result != "hello agus" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldFailNow done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("agus")

	if result != "hello agus" {
		t.Error("result is not hello agus")
	}
	fmt.Println("TestHelloWorldError done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("agus")

	if result != "hello agus" {
		t.Fatal("result is not hello agus")
	}
	fmt.Println("should not execute")
}

func TestHelloWorldAsertion(t *testing.T) {
	result := HelloWorld("Agus")
	assert.Equal(t, "Hello Agus", result, "result must be 'Hello Agus'") //jika gagal otomatis bakal manggil fail

	fmt.Println("TestHelloWorldAsertion done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Agus")
	require.Equal(t, "Hello agus", result, "result must be 'Hello Agus'") //jika gagal otomatis bakal manggil fail

	fmt.Println("should not execute")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Test only run on MAC OS")
	}

	result := HelloWorld("Agus")
	require.Equal(t, "Hello Agus", result, "result must be 'Hello Agus'") //jika gagal otomatis bakal manggil fail

	fmt.Println("should not execute")
}

//run in curent folder
//go test -v (run all test)
//go test -v -run=TestHelloWorld (run test by prefix)

//go test -v -run=TestHelloWorld/TestHelloWorldSubTest1 (filter subtest)
//go test -v -run=/TestHelloWorldSubTest2 (filter subtest)

//run all test recursive
//go test -v ./...

//mengagalkan test
//t.Fail() mengagalkan test namun tetap melanjutkan eksekusi unit test / melanjutkan kode program dlm fungsi tsb
//t.FailNow() mengagalkan unit test saat ini tanpa melanjutkan eksekusi unit test
