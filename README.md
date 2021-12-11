This Go package implements two simple hash algorithms, confusingly known under the same name, Longitudinal Redundancy Check (LRC). The hash sum of both algorithms is just a single byte.

The first algorithm, also known as BCC (for block check character) calculates its single sum byte simply as the XOR of all input bytes. This is the algorithm defined in [ISO 1155](https://www.iso.org/standard/5723.html).

The second algorithm calculates its single sum byte as the 2's complement of the arithmetic sum (mod 2⁸) of the input bytes. This is the algorithm used by the popular Modbus standard, as specified in [Modbus over Serial Line](https://modbus.org/docs/Modbus_over_serial_line_V1_02.pdf), section 6.2.1.
