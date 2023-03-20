package model

import "time"

type PinnedFile struct {
	IpfsHash     string
	PinSize      int64
	Timestamp    time.Time
	IsDuplicated *bool
}
