package model

import "time"

type PinnedFile struct {
	IpfsHash     string
	PinSize      int8
	Timestamp    time.Time
	IsDuplicated *bool
}
