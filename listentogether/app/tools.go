package app

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

func ipFrom(ipv4 string) net.IP {
	ipl := strings.Split(ipv4, ".")
	if len(ipl) != 4 {
		return nil
	}
	a, err := strconv.Atoi(ipl[0])
	if err != nil {
		return nil
	}
	b, _ := strconv.Atoi(ipl[1])
	if err != nil {
		return nil
	}
	c, _ := strconv.Atoi(ipl[2])
	if err != nil {
		return nil
	}
	d, _ := strconv.Atoi(ipl[3])
	if err != nil {
		return nil
	}
	return net.IPv4(byte(a), byte(b), byte(c), byte(d))
}

var cacheFile string = "._listen_together_cached_ip.txt"

func cacheIP(ip string) {
	ioutil.WriteFile(cacheFile, []byte(ip), 0644)
}

func getCachedIP() string {
	ip, err := ioutil.ReadFile(cacheFile)
	if err == nil {
		return strings.TrimSpace(string(ip))
	}
	return ""
}

func Say520() {
	fmt.Println("❤\tBo love Fish, best.\t❤")
}
