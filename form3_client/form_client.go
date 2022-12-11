//go:build darwin || linux || windows
// +build darwin linux windows

package form3_client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Form3_API struct {
	Api_host_url     string `default:"http://localhost:8080"`
	Api_host_version string `default:"v1"`
}

/*
API request for fetching account info, expects account JSON as input.
Returns the success or failure message along with Account data form3 api.
*/

func (client *Form3_API) CreateAccount(customerInfo AccountData) (response ResponseJSON) {
	var jsonData AccountJSON
	jsonData.Data = customerInfo
	inputstr, _ := json.Marshal(jsonData)
	log.Println("Creating a new Account")

	httpClient := http.Client{}
	resp, err := httpClient.Post(client.Api_host_url+"/"+client.Api_host_version+"/"+"organisation/accounts",
		"application/json", bytes.NewReader(inputstr))

	if err != nil {
		log.Println("Failed to create a new account" + err.Error())
		response.Status = "Failure"
		response.ErrorMessage = err.Error()
		return
	}

	if resp.StatusCode != 201 {
		log.Println("failure" + resp.Status)
		response.Status = "Failure"
		response.ErrorCode = resp.StatusCode
		response.ErrorMessage = resp.Status
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body" + err.Error())
	}

	var respJsonUnMarshalled = AccountJSON{}

	errUnmarshall := json.Unmarshal(body, &respJsonUnMarshalled)

	if errUnmarshall != nil {
		log.Println("Error while unmarshalling response body" + errUnmarshall.Error())
		response.Status = "Failure"
		response.ErrorMessage = "Error while unmarshalling response body"
		return
	}

	response.Status = "Success"
	response.Data = append(response.Data, respJsonUnMarshalled.Data)

	return
}

/*
API request for fetching account info, expects account id as input.
Returns the success or failure along with Account data returned by the request to form3 api.
*/
func (client *Form3_API) FetchAccount(id string) (response ResponseJSON) {
	log.Println("Fetching Account details for Id:" + id)
	httpClient := &http.Client{}
	resp, err := httpClient.Get(client.Api_host_url + "/" + client.Api_host_version + "/" + "organisation/accounts/" + id)
	if err != nil {
		log.Println("Failed to fetch account info" + err.Error())
		response.Status = "Failure"
		if resp != nil {
			response.ErrorCode = resp.StatusCode
		}
		response.ErrorMessage = err.Error()
		return
	}
	if resp.StatusCode != 200 {
		response.Status = "Failure"
		response.ErrorCode = resp.StatusCode
		response.ErrorMessage = resp.Status
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Status = "Failure"
		response.ErrorMessage = resp.Status
		log.Println("Error while reading response body" + err.Error())
		return
	}

	var respJsonUnMarshalled = AccountJSON{}
	errUnmarshall := json.Unmarshal(body, &respJsonUnMarshalled)

	if errUnmarshall != nil {
		log.Println("Error while unmarshalling response body" + errUnmarshall.Error())
		response.Status = "Failure"
		response.ErrorMessage = "Error while unmarshalling response body"
		return
	}

	response.Status = "Success"
	response.Data = append(response.Data, respJsonUnMarshalled.Data)
	return
}

/*
API request for fetching account info, expects account id as input.
Returns the success or failure along with Account data returned by the request to form3 api.
*/
func (client *Form3_API) FetchAllAccounts() (response ResponseJSON) {
	log.Println("Fetching Account details for Id:")
	httpClient := &http.Client{}
	resp, err := httpClient.Get(client.Api_host_url + "/" + client.Api_host_version + "/" + "organisation/accounts/")
	if err != nil {
		log.Println("Failed to fetch account info" + err.Error())
		response.Status = "Failure"
		if resp != nil {
			response.ErrorCode = resp.StatusCode
		}
		response.ErrorMessage = err.Error()
		return
	}
	if resp.StatusCode != 200 {
		response.Status = "Failure"
		response.ErrorCode = resp.StatusCode
		response.ErrorMessage = resp.Status
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Status = "Failure"
		response.ErrorMessage = resp.Status
		log.Println("Error while reading response body" + err.Error())
		return
	}
	errUnmarshall := json.Unmarshal(body, &response)

	if errUnmarshall != nil {
		log.Println("Error while unmarshalling response body" + errUnmarshall.Error())
		response.Status = "Failure"
		response.ErrorMessage = "Error while unmarshalling response body"
		return
	}

	response.Status = "Success"
	response.StatusCode = resp.StatusCode
	return
}

/*
API request for delete account info expects account id as input.
Returns the success or failure code based on response from form3 api.
*/

func (client *Form3_API) DeleteAccount(id string) (response ResponseJSON) {
	log.Println("Deleting account ID:", id)
	httpClient := &http.Client{}
	req, err := http.NewRequest("DELETE", client.Api_host_url+"/"+client.Api_host_version+"/"+"organisation/accounts/"+id+"?version=0", bytes.NewReader([]byte{}))
	if err != nil {
		log.Println("Failed to create delete request" + err.Error())
		response.Status = "Failure"
		response.ErrorMessage = err.Error()
		return
	}

	resp, errResp := httpClient.Do(req)

	if errResp != nil {
		log.Println("Failed to delete account" + errResp.Error())
		response.Status = "Failure"
		response.ErrorMessage = errResp.Error()
		return
	}

	if resp.StatusCode != 204 {
		response.Status = "Failure"
		response.ErrorCode = resp.StatusCode
		response.ErrorMessage = resp.Status
		return
	}

	response.Status = "Success"

	return
}
