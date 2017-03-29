package hub

import (
	"bytes"
	"io/ioutil"
	"time"

	"volume/event"
	"volume/log"
	"volume/volume"
)

// Handles updating the volume level
func updateVolumeHandler(e event.Event, c Client) error {
	// Decode payload
	payload := event.VolumeLevelPayload{}
	decoder := event.NewDecoder(bytes.NewReader(e.Payload))
	if err := decoder.Decode(&payload); err != nil {
		return err
	}
	log.WithField("level", payload.Level).Debug("handle volume update event")
	// Set volume level
	volume.SetVolume(payload.Level)
	// Encode increased event payload
	var buff = new(bytes.Buffer)
	encoder := event.NewEncoder(buff)
	err := encoder.Encode(&event.Event{
		Topic:   event.VolumeUpdatedTopic,
		Created: time.Now().UTC(),
		Payload: e.Payload,
	})
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(buff)
	if err != nil {
		return err
	}
	if _, err := c.Write(data); err != nil {
		defer Remove(c)
	}
	return nil
}

// Handles mute
func muteHandler(e event.Event, c Client) error {
	log.Debug("handle mute volume event")
	volume.Mute() // Mute the volume
	// Send volume:muted event
	var buff = new(bytes.Buffer)
	encoder := event.NewEncoder(buff)
	err := encoder.Encode(&event.Event{
		Topic:   event.VolumeMutedTopic,
		Created: time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(buff)
	if err != nil {
		return err
	}
	if _, err := c.Write(data); err != nil {
		defer Remove(c)
	}
	return nil
}

// Handles unmute
func unmuteHandler(e event.Event, c Client) error {
	log.Debug("handle un mute volume event")
	volume.UnMute() // UnMute the volume
	// Send volume:unmuted event
	var buff = new(bytes.Buffer)
	encoder := event.NewEncoder(buff)
	err := encoder.Encode(&event.Event{
		Topic:   event.VolumeUnmutedTopic,
		Created: time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(buff)
	if err != nil {
		return err
	}
	if _, err := c.Write(data); err != nil {
		defer Remove(c)
	}
	return nil
}
