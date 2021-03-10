package main

import (
	"fmt"
	"strings"
	"os"
	"net"
	
	
)

func main() {
	fmt.Println("Environment Variables")
	fmt.Println(strings.Repeat("=", 21))
	// fmt.Println("PATH =", os.Getenv("PATH"))
	for _, lines := range os.Environ() {
		fmt.Println(lines)      	
    	}
	fmt.Println()
	
	fmt.Println("Network Interface Information")
	fmt.Println(strings.Repeat("=", 29))
	
	interfaces, _ := net.Interfaces()
   	for _, interf := range interfaces {
      		fmt.Print("Index:", interf.Index, " ")
      		fmt.Print("MTU:", interf.MTU, " ")
		fmt.Print("Name:", interf.Name, " ")
		fmt.Print("HardwareAddr:", interf.HardwareAddr, " ")
		fmt.Print("Flags:", interf.Flags, " ")
		addrs,_ := interf.Addrs()
		fmt.Println("IP Address:", addrs, " ")
	}
}

// End.
