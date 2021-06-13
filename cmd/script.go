package script

import (
	"fmt"
	"log"
	"os/exec"
)

func executeShellScript(username string, password string, reponame string, tarfile string, provfile string) string {
	fmt.Println("Connecting to the Harbor Api .........")

	//executing shell script using os/exec

	cmd, err := exec.Command("/bin/bash", "secure-push.sh --username "+username+" --password "+password+" --chart-tgz-file "+tarfile+" --prov-file "+provfile+" --repo-name "+reponame).Output()

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error %s", err)
	}

	output := string(cmd)

	return output

}
