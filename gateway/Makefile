IMAGE_TAG=latest

abigen:
	abigen --pkg main --sol ../contract/contracts/Web3Avatar.sol --out web3avatar_contract.go

build:
	go build -o main

docker-build:
	docker build -t ghcr.io/blackstorm/web3-avatar/gateway:$(IMAGE_TAG) .

docker-push:
	docker push ghcr.io/blackstorm/web3-avatar/gateway:$(IMAGE_TAG)

kube-render:
	kustomize build ./kube | sed 's@X-IMAGE@'"ghcr.io/blackstorm/web3-avatar/gateway:$(IMAGE_TAG)"'@' | sed 's@X-INFURA_TOKEN@'"$(INFURA_TOKEN)"'@' > kube.yaml