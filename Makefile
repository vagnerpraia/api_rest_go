
api_gorilla_max: clean
	go get -u github.com/valyala/fasthttp
	go build

clean:
	rm -f api_rest_go