package libknot

// ErrCheck panics if there is an error
func (ctl KnotCtl) ErrCheck(e int) {
	if e != 0 {
		panic(ctl.StrErr(e))
	}
}

// Init initializes the connection to the KnotDNS UNIX sockets. It allocates the control context and connects to the socket.
func (ctl *KnotCtl) Init(path string) int {
	ctl.Alloc()
	err := ctl.Connect(path)
	return err
}

// Terminate Sends END, Closes the UNIX socket, deallocates (Free) the control socket.
func (ctl *KnotCtl) Terminate() {
	ctl.Send(END, nil)
	ctl.Close()
	ctl.Free()
}

// SendBlock sends one control unit with the data and an other one with the BLOCK unit type telling the KnotDNS server that the message is finished and that it can respond to us.
func (ctl KnotCtl) SendBlock(data KnotCtlData) int {
	x := data.ToCtl()
	err := ctl.Send(DATA, &x)
	if err != 0 {
		return err
	}
	ctl.Send(BLOCK, nil)
	return 0
}

// ReceiveBlock retrieves all data (DATA and EXTRA unit type) received by the server until it tells us (END or BLOCK unit type) that it is finished.
func (ctl KnotCtl) ReceiveBlock() []KnotCtlData {
	var answers []KnotCtlData
	for {
		var receivedData KnotCtlData
		msg := KnotCtlData{}.ToCtl()

		err, reqType := ctl.Receive(&msg)
		ctl.ErrCheck(err)

		if reqType != DATA && reqType != EXTRA {
			break
		}

		receivedData.FromCtl(msg)

		if receivedData.Error != "" {
			panic(receivedData.Error)
		}
		answers = append(answers, receivedData)
	}
	return answers
}
