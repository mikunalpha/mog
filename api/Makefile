package := github.com/mikunalpha/mog/api
output := api

vendor := ${PWD}/vendor
tmpvendor := ${PWD}/.vendor

local:.gopath
	mv $(vendor) $(tmpvendor) && ln -sf $(tmpvendor) ${PWD}/.gopath/src
	env GOPATH=${PWD}/.gopath CGO_ENABLED=0 go build -o $(output)
	mv $(tmpvendor) $(vendor) && ln -sf $(vendor) ${PWD}/.gopath/src

init:
	glide create --skip-import --non-interactive
	echo "ignore:" >> glide.yaml
	echo "- "$(package) >> glide.yaml
	echo "excludeDirs:" >> glide.yaml
	echo "- .git" >> glide.yaml
	echo "- node_modules" >> glide.yaml
	glide install --skip-test --strip-vendor
	mkdir .gopath && ln -sf $(vendor) ${PWD}/.gopath/src
	if [ ! -d "$(vendor)/$(package)" ]; then mkdir -p $(vendor)/$(package); fi
	rm -rf $(vendor)/$(package)
	ln -sf ${PWD} $(vendor)/$(package)

get:
	glide get $(P)
	mv $(vendor) $(tmpvendor) && ln -sf $(tmpvendor) ${PWD}/.gopath/src
	if [ ! -d "$(tmpvendor)/$(package)" ]; then mkdir -p $(tmpvendor)/$(package); fi
	rm -rf $(tmpvendor)/$(package)
	ln -sf ${PWD} $(tmpvendor)/$(package)
	mv $(tmpvendor) $(vendor) && ln -sf $(vendor) ${PWD}/.gopath/src

rm:
	glide rm $(P)
	mv $(vendor) $(tmpvendor) && ln -sf $(tmpvendor) ${PWD}/.gopath/src
	if [ ! -d "$(tmpvendor)/$(package)" ]; then mkdir -p $(tmpvendor)/$(package); fi
	rm -rf $(tmpvendor)/$(package)
	ln -sf ${PWD} $(tmpvendor)/$(package)
	mv $(tmpvendor) $(vendor) && ln -sf $(vendor) ${PWD}/.gopath/src

update:
	glide up
	mv $(vendor) $(tmpvendor) && ln -sf $(tmpvendor) ${PWD}/.gopath/src
	if [ ! -d "$(tmpvendor)/$(package)" ]; then mkdir -p $(tmpvendor)/$(package); fi
	rm -rf $(tmpvendor)/$(package)
	ln -sf ${PWD} $(tmpvendor)/$(package)
	mv $(tmpvendor) $(vendor) && ln -sf $(vendor) ${PWD}/.gopath/src

clean:
	rm -rf glide.yaml glide.lock .gopath vendor $(output)