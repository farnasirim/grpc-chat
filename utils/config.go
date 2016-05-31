package utils

import (
	"fmt"
	"os"
)

func ConfigString(name string) string {
	return os.Getenv(fmt.Sprintf("CHAT_%s", name))
}
