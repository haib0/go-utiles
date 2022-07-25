package io_api

import (
	"context"
	// "os"

	"github.com/gen2brain/malgo"
)

// Capture records incoming samples into the provided writer.
// The function initializes a capture device in the default context using
// provide stream configuration.
// Capturing will commence writing the samples to the writer until either the
// writer returns an error, or the context signals done.
func Capture(ctx context.Context, w chan<- []byte, config StreamConfig) error {
	deviceConfig := config.asDeviceConfig(malgo.Loopback)
	abortChan := make(chan error)
	defer close(abortChan)
	aborted := false

	deviceCallbacks := malgo.DeviceCallbacks{
		Data: func(outputSamples, inputSamples []byte, frameCount uint32) {
			if aborted {
				return
			}

			w <- inputSamples
			// os.WriteFile("capture2", inputSamples, os.ModeAppend)
		},
	}

	return stream(ctx, abortChan, deviceConfig, deviceCallbacks)
}

