package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func main() {
	log.Println("starting")
	apikey, err := ioutil.ReadFile(os.Getenv("API_KEY_FILE"))
	if err != nil {
		log.Fatalf("could not read API_KEY_FILE ('%s'): %s", os.Getenv("API_KEY_FILE"), err)
	}
	client := hcloud.NewClient(hcloud.WithToken(strings.TrimSpace(string(apikey))))
	targetHost := os.Getenv("TARGET_HOST")

	targetServer, _, err := client.Server.GetByName(context.Background(), targetHost)
	if err != nil {
		log.Fatalf("could not find target server '%s': %s", targetHost, err)
	}

	floatingIPID, err := strconv.Atoi(os.Getenv("FLOATING_IP_ID"))
	if err != nil {
		log.Fatalf("FLOATING_IP_ID not an int: %s", err)
	}

	// make first tick immediate
	ticker := time.Tick(30 * time.Second)
	for ; true; <-ticker {
		fip, _, err := client.FloatingIP.GetByID(context.Background(), floatingIPID)
		if err != nil {
			log.Fatalf("could not get floating IP %d: %s", floatingIPID, err)
		}
		if fip.Server.ID == targetServer.ID {
			log.Printf("already set to %s", targetServer.Name)
		} else {
			log.Printf("updating to %s", targetServer.Name)
			_, _, err = client.FloatingIP.Assign(context.Background(), fip, targetServer)
			if err != nil {
				log.Printf("ERROR assigning floating IP: %s", err)
			}
		}
	}
}
