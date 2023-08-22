package testing

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff before the tests")
	exitVal := m.Run()
	log.Println("Do stuff after the tests")
	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}
