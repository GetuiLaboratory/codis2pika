package reader

import "codis2pika/internal/entry"

type Reader interface {
	StartRead() chan *entry.Entry
}
