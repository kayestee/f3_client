# Form3 Take Home Exercise
This client library is written for assessment of form3 API. 

## Instructions
Include github.com/kayestee/f3_client in your go.mod file as require dependencies.
And run go mod download 

## API's list
The client api currently utilizes Create, Fetch and delete functions of form3 API.

Here is a sample JSON for create API request:
Reference:: https://www.api-docs.form3.tech/api/tutorials/getting-started/send-a-payment/create-the-payment

``` {
  "data": {
    "id": "{{random_guid}}",
    "organisation_id": "{{organisation_id}}",
    "type": "accounts",
    "attributes": {
       "country": "GB",
        "base_currency": "GBP",
        "bank_id": "400302",
        "bank_id_code": "GBDSC",
        "account_number": "10000004",
        "customer_id": "234",
        "iban": "GB28NWBK40030212764204",
        "bic": "NWBKGB42",
        "account_classification": "personal",
        "name" : ["Test Account"]
    }
  }
}
```

Fetch and Delete uses the account Id created from the above API call.

