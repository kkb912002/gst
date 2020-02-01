package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type GstEventStruct C.GstEvent

type Event struct {
	GstEvent *GstEventStruct
}

func Eos() *Event {
	event := new(Event)
	event.GstEvent = (*GstEventStruct)(C.gst_event_new_eos())
	return event
}


