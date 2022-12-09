package form3_client

import (
	"bytes"
	json "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Form3_API struct {
	Api_host_url     string `default:"http://localhost:8080"`
	Api_host_version string `default:"v1"`
}

func (client *Form3_API) CreateAccount(customerInfo AccountData) (response ResponseJSON) {
	var jsonData AccountJSON
	jsonData.Data = customerInfo
	inputstr, _ := json.Marshal(jsonData)
	log.Println("Creating a new Account")

	httpClient := http.Client{}
	log.Println(client.Api_host_url + "/" + client.Api_host_version + "/" + "organisation/accounts")
	log.Println(string(inputstr))
	resp, err := httpClient.Post(client.Api_host_url+"/"+client.Api_host_version+"/"+"organisation/accounts",
		"application/json", bytes.NewReader(inputstr))

	if err != nil {
		log.Println("Failed to create a new account" + err.Error())
		response.Status = "Failure"
		response.ErrorMessage = err.Error()
		return
	}

	if resp.StatusCode != 201 {
		response.Status = "Failure"
		response.ErrorCode = resp.StatusCode
		response.ErrorMessage = resp.Status
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body" + err.Error())
	}

	log.Println(string(body))

	var respJsonUnMarshalled AccountJSON
	errUnmarshall := json.Unmarshal(body, &respJsonUnMarshalled)
	if errUnmarshall != nil {
		log.Println("Error while unmarshalling response body" + errUnmarshall.Error())
		response.Status = "Failure"
		response.ErrorMessage = "Error while unmarshalling response body"
		return
	}

	response.Status = "Success"
	response.Data = respJsonUnMarshalled.Data

	return
}

func (client *Form3_API) FetchAccount(id string) (response ResponseJSON) {
	log.Println("Fetching Account details for Id:" + id)
	httpClient := &http.Client{}
	resp, err := httpClient.Get(client.Api_host_url + "/" + client.Api_host_version + "/" + "organisation/accounts/" + id)
	if err != nil {
		log.Println("Failed to fetch account info" + err.Error())
		response.Status = "Failure"
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

	log.Println(string(body))

	var respJsonUnMarshalled AccountJSON
	errUnmarshall := json.Unmarshal(body, &respJsonUnMarshalled)
	if errUnmarshall != nil {
		log.Println("Error while unmarshalling response body" + errUnmarshall.Error())
		response.Status = "Failure"
		response.ErrorMessage = "Error while unmarshalling response body"
		return
	}

	response.Status = "Success"
	response.Data = respJsonUnMarshalled.Data
	return
}

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
	response.Message = "Account was deleted successfully"

	return
}
