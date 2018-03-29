.DEFAULT_GOAL := run

init:
	@sh ./reflesh.sh
	@go install
	@clarch --mode init

run:
	@sh ./reflesh.sh
	@go install
	@clarch

bind:
	@go-bindata -o bindata/bindata.go --pkg bindata clarch/templates/...
