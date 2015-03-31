package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	var scheme string
	hostname, _ := os.Hostname()

	pflag.StringVarP(&scheme, "scheme", "s", hostname, "SCHEME")
	pflag.Parse()

	fmt.Printf("%s.disk.usage %f %d\n", scheme, diskUsage(), time.Now().Unix())
}

func diskUsage() float64 {
	out, _ := exec.Command("df", "-lP").Output()
	lines := strings.Split(strings.TrimRight(string(out), "\n"), "\n")[1:]

	var used float64 = 0.0
	var available float64 = 0.0

	for _, line := range lines {
		stats := strings.Fields(line)
		used += ParseFloat(stats[2])
		available += ParseFloat(stats[3])
	}

	return 100.0 * used / (used + available)
}

func ParseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
