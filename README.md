# Steps for Running the Receipt Processor Server:
*Assumes you already have Go installed on your system.*
1. Open your terminal / command line
2. Clone this repo:
`git clone https://github.com/d-halverson/Receipt-Processor.git`
3. Change into the directory of the newly cloned repo:
`cd Receipt-Processor`
4. Run the program:
`go run main/main.go`

# What does this Receipt Processor do?
This Go server implements the [Fetch Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge). You may consume its two endpoints to process receipts, store them in memory, and retrieve how many points each one is worth. For more information on how points are calculated, visit the previous link.

# Now that I have the server running, how do I consume this service's apis?
The above link also contains an overview of the `GET /receipts/{id}/points` and `POST /receipts/process` endpoints, as well as a detailed [api spec](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/api.yml).

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
