package gotemp_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/bllyanos/gotemp"
)

func TestNew(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if g == nil {
		t.Fatal("expected non-nil Gotemp instance")
	}
}

func TestRenderPage(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var buf bytes.Buffer
	err = g.RenderPage(&buf, "app_layout", "home/index.html", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result := buf.String()
	if !strings.Contains(result, "<!DOCTYPE html>") {
		t.Error("expected HTML doctype in output")
	}
	if !strings.Contains(result, "Homepage") {
		t.Error("expected 'Homepage' in output")
	}
	if !strings.Contains(result, "Your APP!!") {
		t.Error("expected header content in output")
	}
}

func TestRenderPageWithData(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data := map[string]string{
		"Title": "Test Page",
		"Body":  "Test content",
	}

	var buf bytes.Buffer
	err = g.RenderPage(&buf, "app_layout", "home/index.html", data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result := buf.String()
	if !strings.Contains(result, "<!DOCTYPE html>") {
		t.Error("expected HTML doctype in output")
	}
}

func TestNonExistentPage(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var buf bytes.Buffer
	err = g.RenderPage(&buf, "app_layout", "nonexistent/page.html", nil)
	if err == nil {
		t.Error("expected error for non-existent page")
	}
}

func TestNonExistentLayout(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var buf bytes.Buffer
	err = g.RenderPage(&buf, "nonexistent_layout", "home/index.html", nil)
	if err == nil {
		t.Error("expected error for non-existent layout")
	}
}

func TestMultiplePages(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var buf1, buf2 bytes.Buffer
	err1 := g.RenderPage(&buf1, "app_layout", "home/index.html", nil)
	err2 := g.RenderPage(&buf2, "app_layout", "home/index.html", nil)

	if err1 != nil || err2 != nil {
		t.Fatalf("expected no errors, got %v and %v", err1, err2)
	}

	result1 := buf1.String()
	result2 := buf2.String()

	if result1 != result2 {
		t.Error("expected same output for same page rendered twice")
	}
}

func TestRenderToStdout(t *testing.T) {
	g, err := gotemp.New("examples")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = g.RenderPage(os.Stdout, "app_layout", "home/index.html", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	result := buf.String()

	if !strings.Contains(result, "Homepage") {
		t.Error("expected 'Homepage' in stdout output")
	}
}
