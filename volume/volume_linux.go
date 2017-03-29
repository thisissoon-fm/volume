package volume

// #cgo LDFLAGS: -lasound -lstdc++
// #include <stdlib.h>
// #include "volume_linux.h"
import "C"

import (
	"math"
	"unsafe"
)

// Initialise Config
var config = Config{}

// Set volume level for device via ALSA
func SetVolume(level int) {
	card := C.CString(config.Card())
	mixer := C.CString(config.Mixer())
	max := config.Max()
	min := config.Min()
	volume := int(math.Floor((float64(level)*((max-min)/100) + min) + .5))
	C.setVolume(card, C.int(volume), mixer)
	C.free(unsafe.Pointer(card))
}
