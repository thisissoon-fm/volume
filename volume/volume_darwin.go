package volume

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreAudio -framework AppKit
#include "volume_darwin.h"
inline float volume() {
	return (float)NSSound.systemVolume;
}
inline int defaultOutputDevice() {
	return (int)NSSound.defaultOutputDevice;
}
inline void setVolume(float f) {
	return (void)[NSSound setSystemVolume:f];
}
*/
import "C"

import "volume/log"

// Sets the volume level
func SetVolume(level uint8) {
	if level > 100 {
		level = 100
	}
	flevel := (float32(level) / 100)
	log.WithField("level", flevel).Debug("set volume level")
	C.setVolume(C.float(flevel))
	setCurrent(level)
}
