package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeAppSuccess(t *testing.T) {

	err := initializeApp()

	if err != nil {
		t.Errorf("failed to initialize app: %v", err)
	}
	assert.NoError(t, err, "Expected no error")
}

// func TestInitializeAppFailed(t *testing.T) {

// 	err := initializeApp()

// 	if err != nil {
// 		t.Errorf("failed to initialize app: %v", err)
// 	}
// 	assert.Error(t, err, "Expected error")
// }

func TestMain(m *testing.M) {
	// old := os.Stdout

	// // Membuat pipe untuk menangkap output
	// r, w, _ := os.Pipe()
	// os.Stdout = w

	// // Menjalankan fungsi main
	// main()

	// // Mengembalikan os.Stdout
	// w.Close()
	// var buf bytes.Buffer
	// io.Copy(&buf, r)
	// os.Stdout = old

	// // Memeriksa output
	// expectedOutput := "Hello, World!\n"
	// actualOutput := buf.String()
	// if actualOutput != expectedOutput {
	// 	m.Error("Output tidak sesuai. Expected: %s, Actual: %s", expectedOutput, actualOutput)
	// }
	os.Exit(m.Run())
}
