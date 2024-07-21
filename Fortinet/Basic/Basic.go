

package main

import (
    "bufio" // Package for reading input
    "fmt"   // Package for formatting and printing
    "os"    // Package for operating system functions
    "github.com/dongri/ssh" // Package for SSH client
)

func main() {
    // Prompt user to enter the username for SSH connection
    fmt.Print("Enter the username for SSH connection: ")
    scanner := bufio.NewScanner(os.Stdin) // Create a new scanner to read user input
    scanner.Scan() // Read the next token from the input
    username := scanner.Text() // Store the entered username

    // Prompt user to enter the password for SSH connection
    fmt.Print("Enter the password for SSH connection: ")
    scanner.Scan() // Read the next token from the input
    password := scanner.Text() // Store the entered password

    // Connect to Fortinet 2000e firewall via SSH using provided username and password
    client, err := ssh.Dial("fortinet2000e-firewall-<Ip address>.com:22", username, password)
    if err != nil {
        fmt.Println("Error connecting to firewall:", err) // Print error message if connection fails
        return // Exit the program
    }

    // Configure 3 GE ports physical interface
    commands := []string{
        "set interface ge0/0/1",
        "set interface ge0/0/2",
        "set interface ge0/0/3",
    }

    // Execute commands to configure GE ports
    for _, cmd := range commands {
        if err := client.RunCommand(cmd); err != nil {
            fmt.Println("Error configuring interface:", err) // Print error message if configuration fails
        }
    }

    // Create a bundle Etherchannel with those interfaces
    client.RunCommand("set interface ethernet-channel")

    // Associate GE interfaces with the Ethernet Channel
    for i := 1; i <= 3; i++ {
        client.RunCommand(fmt.Sprintf("set interface ge0/0/%d ethernet-channel ge-chan", i))
    }

    // Configure hostname as HQ-MX
    client.RunCommand("set hostname HQ-MX")

    // Configure a loopback interface with address 192.168.20.1/27
    client.RunCommand("set interface loopback address 192.168.20.1/27")

    // Close SSH connection
    client.Close()
}
