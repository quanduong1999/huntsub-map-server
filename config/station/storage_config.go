package station

import (
	"fmt"
	"os"
)

type StorageConfig struct {
	Upload      string
	MaxFileSize int64
}

const (
	oneGB = 1 << 30
	tenGB = 10 << 30
)

func (c *StorageConfig) Check() {
	if c.Upload == "" {
		c.Upload = "data/upload"
	}
	// child folder
	if c.MaxFileSize < oneGB {
		c.MaxFileSize = oneGB
	} else if c.MaxFileSize > tenGB {
		c.MaxFileSize = tenGB
	}
	// createFolder(c.Upload)
}

func createFolder(folder string) {
	err := os.MkdirAll(folder, os.ModeAppend)
	if err != nil {
		logger.Fatalf("Storage config: create folder [%v] failed", folder)
	}
}

func (c StorageConfig) String() string {
	return fmt.Sprintf("storage:upload=%s", c.Upload)
}
