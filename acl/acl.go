package acl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
func BootstrapACL(consulAddress, consulPort string) {

	var output bootstrapACLResponse

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:%s/v1/acl/bootstrap", consulAddress, consulPort), nil)
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
		log.Fatal(err)
	}

	secret.CreateSecret(output.Description, output.Policies[0].Name, output.SecretID)

}
