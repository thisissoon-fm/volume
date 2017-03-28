package volume

import "volume/log"

func SetVolume(level int) {
	log.WithField("level", level).Warn("set volume for darwin OS's is not supported")
}
