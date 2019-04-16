# Calvin's trading bot in Golang
## This trading bot support both manual and api trading for exchange:
```coss.io```

## Configuration and Pre-requisite
* Go go1.10.5

### Configure GOPATH


## How to use
### Run on PC/Mac
* Pull the source code
* Configure GOPATH
```export GOPATH=$HOME/trading_bot:$GOPATH```
```cd src/cyan.io/trading_bot```
```go get ./```
```go build```
```bin/trading_bot```

### Run on Docker (Recommended)
* Pull the source code
* Follow the instruction in Dockerfile
```docker build -t bot .```
```docker run -it --rm -p 8080:8080 --name trading_bot bot```
