# MyPayment

Project of new payment system. This service is adopting the Ardan Labs' service template and utilizing mux as the router.
The service consists of three endpoints, which are POST /validation, Post /transfers, and Post /transfers/status. In terms of completing this service,
this project used 3rd party API acted as BANK API. The mock BANK API divided into 2 parts which are:
- Validation API to validate the accounts.
``url: https://65f19893034bdbecc7631ed1.mockapi.io/mock/bank/:validate``
- Transfer API to transfer the money.
``url: https://65f19893034bdbecc7631ed1.mockapi.io/mock/bank/:transaction``

### Prerequisites

Install the following dependencies

```
go get ./...
```

### Run

To run the application, execute the following command

```
make run-local
```

### Endpoints

- GET /validation/accounts
This endpoint is responsible for validating the accounts based on provided account number and name.

Request example:
```
http://localhost:3002/validation/accounts?name=Clifton Franey&bank_number=522628902
```

Respond example:
```
{
    "account_name": "Clifton Franey",
    "account_number": "522628902",
    "is_valid": true
}
```
- POST /transfers
This endpoint is responsible for transferring the money from one account to another account. 
In the process this endpoint will insert the transaction first and then call the mock Bank API to send the transaction.

Request example:
``` json
http://localhost:3002/transfers

{
    "sender_name":"may",
    "sender_bank_number":"12345",
    "sender_bank":"BNI",
    "receiver_name":"april",
    "receiver_bank_number":"123456789",
    "receiver_bank": "BCA",
    "amount": "10000",
    "currency": "IDR",
    "description":"payment"
}

```

Request example to Mock Bank API:
Sending the MyPayment ID to mock Bank API as unique identifier to update the payment status later on.
``` json
http://localhost:3002/transfers

{   
    "id":"c27da6b8-b9cd-48a5-9a66-65b8e747d0dc"
    "sender_name":"may",
    "sender_bank_number":"12345",
    "sender_bank":"BNI",
    "receiver_name":"april",
    "receiver_bank_number":"123456789",
    "receiver_bank": "BCA",
    "amount": "10000",
    "currency": "IDR",
    "description":"payment"
}

```

Respond example:
```json
{
    "id": "c27da6b8-b9cd-48a5-9a66-65b8e747d0dc",
    "sender_name": "may",
    "sender_bank_number": "12345",
    "sender_bank": "BNI",
    "receiver_name": "april",
    "receiver_bank_number": "123456789",
    "receiver_bank": "BCA",
    "amount": "10000",
    "currency": "IDR",
    "status": "PENDING",
    "description": "payment",
    "created": "2024-03-14T17:42:44.446131+08:00",
    "updated": "2024-03-14T17:42:44.446131+08:00"
}

```

- POST /transfers/status
This endpoint is responsible for receiving the status update from bank and update the status in MyPayment DB.

Request example:
``` json
http://localhost:3002/transfers/status

{
    "id":"c27da6b8-b9cd-48a5-9a66-65b8e747d0dc",
    "status":"SUCCESS"
}

```

Response example:

```json
{
  "id": "c27da6b8-b9cd-48a5-9a66-65b8e747d0dc",
  "status": "SUCCESS",
  "updated": "2024-03-14T17:42:44.446131+08:00"
}

```
