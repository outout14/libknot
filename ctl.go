package golibknot

/*
 #cgo LDFLAGS: -lknot
 #cgo CFLAGS: -I /usr/include/libknot/
 #include <dlfcn.h>
 #include <libknot.h>
*/
import "C"

type KnotCtl struct {
	CtlAlloc *C.struct_knot_ctl
}

func (KnotCtl) StrErr(e int) string {
	return C.GoString(C.knot_strerror(C.int(e)))
}

func (ctl *KnotCtl) Alloc() {
	ctl.CtlAlloc = C.knot_ctl_alloc()
}

func (ctl *KnotCtl) Close() {
	C.knot_ctl_close(ctl.CtlAlloc)
}

func (ctl KnotCtl) Connect(path string) int {
	e := C.knot_ctl_connect(ctl.CtlAlloc, C.CString(path))
	return int(e)
}

func (ctl KnotCtl) Send(rType int, data *C.knot_ctl_data_t) int {
	e := C.knot_ctl_send(ctl.CtlAlloc, C.knot_ctl_type_t(rType), data)
	return int(e)
}

func (ctl KnotCtl) Receive(data *C.knot_ctl_data_t) (int, int) {
	var reqType C.knot_ctl_type_t
	e := C.knot_ctl_receive(ctl.CtlAlloc, &reqType, data)
	return int(e), int(reqType)
}
