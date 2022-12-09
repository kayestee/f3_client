package form3_cli

import (
	"github.com/google/uuid"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var generatedId string
var form3cli Form3_API

func TestCreateAccount(t *testing.T) {
	form3cli.Api_host_url = os.Getenv("API_HOST")
	form3cli.Api_host_version = os.Getenv("API_VERSION")

	rand.Seed(time.Now().UnixNano())
	var inputAccount AccountData
	inputAccount.ID = uuid.New().String()
	inputAccount.Type = "accounts"
	inputAccount.OrganisationID = uuid.New().String()

	var name = []string{"First", "Name", "Last", "Name"}

	var accoutAttrs AccountAttributes
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
		log.Println(resp.Data)
		t.Errorf("Error in create account")
		t.Fail()
	}

	if resp.Data.ID != "" {
		generatedId = resp.Data.ID
	}

	log.Println("Generated Id ---" + resp.Data.ID)
}

func TestFetchAccount(t *testing.T) {
	resp := form3cli.FetchAccount(generatedId)
	t.Log(resp)

	if resp.Status != "Success" {
		t.Errorf("Error in fetch account")
	}

}

func TestDeleteAccount(t *testing.T) {
	resp := form3cli.DeleteAccount(generatedId)
	t.Log("Id is ", generatedId)
	t.Log(resp)
	if resp.Status != "Success" {
		t.Errorf("Error in delete account")
	}

}

func TestCreateAccountFail(t *testing.T) {
	form3cli.Api_host_url = os.Getenv("API_HOST")
	form3cli.Api_host_version = os.Getenv("API_VERSION")

	rand.Seed(time.Now().UnixNano())
	var inputAccount AccountData
	inputAccount.ID = uuid.New().String()
	inputAccount.Type = "accounts"
	inputAccount.OrganisationID = uuid.New().String()

	var name = []string{"First", "Name", "Last", "Name"}

	var accoutAttrs AccountAttributes
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
	log.Println(resp)
	if resp.Status != "Failure" {
		log.Println(resp.Data)
		t.Errorf("Creating account with invalid data.")
		t.Fail()
	}
}

func TestFetchAccountFail(t *testing.T) {
	resp := form3cli.FetchAccount("")
	t.Log(resp)
	if resp.Status != "Failure" {
		t.Errorf("Error in fetch account")
	}

}

func TestDeleteAccountFail(t *testing.T) {
	resp := form3cli.DeleteAccount("")
	t.Log("Id is ", generatedId)
	t.Log(resp.Status)
	if resp.Status != "Failure" {
		t.Errorf("Error in delete account")
	}

}
