# LibKnot : Go interface for managing the Knot DNS daemon

LibKnot is a GoLang interface to manage the KnotDNS daemon using the C library ([libknot](https://github.com/CZ-NIC/knot/tree/master/tests/libknot)) based on the ([official Python interface](https://pypi.org/project/libknot/)].


Since the software loads the C library dynamically, you must install it on your development and production machine. 
When publishing compiled software, pay attention to the version of libknot used.


## Usage 

You can uses this library in two ways : 	
- Directly sending message using ``Send()``, one with the ``DATA`` unit type and one with a ``BLOCK`` unit type after and get the data back using ``Receive()`` and ``for`` loops.
- Using the wrapper functions : ``SendBlock()`` who will take care of the unit types and ``ReceiveBlock()`` which will output the received data as an array.


## Examples 

- Basic example using wrapper functions : 

```go
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
	_ = ctl.SendBlock(msg)
	/* End Send */

	/* Receive */
	data := ctl.ReceiveBlock()

	for _, s := range data {
		fmt.Println(s.Data) //Print received data
	}
	/* End Receive*/

	ctl.Terminate() //Clsoe Send(End) + Close() + Free()
}

```