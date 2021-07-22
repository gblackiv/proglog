package log

import {
	"bufio"
	"encoding/binary"
	"os"
	"sync"
}

	var {
		enc = binary.BigEndian
	}

	const {
		lenWidth = 8
	}

	type store struct {
		*os.File
	}
