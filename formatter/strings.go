package formatter

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// InterfaceToString returns a string after being parsed by fmt.Sprintf
// I've noticed that it doesn't work on all structs passed in.
func InterfaceToString(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	return s
}

// StructToString marshalls a struct into bytes and then outputs them as a string
func StructToString(i interface{}) string {
	structBytes, err := json.Marshal(i)
	if err != nil {
		log.WithError(err).Error("unable to marhsal struct into bytes")
	}
	return string(structBytes)
}
