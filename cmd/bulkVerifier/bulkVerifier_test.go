package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestProcessCSV(t *testing.T) {
	input := strings.NewReader("user_agent,ip_address\nMozilla/5.0,66.249.66.1")
	var output bytes.Buffer

	err := processCSV(input, &output)
	if err != nil {
		t.Fatalf("processCSV failed: %v", err)
	}

	expectedOutput := "user_agent,ip_address,is_good_bot,bot_name\nMozilla/5.0,66.249.66.1,false,\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
