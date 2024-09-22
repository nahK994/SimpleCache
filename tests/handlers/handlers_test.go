package handler

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/handlers"
)

func TestHandler(t *testing.T) {
	t.Run("TestHandleGET", func(t *testing.T) {
		handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$5\r\nShomi\r\n")
		response := handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$4\r\nname\r\n")
		expected := "$5\r\nShomi\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleSET", func(t *testing.T) {
		response := handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$8\r\nlanguage\r\n$2\r\nGo\r\n")
		expected := "+OK\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleINCR", func(t *testing.T) {
		handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$7\r\ncounter\r\n$2\r\n10\r\n")
		response := handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$7\r\ncounter\r\n")
		expected := ":11\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleNotExistingINCR", func(t *testing.T) {
		response := handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$8\r\ncounter1\r\n")
		expected := ":1\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleDECR", func(t *testing.T) {
		handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$7\r\ncounter\r\n$2\r\n10\r\n")
		response := handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$7\r\ncounter\r\n")
		expected := ":9\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleNotExistingDECR", func(t *testing.T) {
		response := handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$8\r\ncounter2\r\n")
		expected := ":-1\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleDEL", func(t *testing.T) {
		handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$7\r\ntempKey\r\n$4\r\ntest\r\n")
		response := handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$7\r\ntempKey\r\n")
		expected := ":1\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleLPUSHAndLRANGE", func(t *testing.T) {
		handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$6\r\nmyList\r\n$3\r\none\r\n$3\r\ntwo\r\n$5\r\nthree\r\n")
		response := handlers.HandleCommand("*4\r\n$6\r\nLRANGE\r\n$6\r\nmyList\r\n$1\r\n1\r\n$2\r\n-1\r\n")
		expected := "*3\r\n$5\r\nthree\r\n$3\r\ntwo\r\n$3\r\none\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleLPOP", func(t *testing.T) {
		handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$5\r\nitems\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
		response := handlers.HandleCommand("*2\r\n$4\r\nLPOP\r\n$5\r\nitems\r\n")
		expected := "*1\r\n$1\r\nc\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleInvalidINCR", func(t *testing.T) {
		response := handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$4\r\nname\r\n")
		err := errors.Err{
			Type: errors.GetErrorTypes().TypeError,
		}
		expected := err.Error()
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleInvalidDECR", func(t *testing.T) {
		response := handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$4\r\nname\r\n")
		err := errors.Err{
			Type: errors.GetErrorTypes().TypeError,
		}
		expected := err.Error()
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})

	t.Run("TestHandleEmptyListLPOP", func(t *testing.T) {
		response := handlers.HandleCommand("*2\r\n$4\r\nLPOP\r\n$4\r\nname\r\n")
		expected := "*0\r\n"
		if response != expected {
			t.Errorf("Expected '%s', got '%s'", expected, response)
		}
	})
}
