package volume

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"volume/log"
)

// Peristent temp file name
var tmpFile string = filepath.Join(os.TempDir(), "sfmvolume.vol")

// Current level and lock
var (
	currentLock sync.Mutex
	current     uint8 = 30 // Default value set on start if not loaded
)

// Returns current volume level
func Current() uint8 {
	currentLock.Lock()
	defer currentLock.Unlock()
	return current
}

// Sets the volume level and saves the level
func SetVolume(level uint8) {
	currentLock.Lock()
	setVolume(level) // OS specific function call
	current = level
	currentLock.Unlock()
}

// Saves the current volume to a temporary file for persistence
func Save() error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, Current())
	if err != nil {
		return err
	}
	log.WithField("path", tmpFile).Debug("write current volume level to temp file")
	return ioutil.WriteFile(tmpFile, buf.Bytes(), 0644)
}

// Loads and sets the volume from a temporary file
func Load() {
	data, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		log.WithError(err).Error("error reading volume store")
		return
	}
	var level uint8
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &level); err != nil {
		log.WithError(err).Error("error reading volume store")
		return
	}
	SetVolume(level)
}
