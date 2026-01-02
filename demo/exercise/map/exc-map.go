//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
func printStatus(servers map[string]int) {
	totalServers := len(servers)
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println("There are",totalServers," of server")
	fmt.Println("Total Online", stats[Online])
	fmt.Println("Total Offline", stats[Offline])
	fmt.Println("Total Maintenance", stats[Maintenance])
	fmt.Println("Total Retired", stats[Retired], "\n")
}


func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
//  as the value
	serversStatus := make(map[string]int)

	for _, server := range servers {
		serversStatus[server] = Online
	}

	//  - call display server info function
	printStatus(serversStatus)
	//  - change server status of `darkstar` to `Retired`
	serversStatus["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serversStatus["aiur"] = Offline
	//  - call display server info function
	printStatus(serversStatus)
	//  - change server status of all servers to `Maintenance`
	for server, _ := range serversStatus {
		serversStatus[server] = Maintenance
	}
	//  - call display server info function
	printStatus(serversStatus)
}

