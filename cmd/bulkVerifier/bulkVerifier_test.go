package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestBulkVerify(t *testing.T) {
	inputFile, err := ioutil.TempFile("", "test_input_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file for input: %v", err)
	}
	defer os.Remove(inputFile.Name())

	testData := "user_agent,ip_address\nMozilla/5.0,66.249.66.1"
	if _, err := inputFile.WriteString(testData); err != nil {
		t.Fatalf("Failed to write to input temp file: %v", err)
	}
	inputFile.Close()

	// Create a temporary output file
	outputFile, err := ioutil.TempFile("", "test_output_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file for output: %v", err)
	}
	defer os.Remove(outputFile.Name())
	outputFile.Close() // Close the file so it can be opened by BulkVerify

	if err := BulkVerify(inputFile.Name(), outputFile.Name()); err != nil {
		t.Errorf("BulkVerify failed: %v", err)
	}

	outputContent, err := ioutil.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output temp file: %v", err)
	}

	expectedOutput := "user_agent,ip_address,bot_status,bot_name\nMozilla/5.0,66.249.66.1,unknown,\n"
	if !strings.Contains(string(outputContent), expectedOutput) {
		t.Errorf("Output file content did not match expected output. Got: %s Want: %s", string(outputContent), expectedOutput)
	}

}
