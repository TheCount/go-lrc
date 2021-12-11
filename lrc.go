package lrc

import (
	"encoding"
	"fmt"
	"hash"
)

// LRC implements the Modbus longitudinal redundancy check hash, which is
// simply the 2's complements of the sum (mod 2⁸) of all input bytes.
// The zero value is ready to use.
type LRC byte

// interface checks
var _ hash.Hash = (*LRC)(nil)
var _ encoding.BinaryMarshaler = LRC(0)
var _ encoding.BinaryUnmarshaler = (*LRC)(nil)

// BlockSize implements hash.Hash.BlockSize.
func (LRC) BlockSize() int {
	return 1
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (b LRC) MarshalBinary() ([]byte, error) {
	return []byte{byte(b)}, nil
}

// Reset implements hash.Hash.Reset.
func (b *LRC) Reset() {
	*b = 0
}

// Size implements hash.Hash.Size.
func (LRC) Size() int {
	return 1
}

// Sum implements hash.Hash.Sum.
func (b LRC) Sum(buf []byte) []byte {
	return append(buf, b.Sum8())
}

// Sum8 returns the 8-bit LRC hash accumulated so far.
func (b LRC) Sum8() uint8 {
	return uint8(-int8(b))
}

// HexSum appends the current hash to buf in hexadecimal format and returns the
// resulting slice. It does not change the underlying hash state.
// The hexadecimal representation is high-nibble first, using upper case
// hexadecimal digits. This format conforms to the specification of the ASCII
// transmission mode in the Modbus over serial line specification
// (section 2.5.2).
func (b LRC) HexSum(buf []byte) []byte {
	const hexdigits = "0123456789ABCDEF"
	sum := b.Sum8()
	return append(buf, hexdigits[sum>>4], hexdigits[sum&0x0F])
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (b *LRC) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return fmt.Errorf("expected data size 1, got %d", len(data))
	}
	*b = LRC(data[0])
	return nil
}

// Write implements io.Writer.Write for hash.Hash.
func (b *LRC) Write(buf []byte) (int, error) {
	for _, v := range buf {
		*b += LRC(v)
	}
	return len(buf), nil
}
