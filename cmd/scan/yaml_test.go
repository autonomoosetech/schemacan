package main

import (
	"testing"
)

func TestSplitStreamLength(t *testing.T) {
	out := splitStream([]byte("a:1\n---\nb:2\n---\nc:3\n"))

	length := len(out)

	if length != 3 {
		t.Errorf("Stream splitting got %d; want 3", length)
	}
}

func TestSplitStreamNotRequired(t *testing.T) {
	out := splitStream([]byte("a:1\nb:2\nc:2\n"))

	length := len(out)

	if length != 1 {
		t.Errorf("Stream splitting got %d; want 1", length)
	}
}

func TestUnmarshalLength(t *testing.T) {
	example := []byte(`
version: v1
kind: slot
metadata:
  name: example-a
  namespace: default
spec:
  min: 0
  max: 100
  offset: 50
  size: 8
---
version: v1
kind: slot
metadata:
  name: example-b
  namespace: default
spec:
  min: 0
  max: 100
  offset: 50
  size: 8
`)

	out, err := Unmarshal(example)

	if err != nil {
		t.Errorf("unmarshal failed: %v", err)
		return
	}

	length := len(out)

	if length != 2 {
		t.Errorf("Stream splitting got %d; want 2", length)
	}
}
