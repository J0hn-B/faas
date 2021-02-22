package function

import (
	"fmt"
	"os"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {
	GetTime()
	GetOs()
	return fmt.Sprintf("This function returns container information. %s", string(req))
}

// GetTime exports time
func GetTime() {
	t := time.Now()
	fmt.Println("Time:", t.Format("15:04:05"))
	fmt.Println("Date:", t.Format("02-01-2006"))
}

// GetOs returns the OS hostname
func GetOs() {
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Container Hostname: ", hostname)
	} else {
		panic(err)
	}
}
