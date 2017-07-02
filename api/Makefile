local:vendor
	CGO_ENABLED=0 goat build -o api && strip -s api
deps:vendor
	CGO_ENABLED=0 goat deps
clean:
	rm -f api
