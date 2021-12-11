// Package lrc implements two simple hash algorithms, confusingly known under
// the same name, Longitudinal Redundancy Check (LRC). The hash sum of both
// algorithms is just a single byte.
//
// The first algorithm, also known as BCC (for block check character)
// calculates its single sum byte simply as the XOR of all input bytes.
// This is the algorithm defined in ISO 1155.
//
// The second algorithm calculates its single sum byte as the 2's complement
// of the arithmetic sum (mod 2⁸) of the input bytes. This is the algorithm
// used by the popular Modbus standard, as specified in Modbus over Serial Line,
// section 6.2.1.
//
// The BCC and LRC types in this package implement the hash.Hash,
// encoding.BinaryMarshaler and encoding.BinaryUnmarshaler interfaces.
package lrc
