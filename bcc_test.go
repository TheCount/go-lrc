package lrc

import (
	"bytes"
	"testing"
)

// TestBCC tests BCC hashing.
func TestBCC(t *testing.T) {
	const hello = "Hello, world!"
	testString := []byte(hello)
	var h BCC
	n, err := h.Write(testString)
	if err != nil {
		t.Fatalf("unexpected write error: %s", err)
	}
	if n != len(testString) {
		t.Fatalf("expected write length=%d, got %d", len(testString), n)
	}
	sum8 := h.Sum8()
	if sum8 != 0x0d {
		t.Fatalf("expected bcc=0x0d, got 0x%02x", sum8)
	}
	testString = h.Sum(testString)
	if len(testString) != n+1 {
		t.Errorf("expected bcc'd length=%d, got %d", n+1, len(testString))
	}
	if !bytes.Equal([]byte(hello), testString[:n]) {
		t.Errorf("hash input got altered: %s", testString[:n])
	}
	if testString[n] != 0x0d {
		t.Errorf("expected appended bcc=0x0d, got 0x%02x", testString[n])
	}
	h.Reset()
	if h.Sum8() != 0 {
		t.Errorf("expected bcc=0 after reset, got 0x%02x", h.Sum8())
	}
}

// TestBCCBlockSize tests the BCC.BlockSize method.
func TestBCCBlockSize(t *testing.T) {
	var h BCC
	if h.BlockSize() != 1 {
		t.Errorf("expected blocksize 1, got %d", h.BlockSize())
	}
}

// TestBCCMarshal tests (un-)marshalling a hash.
func TestBCCMarshal(t *testing.T) {
	const (
		hello1 = "Hello, "
		hello2 = "world!"
		hello  = hello1 + hello2
	)
	var h, h1, h2 BCC
	if _, err := h.Write([]byte(hello)); err != nil {
		t.Fatalf("write full message: %s", err)
	}
	if _, err := h1.Write([]byte(hello1)); err != nil {
		t.Fatalf("write first part of message: %s", err)
	}
	b, err := h1.MarshalBinary()
	if err != nil {
		t.Fatalf("marshal: %s", err)
	}
	b = append(b, 0x42)
	if err = h2.UnmarshalBinary(b); err == nil {
		t.Error("expected error after appending bogus data")
	}
	b = b[:len(b)-1]
	if err = h2.UnmarshalBinary(b); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if _, err = h2.Write([]byte(hello2)); err != nil {
		t.Fatalf("write second part of message: %s", err)
	}
	if h.Sum8() != h2.Sum8() {
		t.Errorf("expected bcc=0x%02x, got 0x%02x", h.Sum8(), h2.Sum8())
	}
}

// TestBCCSize tests the BCC.Size method.
func TestBCCSize(t *testing.T) {
	var h BCC
	if h.Size() != 1 {
		t.Errorf("expected hash size 1, got %d", h.Size())
	}
}
