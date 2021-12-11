package lrc

import (
	"bytes"
	"testing"
)

// TestLRC tests LRC hashing.
func TestLRC(t *testing.T) {
	const hello = "Hello, world!"
	testString := []byte(hello)
	var h LRC
	n, err := h.Write(testString)
	if err != nil {
		t.Fatalf("unexpected write error: %s", err)
	}
	if n != len(testString) {
		t.Fatalf("expected write length=%d, got %d", len(testString), n)
	}
	sum8 := h.Sum8()
	if sum8 != 0x77 {
		t.Fatalf("expected lrc=0x77, got 0x%02x", sum8)
	}
	testString = h.Sum(testString)
	if len(testString) != n+1 {
		t.Errorf("expected lrc'd length=%d, got %d", n+1, len(testString))
	}
	if !bytes.Equal([]byte(hello), testString[:n]) {
		t.Errorf("hash input got altered: %s", testString[:n])
	}
	if testString[n] != 0x77 {
		t.Errorf("expected appended lrc=0x77, got 0x%02x", testString[n])
	}
	testString = h.HexSum(testString[:n])
	if len(testString) != n+2 {
		t.Errorf("expected lrc'd length=%d, got %d", n+2, len(testString))
	}
	if !bytes.Equal([]byte(hello), testString[:n]) {
		t.Errorf("hash input got altered: %s", testString[:n])
	}
	if testString[n] != '7' && testString[n+1] != '7' {
		t.Errorf("expected appended lrc=77, got %s", testString[n:])
	}
	h.Reset()
	if h.Sum8() != 0 {
		t.Errorf("expected lrc=0 after reset, got 0x%02x", h.Sum8())
	}
}

// TestLRCBlockSize tests the LRC.BlockSize method.
func TestLRCBlockSize(t *testing.T) {
	var h LRC
	if h.BlockSize() != 1 {
		t.Errorf("expected blocksize 1, got %d", h.BlockSize())
	}
}

// TestLRCMarshal tests (un-)marshalling a hash.
func TestLRCMarshal(t *testing.T) {
	const (
		hello1 = "Hello, "
		hello2 = "world!"
		hello  = hello1 + hello2
	)
	var h, h1, h2 LRC
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
		t.Errorf("expected lrc=0x%02x, got 0x%02x", h.Sum8(), h2.Sum8())
	}
}

// TestLRCSize tests the LRC.Size method.
func TestLRCSize(t *testing.T) {
	var h LRC
	if h.Size() != 1 {
		t.Errorf("expected hash size 1, got %d", h.Size())
	}
}
