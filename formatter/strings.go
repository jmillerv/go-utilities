package formatter

import "fmt"

func InterfaceToString(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	return s
}