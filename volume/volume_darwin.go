package volume

/*
#cgo CFLAGS: -v -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreAudio -framework AppKit
#include "volume_darwin.h"
inline float volume() {
	return (float)NSSound.systemVolume;
}
inline int defaultOutputDevice() {
	return (int)NSSound.defaultOutputDevice;
}
*/
import "C"

import (
	"volume/log"
)

func SetVolume(level int) {
	log.WithField("device", C.defaultOutputDevice()).Debug("device id")
	log.WithField("volume", C.volume()).Debug("device volume")
	log.WithField("level", level).Debug("set level darwin")
}
