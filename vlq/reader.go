package vlq

import (
	"errors"
	"fmt"
	"io"
)

// The maximum allowed VLQ value as defined by the spec
const MAX = 0x0FFFFFFF

// ReadVLQ reads a single VLQ value from a bytes.Reader
func ReadVLQ(r io.Reader) (n int, read int, err error) {

	mada := true
	for mada {
		// read the next byte
		buff := []byte{0}
		_, err = io.ReadFull(r, buff)
		if err != nil {
			return 0, read, err
		}
		b := buff[0]

		// increment read bytes counter
		read++

		// add the 7 LSBs to the result
		n ^= int(b & 0x7f)

		// if the MSB is 1, prepare for the next iteration
		if mada = 1 == b>>7; mada {
			n <<= 7
		}

		// check for exceeding max value defined in spec
		if n > MAX {
			return 0, read, errors.New(fmt.Sprintf("exceeded maximum vlq value [%d]", MAX))
		}
	}

	return
}
