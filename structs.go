package golibknot

/*
 #cgo LDFLAGS: -lknot
 #cgo CFLAGS: -I /usr/include/libknot/
 #include <dlfcn.h>
 #include <libknot.h>
*/
import "C"

// Representation of the types of messages received / sent to the socket as int
const (
	END   = iota
	DATA  = iota
	EXTRA = iota
	BLOCK = iota
)

// Representation of data received / sent by the socket
type KnotCtlData struct {
	Command, Flags, Error, Section, Item, Id, Zone, Owner, Ttl, Type, Data, Filter string
}

// ToCtl converts the structure to a structure used to send data to the C library (C.knot_ctl_data_t).
func (obj KnotCtlData) ToCtl() C.knot_ctl_data_t {
	var ctlStruct C.knot_ctl_data_t
	objs := []string{obj.Command, obj.Flags, obj.Error, obj.Section, obj.Item, obj.Id, obj.Zone, obj.Owner, obj.Ttl, obj.Type, obj.Data, obj.Filter}
	for k, v := range objs {
		ctlStruct[k] = C.CString(v)
	}
	return ctlStruct
}

// FromCtl transforms the data received by the C library into a struct accessible in Go
func (obj *KnotCtlData) FromCtl(data C.knot_ctl_data_t) {
	objs := []*string{&obj.Command, &obj.Flags, &obj.Error, &obj.Section, &obj.Item, &obj.Id, &obj.Zone, &obj.Owner, &obj.Ttl, &obj.Type, &obj.Data, &obj.Filter}
	for i, v := range objs {
		*v = C.GoString(data[i])
	}
}
