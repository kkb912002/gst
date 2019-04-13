// Bindings for GStreamer API
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

#cgo pkg-config: gstreamer-1.0
*/
import "C"

var CLOCK_TIME_NONE = int64(C.GST_CLOCK_TIME_NONE)
