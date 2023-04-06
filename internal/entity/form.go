package entity

import "mime/multipart"

type Form struct {
	Image       *multipart.FileHeader
	Title       string
	Description string
}
