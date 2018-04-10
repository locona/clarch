.DEFAULT_GOAL := run

init:
	@sh ./reflesh.sh
	@make bind
	@go install
	@clarch

run:
	@sh ./reflesh.sh
	@make bind
	@go install
	@clarch

bind:
	@go-bindata -o bindata/bindata.go --pkg bindata clarch/templates/...
