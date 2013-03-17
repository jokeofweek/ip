# Go IP API
This simple library allows you to get IP address location information based on the [hostip.info](http://www.hostip.info/) API.
I was just using this as an excuse to learn Go!

## Example
```Go
package main

import (
	"github.com/jokeofweek/ip"
	"fmt"
)

func main() {
	record, err := ip.GetOwnIpRecord()
	if err != nil {
		fmt.Printf("Error getting own IP record: %v\n", err)
	} else {
		fmt.Printf("Own IP Data:\n %v\n", record)
	}

	record, err = ip.GetIpRecord("74.125.224.72")
	if err != nil {
		fmt.Printf("Error getting 74.125.224.72's record: %v\n", err)
	} else {
		fmt.Printf("74.125.224.72 IP Data:\n %v\n", record)
	}

	record, err = ip.GetIpRecord("184.72.186.1")
	if err != nil {
		fmt.Printf("Error getting 184.72.186.1's record: %v\n", err)
	} else {
		fmt.Printf("184.72.186.1 IP Data:\n %v\n", record)
	}
}
```