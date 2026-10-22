package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "example-alfred")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

func runBinary(t *testing.T, bin string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	code := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("running binary: %v", err)
	}
	return outBuf.String(), errBuf.String(), code
}

type scriptFilterItem struct {
	Title string `json:"title"`
	Arg   string `json:"arg"`
	Valid *bool  `json:"valid"`
}

type scriptFilterResponse struct {
	Items []scriptFilterItem `json:"items"`
}

func TestNoArgsPrintsAllShortcuts(t *testing.T) {
	bin := buildBinary(t)

	stdout, stderr, code := runBinary(t, bin)
	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %s", code, stderr)
	}

	var resp scriptFilterResponse
	if err := json.Unmarshal([]byte(stdout), &resp); err != nil {
		t.Fatalf("Unmarshal(%q): %v", stdout, err)
	}
	if len(resp.Items) != 3 {
		t.Fatalf("got %d items, want 3", len(resp.Items))
	}
}

func TestQueryFiltersShortcuts(t *testing.T) {
	bin := buildBinary(t)

	stdout, stderr, code := runBinary(t, bin, "doc")
	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %s", code, stderr)
	}

	var resp scriptFilterResponse
	if err := json.Unmarshal([]byte(stdout), &resp); err != nil {
		t.Fatalf("Unmarshal(%q): %v", stdout, err)
	}
	if len(resp.Items) != 1 || resp.Items[0].Title != "docs" {
		t.Fatalf("items = %+v, want a single \"docs\" item", resp.Items)
	}
}
