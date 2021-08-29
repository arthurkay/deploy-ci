package main

import (
	"deploy-ci/ansible"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Please specify the port, ansible playbook and user to use in ansible execution")
		fmt.Println("e.g: deploy-ci 3000 test.yml arthur")
		os.Exit(1)
	}
	port := os.Args[1]
	route := http.NewServeMux()
	route.HandleFunc("/", RequestHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), route)
	if err != nil {
		log.Fatal(err)
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var playbook string = os.Args[2]
	var user string = os.Args[3]

	if playbook == "" && user == "" {
		err := ansible.Run("test.yml", "arthur")
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
	}

	err := ansible.Run(playbook, user)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}
