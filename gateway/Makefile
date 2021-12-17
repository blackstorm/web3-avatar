abigen:
	abigen --pkg main --sol ../contract/contracts/Web3Avatar.sol --out web3avatar_contract.go

build:
	go build -o main

docker-build:
	docker build -t w3a/gateway .