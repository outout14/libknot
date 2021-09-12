# GoLibknot : Go interface for managing the Knot DNS daemon

Basic example : 

````go
package main

import "fmt"

func main() {
	ctl := KnotCtl{} //Define the connection

	err := ctl.Init("/var/run/knot/knot.sock") //Alloc + Connect to socket
	ctl.ErrCheck(err)

	/* Send */
	msg := KnotCtlData{
		Command: "conf-read",
		Section: "zone",
		Item:    "domain",
	}
	_ = ctl.SendBlock(&msg)
	/* End Send */ 
	
	/* Receive */
	data := ctl.ReceiveBlock()

	for _, s := range data {
		fmt.Println(s.Data) //Print received data 
	}
	/* End receive*/

	ctl.Send(END, nil) //Always terminate the connection
	ctl.Close()
}

````