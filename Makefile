.DEFAULT_GOAL := run

init:
	@sh ./reflesh.sh
	@make bind
	@go install

run:
	@sh ./reflesh.sh
	@make bind
	@go install

bind:
	@go-bindata -o bindata/bindata.go --pkg bindata clarch/templates/...
