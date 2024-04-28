build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./output/customer_svc && docker build -t customer_svc:0.1 .