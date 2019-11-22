package newlib

import "runtime"

//Version function exported
func Version() string {
	return runtime.Version()
}
