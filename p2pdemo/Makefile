all:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client/main github.com/DemoLiang/demo/p2pdemo/client
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server/main github.com/DemoLiang/demo/p2pdemo/server