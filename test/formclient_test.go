//go:build darwin || linux || windows
// +build darwin linux windows

package sampleclient

import (
	"github.com/google/uuid"
	"github.com/kayestee/f3_client/form3_client"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var generatedId string
var form3cli = form3_client.Form3_API{
	Api_host_url:     os.Getenv("API_HOST"),
	Api_host_version: os.Getenv("API_VERSION"),
}

func TestCreateAccount(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var inputAccount form3_client.AccountData
	inputAccount.ID = uuid.New().String()
	inputAccount.Type = "accounts"
	inputAccount.OrganisationID = uuid.New().String()

	var name = []string{"First", "Name", "Last", "Name"}

	var accoutAttrs form3_client.AccountAttributes
	accoutAttrs.Country = "GB"
	accoutAttrs.BankID = "400302"
	accoutAttrs.BankIDCode = "GBDSC"
	accoutAttrs.AccountNumber = strconv.Itoa(rand.Int())
	accoutAttrs.CustomerID = strconv.Itoa(rand.Int())
	accoutAttrs.Iban = "GB28NWBK40030212764204"
	accoutAttrs.Bic = "NWBKGB42"
	accoutAttrs.AccountClassification = "Personal"
	accoutAttrs.Name = name
	inputAccount.Attributes = &accoutAttrs

	resp := form3cli.CreateAccount(inputAccount)

	if resp.Status != "Success" {
		t.Errorf("Error in create account")
		t.Fail()
	}

	if resp.Data != nil {
		generatedId = resp.Data[0].ID
		log.Println("Generated Id ---" + resp.Data[0].ID)
	}

}

func TestFetchAccount(t *testing.T) {
	resp := form3cli.FetchAccount(generatedId)
	if resp.Status != "Success" {
		t.Errorf("Error in fetch account")
	}

}

func TestDeleteAccount(t *testing.T) {
	resp := form3cli.DeleteAccount(generatedId)
	if resp.Status != "Success" {
		t.Errorf("Error in delete account")
	}

}

func TestCreateAccountFail(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	var inputAccount form3_client.AccountData
	inputAccount.ID = uuid.New().String()
	inputAccount.Type = "accounts"
	inputAccount.OrganisationID = uuid.New().String()

	var name = []string{"First", "Name", "Last", "Name"}

	var accoutAttrs form3_client.AccountAttributes
	accoutAttrs.Country = "GB"
	accoutAttrs.BankID = "400302"
	accoutAttrs.BankIDCode = "GBDSC"
	accoutAttrs.AccountNumber = strconv.Itoa(rand.Int())
	accoutAttrs.CustomerID = strconv.Itoa(rand.Int())
	accoutAttrs.Iban = "GB28NWBK40030212764204"
	accoutAttrs.Bic = "NWBKGB42"
	accoutAttrs.AccountClassification = "personal"
	accoutAttrs.Name = name
	inputAccount.Attributes = &accoutAttrs

	resp := form3cli.CreateAccount(inputAccount)
	if resp.ErrorCode != 400 {
		t.Errorf("Creating account with invalid data.")
		t.Fail()
	}
}

func TestFetchAccountFail(t *testing.T) {
	resp := form3cli.FetchAccount("123123213")
	if resp.ErrorCode != 400 {
		t.Errorf("Error in fetch account")
	}

}

func TestFetchAccountFailParsing(t *testing.T) {
	resp := form3cli.FetchAccount("")

	if resp.Status == "Failure" && resp.ErrorMessage != "Error while unmarshalling response body" {
		t.Errorf("Error in fetch account")
	}
}

func TestFetchAllAccount(t *testing.T) {
	resp := form3cli.FetchAllAccounts()
	if resp.StatusCode != 200 && len(resp.Data) > 0 {
		t.Errorf("Error in fetch account")
	}
}

func TestDeleteAccountFail(t *testing.T) {
	resp := form3cli.DeleteAccount("")
	if resp.ErrorCode != 404 {
		t.Errorf("Error in delete account")
	}

}
