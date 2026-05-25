package cmd

import "testing"

func TestInitCommandRequiresName(t *testing.T) {
	rootCmd.SetArgs([]string{"init"})
	_, err := rootCmd.ExecuteC()
	if err == nil {
		t.Fatalf("expected missing --name to fail")
	}
}

func TestInitCommandRejectsUnknownStack(t *testing.T) {
	rootCmd.SetArgs([]string{"init", "--name", "demo", "--stack", "unknown"})
	_, err := rootCmd.ExecuteC()
	if err == nil {
		t.Fatalf("expected unknown stack to fail")
	}
}
