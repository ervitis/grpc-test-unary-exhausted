# Testing grpc unary purpose

## Testing

Run the server:

```bash
cd server/cmd

go run main.go
```

Run the client:

```bash
cd client/cmd

go run main.go
```

### Results

Reading a big file using unary client-server throws error:

```
panic: rpc error: code = ResourceExhausted desc = grpc: received message larger than max (10485765 vs. 4194304)

goroutine 1 [running]:
main.main()
	/github.com/ervitis/grpc-test-unary-exhausted/client/cmd/main.go:38 +0x351
exit status 2
```