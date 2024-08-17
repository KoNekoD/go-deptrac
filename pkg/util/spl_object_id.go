package util

import "fmt"

func SplObjectID(v interface{}) string {
	return fmt.Sprintf("%p", v)
}
