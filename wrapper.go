package golibknot

func (ctl KnotCtl) ErrCheck(e int) {
	if e != 0 {
		panic(ctl.StrErr(e))
	}
}

func (ctl *KnotCtl) Init(path string) int {
	ctl.Alloc()
	err := ctl.Connect(path)
	return err
}

func (ctl KnotCtl) SendBlock(data KnotCtlData) int {
	x := data.ToCtl()
	err := ctl.Send(DATA, &x)
	if err != 0 {
		return err
	}
	ctl.Send(BLOCK, nil)
	return 0
}

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
