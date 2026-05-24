package stack

import "testing"

func TestRegistryGetLatest(t *testing.T) {
	r := NewDefaultRegistry()

	got, ok := r.GetLatest("nextjs-simple")
	if !ok {
		t.Fatalf("expected default stack to exist")
	}

	if got.Version != "1.0.0" {
		t.Fatalf("expected latest version 1.0.0, got %q", got.Version)
	}
}

func TestRegistryVersionNormalization(t *testing.T) {
	r := &Registry{
		byKey:    map[string]Stack{},
		latestID: map[string]string{},
	}

	if err := r.Register(Stack{ID: "x", Version: "01.002.0003"}); err != nil {
		t.Fatalf("expected valid version, got error: %v", err)
	}

	_, ok := r.Get("x", "1.2.3")
	if !ok {
		t.Fatalf("expected normalized lookup to succeed")
	}
}

func TestRegistryRejectsInvalidVersion(t *testing.T) {
	r := &Registry{
		byKey:    map[string]Stack{},
		latestID: map[string]string{},
	}

	if err := r.Register(Stack{ID: "x", Version: "1.0"}); err == nil {
		t.Fatalf("expected invalid version to fail")
	}
}
