package acl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	secret "github.com/blinchik/go-aws/lib/secrets"
)

type bootstrapACLResponse struct {
	ID          string `json:"ID"`
	AccessorID  string `json:"AccessorID"`
	SecretID    string `json:"SecretID"`
	Description string `json:"Description"`
	Policies    []struct {
		ID   string `json:"ID"`
		Name string `json:"Name"`
	} `json:"Policies"`
	Local       bool   `json:"Local"`
	CreateTime  string `json:"CreateTime"`
	Hash        string `json:"Hash"`
	CreateIndex int    `json:"CreateIndex"`
	ModifyIndex int    `json:"ModifyIndex"`
}

//BootstrapACL This endpoint does a special one-time bootstrap of the ACL system, making the first management token if the acl.tokens.master configuration entry is not
//specified in the Consul server configuration and if the cluster has not been bootstrapped previously.
func BootstrapACL(consulAddress, consulRootPath, consulPort string) {

	var output bootstrapACLResponse

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:%s/%s/acl/bootstrap", consulAddress, consulPort, consulRootPath), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bodyBytes, &output)
	if err != nil {
		log.Println(string(bodyBytes))

		if strings.Contains(string(bodyBytes), "ACL bootstrap no longer allowed") {

			return
		} else {

			log.Fatal(err)

		}
	}

	secret.CreateSecret(output.Description, output.Policies[0].Name, output.SecretID)

}

func UpdateACLToken(consulAddress, consulRootPath, consulPort, token, consulToken string) {

	payload := fmt.Sprintf(` {"token": "%s" }`, token)

	body := strings.NewReader(payload)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:%s/%s/agent/token/agent", consulAddress, consulPort, consulRootPath), body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", consulToken))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bodyBytes))

}
