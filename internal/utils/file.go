package utils

import (
	"codis2pika/internal/log"
	"os"
)

func DoesFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			log.PanicError(err)
		}
	}
	return true
}
