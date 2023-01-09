package writer

import "codis2pika/internal/entry"

type Writer interface {
	Write(entry *entry.Entry)
}
