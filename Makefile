build:
	rm -rf dist && mkdir dist
	go build -o ./dist ./cmd/...

abigen:
	mkdir -p internal/contracts/erc20
	abigen --abi=internal/contracts/abi/ERC20.json --pkg=erc20 --out=internal/contracts/ERC20/ERC20.go
	mkdir -p internal/contracts/lockinginfo
	abigen --abi=internal/contracts/abi/LockingInfo.json --pkg=lockinginfo --out=internal/contracts/lockinginfo/LockingInfo.go
	mkdir -p internal/contracts/lockingpool
	abigen --abi=internal/contracts/abi/LockingPool.json --pkg=lockingpool --out=internal/contracts/lockingpool/LockingPool.go
	mkdir -p internal/contracts/sequencerset
	abigen --abi=internal/contracts/abi/MetisSequencerSet.json --pkg=sequencerset --out=internal/contracts/sequencerset/MetisSequencerSet.go
