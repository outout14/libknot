package libknot

/*
 #cgo LDFLAGS: -lknot
 #cgo CFLAGS: -I /usr/include/libknot/
 #include <dlfcn.h>
 #include <libknot.h>
*/
import "C"

// KnotCtl keeps control ctx and is used to perform functions related to the KnotDNS server.
type KnotCtl struct {
	CtlAlloc *C.struct_knot_ctl
}

// StrErr converts a Knot Error (int) to a usefull human readable error (string).
func (KnotCtl) StrErr(e int) string {
	return C.GoString(C.knot_strerror(C.int(e)))
}

// Alloc creates a control context.
func (ctl *KnotCtl) Alloc() {
	ctl.CtlAlloc = C.knot_ctl_alloc()
}

// Free deallocates a control socket.
func (ctl *KnotCtl) Free() {
	C.knot_ctl_free(ctl.CtlAlloc)
}

// Close closes connections to the server.
func (ctl *KnotCtl) Close() {
	C.knot_ctl_close(ctl.CtlAlloc)
}

// Connect connects to the KnotDNS UNIX socket.
func (ctl KnotCtl) Connect(path string) int {
	e := C.knot_ctl_connect(ctl.CtlAlloc, C.CString(path))
	return int(e)
}

// Send sends one control unit to the socket. The unit type to send should be precised. Data can be nil for non-data type (eg: BLOCK).
func (ctl KnotCtl) Send(unitType int, data *C.knot_ctl_data_t) int {
	e := C.knot_ctl_send(ctl.CtlAlloc, C.knot_ctl_type_t(unitType), data)
	return int(e)
}

// Receive receives one control unit from the socket.
func (ctl KnotCtl) Receive(data *C.knot_ctl_data_t) (int, int) {
	var reqType C.knot_ctl_type_t
	e := C.knot_ctl_receive(ctl.CtlAlloc, &reqType, data)
	return int(e), int(reqType)
}
