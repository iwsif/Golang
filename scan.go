package main

import "C"
import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
    "bufio"
    
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Scanner struct {
	Port int
	Is_open chan string
}

func options() {
	fmt.Println()
	fmt.Println("\t\t\tScanner written in go(VERSION 1.0)")
	fmt.Println("\tOPTIONS:")
	fmt.Println("\t\t--help:Show help menu!!")
	fmt.Println("\t\t--ping:Ping a domain_name or ip to see if host is up!!")
	fmt.Println("\t\t--get_ip:Get ip of domain name")
	fmt.Println("\t\t--host:Type --host and ip of your target(You can type also your hostname if you want to scan your system)")
	fmt.Println("\t\t--port_list:Type --port_list to generate list of ports for known services")
	fmt.Println("\t\t--get_port:Type --get_port + service name to get port number for a specific service")
	fmt.Println("\t\t--port:Scan for specific tcp port(Type --port + port number)")
	fmt.Println("\t\t--tcp_scan:Scan all 65535 tcp_ports(Type --tcp_scan + 'all')")
	fmt.Println("\t\t--udp_scan:Scan all 65535 udp_ports(Type --udp_scan + 'all'")
	fmt.Println()

}

func scan_port(port int,hostname string) {
   if port > 0 && port < 65535 {
      fmt.Println("Starting...")
      fmt.Println("Please wait...")
      time.Sleep(time.Duration(time.Millisecond * 1000))
      channel := scan_function("tcp",hostname,port)
      if fmt.Sprintln(<-channel) == "True" {
        fmt.Println("Port is open!!")
      }else {
        fmt.Println("Port is closed")
      }
    }
  
}


func if_is_up( ip string) string {
	value1 := "Error running command..."
	value2 := "TRUE"
	value3 := "FALSE"
	var values []string
	var result string

	values = append(values,value1)
	values = append(values,value2)
	values = append(values,value3)


	out,err := exec.Command("ping " + ip).Output()
	if err != nil {
		result = values[0]
	}
	string_output := string(out)
	substring := "Destination Host Unreachable"
	true_or_false := strings.Contains(string_output,substring)
	if true_or_false == true {
		result = values[2]
	}else {
		result = values[1]
	}
	return result
}

func ping(domain_name string) ([]string,error) {
	var myerror error
	var domain_ip []string
	result,err  := net.LookupIP(domain_name)
	if err != nil {
		fmt.Println(err)
		myerror = err
	}
  	myerror = nil
	for _,ip := range result {
		domain_ip = append(domain_ip,ip.String())
	}
	
	return domain_ip,myerror
}


func enumerate(slice []string) int {
	var counter int
	for _,elements := range slice {
		fmt.Sprint(elements)
		counter = counter + 1
	}
	return counter
}

func get_opt() []string {
	var get_options []string
	get_options= append(get_options,"--help")
	get_options = append(get_options,"--ping")
	get_options = append(get_options,"--get_ip")
	get_options = append(get_options,"--host")
	get_options = append(get_options,"--port")
	get_options = append(get_options,"--tcp_scan")
	get_options = append(get_options,"--udp_scan")
	get_options = append(get_options,"--get_port")
	return get_options
}

func iter(word string) string {
	var result string
	values := get_opt()
	for _,value :=range values {
		if value == word {
			result = "TRUE"
		}else {
			result = "FALSE"
		}
	}
	return result
}

func port_list() []string {

	var list []string
	out,err := exec.Command("cat","/etc/services").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Exiting...")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		os.Exit(0)
	}else {
		fmt.Println("Generating list with ports...")
		fmt.Println("Please wait...")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		list = append(list, string(out))
	}
	return list
}


func scan_function(protocol string,hostname string,port int) chan string{

	result := make(chan string)
	address := hostname + ":" + strconv.Itoa(port)
	var myerror error
	go func() {
		connection, err := net.DialTimeout(protocol, address, time.Duration(time.Millisecond*4000))
		myerror = err
		if err != nil {
			result <- "false"
			//log.Fatal(err)
		}
		result <- "true"
		wg.Done()
		connection.Close()
	}()
	return result
}

func scanner(hostname string,scan_type string) {
	var tcp_scanner []Scanner
	var udp_scanner []Scanner
	if scan_type == "tcp" {
		for i := 0; i < 65536; i++ {
			tcp_scanner = append(tcp_scanner, Scanner{Port: i, Is_open: scan_function("tcp", hostname, i)})
		}
	}else if scan_type == "udp " {
		for i:=0; i < 65536; i++ {
			udp_scanner = append(udp_scanner,Scanner{Port: i,Is_open: scan_function("udp",hostname,i)})
		}
	}
	if scan_type == "tcp" {
		for _,values := range tcp_scanner {
			fmt.Println("Port:" + strconv.Itoa(values.Port) + "/" + <-values.Is_open)
			if values.Port == 65535 {
				fmt.Println("Scanned 65535-tcp ports...")
				fmt.Println("Exiting...")
				time.Sleep(time.Duration(time.Millisecond * 2000))
				os.Exit(0)
			}
		}
	}else if scan_type == "udp" {
		for _,values := range udp_scanner {
			fmt.Println("Port: " + strconv.Itoa(values.Port) + "/" + <-values.Is_open)
			if values.Port == 65535 {
				fmt.Println("Scanned 65535-udp ports...")
				fmt.Println("Exiting...")
				time.Sleep(time.Duration(time.Millisecond * 2000))
				os.Exit(0)
			}
		}
	}
}

func main() {
	err := exec.Command("uname").Run()
	if err != nil {
		fmt.Println("No supported operating system detected...")
		fmt.Println("Exiting....")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		os.Exit(0)
	}

	var length int
	length = 0
	length = enumerate(os.Args)
	if length == 1 {
		fmt.Println("No arguments supplied...")
		fmt.Println("Starting help_menu...")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		options()
	}

	if length == 2 && os.Args[1] == "--help" {
		fmt.Println("Starting help menu....")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		options()
	}
	if length == 2 && os.Args[1] == "--ping" {
		fmt.Println("You have to type an ip or a domain name in order to see if the remote server is up!!")
		fmt.Println("See --help for more info...")
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	if length == 2 && os.Args[1] == "get_ip" {
		fmt.Println("You have to type the domain_name in order to get the ip!!")
		fmt.Println("Type --help for more info")
		fmt.Println("Exiting...")
		os.Exit(0)
	}

	if length == 2 && os.Args[1] == "--host" {
		fmt.Println("You have to type also the ip of your target or hostname")
		fmt.Println("Type --help for more info")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		os.Exit(0)
	}
	if length == 2 && os.Args[1] == "--port" {
		fmt.Println("You have to provide also the number of port you want to scan...")
		list := port_list()
		for _, output := range list {
			fmt.Println(output)
		}
		time.Sleep(time.Duration(time.Millisecond * 2000))
	}

	if length == 2 && os.Args[1] == "--tcp_scan" {
		fmt.Println("You have to provide also 'all' keyword")
		fmt.Println("Type --help for more info ")
		fmt.Println("Exiting...")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		os.Exit(0)
	}

	if length == 2 && os.Args[1] == "--udp_scan" {
		fmt.Println("You have to provide also 'all' keyword")
		fmt.Println("Type --help for more info ")
		fmt.Println("Exiting...")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		os.Exit(0)
	}

	if length == 2 && os.Args[1] == "--get_port" {
		fmt.Println("You have to provide also the service you want to search for..")
		fmt.Println("Type --help for more info...")
		fmt.Println("Exiting...")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		os.Exit(0)
	}

	if length == 3 && os.Args[1] == "--get_ip" && os.Args[2] != "" {
		fmt.Println("Getting the ip of the target...")
		fmt.Println("Please wait..")
		time.Sleep(time.Duration(time.Millisecond * 2000))
		result, myerror := ping(os.Args[2])
		if myerror != nil {
			fmt.Println("Domain name doesnt exist...")
			fmt.Println("Exiting...")
			time.Sleep(time.Duration(time.Millisecond * 2000))
		} else {
			slice_length := enumerate(result)
			for i := 0; i < slice_length; i++ {
				fmt.Println(result[i])
				time.Sleep(time.Duration(time.Millisecond * 1000))
			}
		}
	}
	if length == 3 && os.Args[1] == "--ping" && os.Args[2] != "" {
		result := if_is_up(os.Args[2])
		if result == "TRUE" {
			fmt.Println("Host is up!!")
			fmt.Println()
			time.Sleep(time.Duration(time.Millisecond * 2000))
		} else {
			fmt.Println("Host is down")
		}
	}
	if length == 3 && os.Args[1] == "--get_port" && os.Args[2] != "" {

	  pattern, err := regexp.Compile("(?:" + os.Args[2] + ")" + "(.*)")
	  if err != nil {
	    fmt.Println(err)
	  }
	  string_pattern := pattern.String()
	  ports := port_list()
	  var result interface{}
	  //var Port string
      for _, port := range ports {
		result, err = regexp.MatchString(string_pattern, port)
		if err != nil {
		  fmt.Println(err)
		}
      }
      if result != true {
	    fmt.Println("Unknown service..")
	    fmt.Println("Exiting..")
	    time.Sleep(time.Duration(time.Millisecond * 2000))
	    os.Exit(0)
      }
      first_command := exec.Command("cat", "/etc/services")
      second_command := exec.Command("grep",os.Args[2])
      second_command.Stdin,err = first_command.StdoutPipe()
      first_command.Start()
      second_command.Wait()
      if err != nil {
          fmt.Println(err)
        }
      grep,err := second_command.Output()
      if err != nil {
        fmt.Println(err)
      }
      fmt.Println(string(grep))
    }
    if length == 3 && os.Args[1] == "--port" && os.Args[2] != "" {
      var name string
      fmt.Println("Type the domain name you want to scan for specific port")
      scanner := bufio.NewScanner(os.Stdin)
      scanner.Scan()
      name = scanner.Text()
      fmt.Println("Checking if the domain_name is valid...")
      time.Sleep(time.Duration(time.Millisecond*2000))
      result,err := ping(name)
      fmt.Sprintln(result)
      if err != nil {
        fmt.Println()
        fmt.Println("Domain name doesnt exist!!")
        time.Sleep(time.Duration(time.Millisecond*2000))
        os.Exit(0)
      }
      port,err := strconv.Atoi(os.Args[2])
      scan_port(port,name)
    }
	if length == 5 {
		wg.Add(65535)
		fmt.Println("Welcome this is scanner written in golang(Version 1.0.0)")
		fmt.Println("Starting...")
		time.Sleep(time.Duration(time.Millisecond * 3000))
		uname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error exiting...")
			os.Exit(0)
		}
		pattern1, err := regexp.Compile(uname)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error exiting...")
			os.Exit(0)
		}
		pattern2, err := regexp.Compile("[0-9]{1,3}'.'[0-9]{1,3}'.'[0-9]{1,3}'.'[0-9]{1,3}")
		if err != nil {
			fmt.Println(err)
			fmt.Println("Exiting...")
			  os.Exit(0)
			}
		string_pattern1 := pattern1.String()
	  	string_pattern2 := pattern2.String()
		if os.Args[1] == "--tcp_scan" && os.Args[2] == "all" && os.Args[3] == "--host" && os.Args[4] != "" || os.Args[1] == "--host" && os.Args[2] != "" && os.Args[3] == "--tcp_scan" && os.Args[4] == "all" {
			var hostname string
			if os.Args[1] == "--host" {
				if iter(os.Args[2]) == "TRUE" {
					fmt.Println("Error...wrong value")
					fmt.Println("Exiting...")
					time.Sleep(time.Duration(time.Millisecond * 2000))
					os.Exit(0)
				} else {
					result, err := regexp.MatchString(string_pattern1, os.Args[2])
					if err != nil {
						fmt.Println(err)
					}
					result2, err := regexp.MatchString(string_pattern2, os.Args[2])
					if err != nil {
						fmt.Println(err)
					}
					if result == true || result2 == true {
						hostname = os.Args[2]
						scanner(hostname, "tcp")
						wg.Wait()
					}
				}
			} else if os.Args[3] == "--host" {
				if iter(os.Args[4]) == "TRUE" {
					fmt.Println("Wrong value...")
					fmt.Println("Exiting...")
					time.Sleep(time.Duration(time.Millisecond * 2000))
					os.Exit(0)
				} else {
					result1, err := regexp.MatchString(string_pattern1, os.Args[4])
					if err != nil {
						fmt.Println(err)
					}
					result2, err := regexp.MatchString(string_pattern2, os.Args[4])
					if err != nil {
						fmt.Println(err)
					}
					if result1 == true || result2 == true {
						hostname = os.Args[4]
						scanner(hostname, "tcp")
						wg.Wait()
					} else {
						fmt.Println("Wrong value...")
						fmt.Println("Type --help for more info...")
						fmt.Println("Exiting...")
						time.Sleep(time.Duration(time.Millisecond * 2000))
						os.Exit(0)
					}
				}
			}
		} else if os.Args[1] == "--udp_scan" && os.Args[2] == "all" && os.Args[3] == "--host" && os.Args[4] != "" || os.Args[1] == "--host" && os.Args[2] != "" && os.Args[3] == "--udp_scan" && os.Args[4] == "all" {
				if os.Args[3] == "--host" {
					result1, err := regexp.MatchString(string_pattern1, os.Args[4])
					if err != nil {
						fmt.Println(err)
					}
					result2, err := regexp.MatchString(string_pattern2, os.Args[4])
					if err != nil {
						fmt.Println(err)
					}
					if result1 == true || result2 == true {
						hostname := os.Args[4]
						scanner(hostname, "udp")
					} else {
						fmt.Println("Wrong value..")
						fmt.Println("Type --help for more info")
						fmt.Println("Exiting...")
						time.Sleep(time.Duration(time.Millisecond * 2000))
						os.Exit(0)
					}
				}
				if os.Args[1] == "--host" {
					result1, err := regexp.MatchString(string_pattern1, os.Args[2])
					if err != nil {
						fmt.Println(err)
					}
					result2, err := regexp.MatchString(string_pattern2, os.Args[2])
					if err != nil {
						fmt.Println(err)
					}
					if result1 == true || result2 == true {
						hostname := os.Args[2]
						scanner(hostname, "udp")
					} else {
						fmt.Println("Error incorrect value....")
						fmt.Println("Type --help for more info...")
						fmt.Println("Exiting...")
						time.Sleep(time.Duration(time.Millisecond * 2000))
						os.Exit(0)
					}
				}
            }
      }
}
















































