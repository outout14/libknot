# LibKnot : Go interface for managing the Knot DNS daemon

LibKnot is a GoLang interface to manage the KnotDNS daemon using the C library ([libknot](https://github.com/CZ-NIC/knot/tree/master/tests/libknot)) based on the ([official Python interface](https://pypi.org/project/libknot/)].


Since the software loads the C library dynamically, you must install it on your development and production machine. 
When publishing compiled software, pay attention to the version of libknot used.


## Usage 

You can uses this library in two ways : 
- Use wrapper functions, easier to use. 	
- Directly control the server with the functions ``Alloc()``, ``Connect()``, ``Close()``, ``Free()``, ``Send()``, ``Receive()`` and manage data types. 

## Examples 

- Basic example using wrapper functions : 

```go
package main

import (
	"fmt"
	
	"github.com/outout14/libknot"
)

func main() {
	ctl := libknot.KnotCtl{} //Define the connection

	err := ctl.Init("/var/run/knot/knot.sock") // Allocate de control context + Connect(socket)
	ctl.ErrCheck(err)                          //Panic if err != 0 and print the error as a human readble string

	/* Send */
	msg := libknot.KnotCtlData{
		Command: "conf-read",
		Section: "zone",
		Item:    "domain",
	}
	_ = ctl.SendBlock(msg) //Send DATA msg + BLOCK (BLOCK = End of message)
	/* End Send */

	/* Receive */
	data := ctl.ReceiveBlock() //Retrieves all DATA and EXTRA type messages

	for _, s := range data {
		fmt.Println(s.Data) //Print received data
	}
	/* End Receive*/

	ctl.Terminate() //Send(END) + Close() + Free() control ctx
}
```
