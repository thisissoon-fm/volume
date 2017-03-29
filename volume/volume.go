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

// State persisntence
var (
	volTmpFile  string = filepath.Join(os.TempDir(), "sfmvolume.vol")
	muteTmpFile string = filepath.Join(os.TempDir(), "sfmvolume.mute")
)

// State
var (
	lock    sync.Mutex
	current uint8 = 30 // Default value set on start if not loaded
	muted   bool
)

// Load volume state
func loadVolumeState() (level uint8, err error) {
	var data []byte
	data, err = ioutil.ReadFile(volTmpFile)
	if err != nil {
		return
	}
	err = binary.Read(bytes.NewReader(data), binary.LittleEndian, &level)
	return
}

// Save volume state
func saveVolumeState() error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, Current())
	if err != nil {
		return err
	}
	log.WithField("path", volTmpFile).Debug("write current volume level to temp file")
	return ioutil.WriteFile(volTmpFile, buf.Bytes(), 0644)
}

// Load mute state
func loadMuteState() (muted bool, err error) {
	var data []byte
	data, err = ioutil.ReadFile(muteTmpFile)
	if err != nil {
		return
	}
	err = binary.Read(bytes.NewReader(data), binary.LittleEndian, &muted)
	return
}

// Save mute state
func saveMuteState() error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, Muted())
	if err != nil {
		return err
	}
	log.WithField("path", muteTmpFile).Debug("write current mute state to temp file")
	return ioutil.WriteFile(muteTmpFile, buf.Bytes(), 0644)
}

// Sets the volume level and saves the level
func SetVolume(level uint8) {
	lock.Lock()
	if !muted {
		setVolume(level) // OS specific function call
	}
	current = level
	lock.Unlock()
}

// Mutes the player, effectivly setting volume to 0
func Mute() {
	lock.Lock()
	if !muted {
		setVolume(0) // OS specific function call
		muted = true
	}
	lock.Unlock()
}

// Unmutes the player, restoring the current volume level
func UnMute() {
	lock.Lock()
	if muted {
		setVolume(current)
		muted = false
	}
	lock.Unlock()
}

// Returns muted state
func Muted() bool {
	lock.Lock()
	defer lock.Unlock()
	return muted
}

// Returns current level state
func Current() uint8 {
	lock.Lock()
	defer lock.Unlock()
	return current
}

// Saves the current volume to a temporary file for persistence
func SaveSate() {
	if err := saveVolumeState(); err != nil {
		log.WithError(err).Error("error saving volume state")
	}
	if err := saveMuteState(); err != nil {
		log.WithError(err).Error("error saving mute state")
	}
}

// Loads and sets the volume from a temporary file
func LoadState() {
	mute, err := loadMuteState()
	if err != nil {
		log.WithError(err).Error("error loading volume state")
		return
	}
	if mute {
		Mute()
	}
	level, err := loadVolumeState()
	if err != nil {
		log.WithError(err).Error("error loading volume state")
		return
	}
	SetVolume(level)
}
