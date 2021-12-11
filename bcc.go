package lrc

import (
	"encoding"
	"fmt"
	"hash"
)

// BCC implements the block check character hash, which is simply the
// XOR of all input bytes. The zero value is ready to use.
type BCC byte

// interface checks
var _ hash.Hash = (*BCC)(nil)
var _ encoding.BinaryMarshaler = BCC(0)
var _ encoding.BinaryUnmarshaler = (*BCC)(nil)

// BlockSize implements hash.Hash.BlockSize.
func (BCC) BlockSize() int {
	return 1
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (b BCC) MarshalBinary() ([]byte, error) {
	return []byte{byte(b)}, nil
}

// Reset implements hash.Hash.Reset.
func (b *BCC) Reset() {
	*b = 0
}

// Size implements hash.Hash.Size.
func (BCC) Size() int {
	return 1
}

// Sum implements hash.Hash.Sum.
func (b BCC) Sum(buf []byte) []byte {
	return append(buf, byte(b))
}

// Sum8 returns the 8-bit BCC hash accumulated so far.
func (b BCC) Sum8() uint8 {
	return uint8(b)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (b *BCC) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return fmt.Errorf("expected data size 1, got %d", len(data))
	}
	*b = BCC(data[0])
	return nil
}

// Write implements io.Writer.Write for hash.Hash.
func (b *BCC) Write(buf []byte) (int, error) {
	for _, v := range buf {
		*b ^= BCC(v)
	}
	return len(buf), nil
}
