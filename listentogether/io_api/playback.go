package io_api

import (
	"context"
	// "fmt"

	"github.com/gen2brain/malgo"
)

// Playback streams samples from a reader to the sound device.
// The function initializes a playback device in the default context using
// provide stream configuration.
// Playback will commence playing the samples provided from the reader until either the
// reader returns an error, or the context signals done.
func Playback(ctx context.Context, r <-chan []byte, config StreamConfig) error {
	deviceConfig := config.asDeviceConfig(malgo.Playback)
	abortChan := make(chan error)
	defer close(abortChan)
	aborted := false

	deviceCallbacks := malgo.DeviceCallbacks{
		Data: func(outputSamples, nil []byte, frameCount uint32) {
			if aborted {
				return
			}
			if frameCount == 0 {
				return
			}
			
			copy(outputSamples, <-r)
			// fmt.Println("PlayBack: ", outputSamples)
		},
	}

	return stream(ctx, abortChan, deviceConfig, deviceCallbacks)
}
