package golibknot

/*
 #cgo LDFLAGS: -lknot
 #cgo CFLAGS: -I /usr/include/libknot/
 #include <dlfcn.h>
 #include <libknot.h>
*/
import "C"

type KnotCtlData struct {
	Command, Flags, Error, Section, Item, Id, Zone, Owner, Ttl, Type, Data, Filter string
}

func (obj KnotCtlData) ToCtl() C.knot_ctl_data_t {
	var ctlStruct C.knot_ctl_data_t
	objs := []string{obj.Command, obj.Flags, obj.Error, obj.Section, obj.Item, obj.Id, obj.Zone, obj.Owner, obj.Ttl, obj.Type, obj.Data, obj.Filter}
	for k, v := range objs {
		ctlStruct[k] = C.CString(v)
	}
	return ctlStruct
}

func (obj *KnotCtlData) FromCtl(data C.knot_ctl_data_t) {
	objs := []*string{&obj.Command, &obj.Flags, &obj.Error, &obj.Section, &obj.Item, &obj.Id, &obj.Zone, &obj.Owner, &obj.Ttl, &obj.Type, &obj.Data, &obj.Filter}
	for i, v := range objs {
		*v = C.GoString(data[i])
	}
}

const (
	END   = iota
	DATA  = iota
	EXTRA = iota
	BLOCK = iota
)
