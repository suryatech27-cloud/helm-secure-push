package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func handleapiRequest() string {

	//fmt.Println("Push Process Initiated .........................")

	//flag variable declaration
	var username string
	var password string
	var reponame string
	var tarfile string
	var provfile string

	//Variable Initialization for flag
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "Password")
	flag.StringVar(&reponame, "repo", "", "Specified your chart repository name")
	flag.StringVar(&tarfile, "chart", "foo.tgz", "Generated secured chart package (.tgz)")
	flag.StringVar(&provfile, "prov", "foo.tgz.prov", "Generated chart provenance file (.prov)")

	//Parsing all the flag value
	flag.Parse()

	//In order to push chart to repo , we need to execute bash script

	//we need to call script function execute the script

	output := executeShellScript(username, password, reponame, tarfile, provfile)

	return output

}

//function responsible executing the shell script
func executeShellScript(username string, password string, reponame string, tarfile string, provfile string) string {
	//fmt.Println("Connecting to the Harbor Api .......................")

	//executing bash script
	cmd, err := exec.Command("/bin/bash", "secure-push.sh", "--username", username, "--password", password, "--chart-tgz-file", tarfile, "--prov-file", provfile, "--repo-name", reponame).Output()

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error %s", err)
	}

	output := string(cmd)

	return output

}

func main() {
	response := handleapiRequest()
	fmt.Println(response)
}
