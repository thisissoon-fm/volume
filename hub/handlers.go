package hub

import (
	"bytes"
	"io/ioutil"
	"time"

	"volume/event"
	"volume/log"
)

// Handles a increase volume event
func increaseVolHandler(e event.Event, c Client) error {
	// Decode payload
	payload := event.VolumeLevelPayload{}
	decoder := event.NewDecoder(bytes.NewReader(e.Payload))
	if err := decoder.Decode(&payload); err != nil {
		return err
	}
	log.WithField("level", payload.Level).Debug("increase volume")
	// TODO: change volume level
	// Encode increased event payload
	var buff = new(bytes.Buffer)
	encoder := event.NewEncoder(buff)
	err := encoder.Encode(&event.Event{
		Topic:   event.VolumeIncreasedTopic,
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
	if _, err := rw.Write(data); err != nil {
		defer Remove(rw)
	}
	return nil
}

// Handles a decrease volume event
func decreaseVolHandler(e event.Event, c Client) error {
	// Decode payload
	payload := event.VolumeLevelPayload{}
	decoder := event.NewDecoder(bytes.NewReader(e.Payload))
	if err := decoder.Decode(&payload); err != nil {
		return err
	}
	log.WithField("level", payload.Level).Debug("decrease volume")
	// TODO: change volume level
	// Encode increased event payload
	var buff = new(bytes.Buffer)
	encoder := event.NewEncoder(buff)
	err := encoder.Encode(&event.Event{
		Topic:   event.VolumeDecreasedTopic,
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
		defer Remove(rw)
	}
	return nil
}
