# client-demo
client-demo demonstrates how to use the client-api.
- The mock server must be running.
- Uses the mock data in the mock-datastore
- demo1 Authenticates the user
```
go run \demo1\main.go 
```

- demo2 use the beartoken from demo1
```
go run demo2\main.go -AccessToken="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MzI2NjIzODEsInVzZXJuYW1lIjoiS2V2aW5LZWxjaGUifQ.xNBn9Km_Ps8UMAGx0QP9dGdJ3Bu75COVWOrjktTjQcQ"
```
