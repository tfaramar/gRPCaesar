package main

import (
	"os"
	"testing"
) 


// test input validation function
func TestValidate( t *testing.T ) {
	//suppress stdout: https://stackoverflow.com/a/31663374
	sout := os.Stdout
	os.Stdout,_ = os.Open(os.DevNull)

	// test for empty text argument
	emptyText, ok := validate("", 2, "e")
	
	if emptyText != nil {
		t.Errorf("expected nil parsed struct for empty text input")
	}
	if ok {
		t.Errorf("expected ok to be false for empty text input")
	}
	
	// test for text with numeric characters/punctuation
	invalidText, ok := validate("I have 300 cats!", 2, "d")

	if invalidText != nil {
		t.Errorf("expected nil parsed struct for text input with numeric and/or punctuation characters")
	}
	if ok {
		t.Errorf("expected ok to be false for invalid text input")
	}

	// test for non e/d mode flag
	invalidMode, ok := validate("This is a test", 4, "z")

	if invalidMode != nil {
		t.Errorf("expected nil parsed struct for invalid mode input")
	}
	if ok {
		t.Errorf("expected ok to be false for invalid mode input")
	}

	// test for valid argument
	validArgs, ok := validate("hi", 2, "e")

	if validArgs.text != "hi" {
		t.Errorf("expected parsed struct text to be \"hi\"")
	}
	if validArgs.shift != 2 {
		t.Errorf("expected parsed struct shift to be 2")
	}
	if validArgs.encode != true {
		t.Errorf("expected parsed struct encode value to be true")
	}
	if !ok {
		t.Errorf("expected ok to be true for valid arguments")
	}
	
	os.Stdout = sout
}