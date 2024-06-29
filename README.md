# What is this Receipt Processor?
This Go server implements the [Fetch Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge). You may consume its two endpoints to process receipts, store them in memory, and retrieve how many points each one is worth. For more information on how points are calculated, visit the previous link.

# Steps for Running the Receipt Processor Server:
*Assumes you already have Go installed on your system.*
1. Open your terminal / command line
2. Clone this repo:
```
git clone https://github.com/d-halverson/Receipt-Processor.git
```
4. Change into the directory of the newly cloned repo:
```
cd Receipt-Processor
```
6. Run the program:
```
go run main/main.go
```

# Now that I have the server running, how do I consume this service's apis?
The [Fetch Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge) link also contains an overview of the `GET /receipts/{id}/points` and `POST /receipts/process` endpoints, as well as a detailed [api spec](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/api.yml).

## Postman
If you have Postman installed, you may use this button below to use the collection I have created to do basic interaction with the service:
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/13928979-6f22523b-9509-4503-b99f-f0ce192a3c22?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D13928979-6f22523b-9509-4503-b99f-f0ce192a3c22%26entityType%3Dcollection%26workspaceId%3D64452a18-21f5-46d9-a5de-751bdc34fe83)

## Curl
1. ProcessReceipt Endpoint:
Example Request:
```
curl --location --request POST 'http://localhost:8080/receipts/process' \
--header 'Content-Type: text/plain' \
--data-raw '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}'
```
Example Response:
```
{"id":"c163bab9-230f-4555-9e0c-90b33a9841c9"}
```
2. GetPoints Endpoint:
Example Request:
*Replace {id} with the id returned in the ProcessEndpoint response*
```
curl --location --request GET 'http://localhost:8080/receipts/{id}/points'
```
Example Response:
```
{"points":28}
```

# Implementation Details and Thoughts
## Concurrency
I implemented the `ReceiptStorage` struct with concurrency in mind using locks around reads and writes. Right now, there isn't a huge need for this, because if the API consumers only call GetPoints with a real id they have from a previous call, they know the returned points will always be the same because there will not be any updates to this receipt id in the future (subsequent POSTs of the same receipt will create separate ids). Since this is the case, even without the usage of locks in `ReceiptStorage` the GetPoints API would still have been accurate. I chose to implement it with locks though, because this allows further expansion of features for the service in the future. If there is ever a need to update a receipt's contents or delete a receipt entirely, concurrency would become an absolute _must_.

## Package Structure
I separated my code into the following packages:
- main -> Has code to execute the server and start listening for requests
- handlers -> Contains API handler functions
- models -> Contains structs for input and output formats of the APIs, and validation of input
- points -> Logic to calculate points for a receipt
- storage -> Logic to store receipts in a thread safe manner

Even though some of the packages do not have a lot of code in them, I still chose to follow this structure because it allows for further code to be added on more easily in the future.

## Unit Testing
I created unit testing on all areas of the service's code, especially extensively on the validation of inputted receipts' formats. To run all unit tests, run this command in the repo's root:
```
go test ./...
```



