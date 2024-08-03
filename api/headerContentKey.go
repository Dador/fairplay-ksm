package main

import (
	"encoding/base64"
	"errors"

	"github.com/Dador/fairplay-ksm/ksm"
)

type HeaderContentKey struct {
	Key string
	Iv  string
}

// Implement FetchContentKey func
func (h HeaderContentKey) FetchContentKey(assetID []byte) ([]byte, []byte, error) {
	if len(h.Key) == 0 || len(h.Iv) == 0 {
		return nil, nil, errors.New("key or iv not provided in headers")
	}
	key, err := base64.StdEncoding.DecodeString(h.Key)
	if err != nil {
		return nil, nil, err
	}
	iv, err := base64.StdEncoding.DecodeString(h.Iv)
	if err != nil {
		return nil, nil, err
	}
	return key, iv, nil
}

// Implement FetchContentKeyDuration func
func (HeaderContentKey) FetchContentKeyDuration(assetID []byte) (*ksm.CkcContentKeyDurationBlock, error) {
	// Unlimited duration by default
	LeaseDuration := uint32(0xFFFFFFFF)  // The duration of the lease, if any, in seconds.
	RentalDuration := uint32(0xFFFFFFFF) // The duration of the rental, if any, in seconds.

	return ksm.NewCkcContentKeyDurationBlock(LeaseDuration, RentalDuration), nil
}
