package volume

// #cgo LDFLAGS: -lasound -lstdc++
// #include <stdlib.h>
// #include "volume_linux.h"
import "C"

import (
	"math"
	"unsafe"
	"volume/log"
)

// Initialise Config
var config = Config{}

// Set volume level for device via ALSA
func SetVolume(level uint8) {
	card := C.CString(config.Card())
	mixer := C.CString(config.Mixer())
	max := config.Max()
	min := config.Min()
	volume := int(math.Floor((float64(level)*((max-min)/100) + min) + .5))
	log.WithField("level", volume).Debug("set volume level")
	C.setVolume(card, C.int(volume), mixer)
	C.free(unsafe.Pointer(card))
	setCurrent(level)
}
