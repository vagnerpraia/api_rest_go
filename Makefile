
api_gorilla_max: clean
	go get -u github.com/valyala/fasthttp
	go get -u github.com/globalsign/mgo
	go build -o ./bin/api_rest_go ./src/*.go

clean:
	rm -f ./bin/*