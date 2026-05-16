package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewCode(prefix string) string {
	return fmt.Sprintf("%s%s%s", prefix, time.Now().Format("20060102150405"), uuid.NewString()[:8])
}
