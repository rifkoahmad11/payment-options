# Payment Options App

A simple Go web service to simulate multiple payment methods for implementing goroutine-based parallel processing.

## Run the App
```bash
go mod init payment-options
go mod tidy
go run cmd/main.go
```

The server will start on `http://localhost:8081`

## Test the Endpoint
You can test using `curl`:
```bash
curl http://localhost:8081/payment/options
```

### Example Response
```json
{
  "returnCode": "200",
  "returnDesc": "success",
  "data": {
    "ovo": {
      "account": "6288xx",
      "status": "Active",
      "balance": "10000",
      "icon": "https://sampleurl.com/ovo.jpg"
    },
    "dana": {
      "account": "6288xx",
      "status": "Active",
      "balance": "10000",
      "icon": "https://sampleurl.com/dana.jpg"
    },
    ...
  }
}
```

## License
MIT