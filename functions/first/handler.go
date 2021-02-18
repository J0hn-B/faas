package function

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {
	GetTime()
	GetOs()
	GetID()
	CurrentUser()
	//GetEnvVars()
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

// GetID returns the numeric effective group and user id of the caller
func GetID() {
	GroupID := os.Getegid()
	UserID := os.Geteuid()
	fmt.Println("Group ID is:", GroupID, "User ID is:", UserID)
}

// CurrentUser gets current user info
func CurrentUser() {
	currentUser, err := user.Current()
	if err == nil {
		fmt.Println("Username:", currentUser.Username)
		fmt.Println("HomeDir:", currentUser.HomeDir)
	} else {
		panic(err)
	}

}

// // GetEnvVars returns the environment variables
// func GetEnvVars() {
// 	// get all env vars
// 	var env []string = os.Environ()
// 	fmt.Println("Environment Variables:")
// 	//iterate over each variable
// 	for index, variable := range env {
// 		nameVal := strings.Split(variable, "=")

// 		fmt.Printf("[%d] %v => %v\n", index, nameVal[0], nameVal[1])
// 	}
// }
