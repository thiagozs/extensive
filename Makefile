GOCMD=go
GOBUILD=$(GOCMD) build
GOENV=$(GOCMD) env
GREP=grep
OUTDIR=out
VERSION=1.0.0
LDFLAGS=-ldflags "-X main.version=${VERSION}"
NAME=server
MAIN=main.go

build:
	@rm -fr ${OUTDIR}
	@mkdir -p ${OUTDIR}
	GOOS=linux GOARCH=arm GOARM=6 ${GOBUILD} ${LDFLAGS} -o ${OUTDIR}/${NAME}.rpi ${MAIN}
	GOOS=linux ${GOBUILD}  ${LDFLAGS} -o ${OUTDIR}/${NAME}.lin ${MAIN}
	GOOS=darwin ${GOBUILD} ${LDFLAGS} -o ${OUTDIR}/${NAME}.mac ${MAIN}
	zip ${OUTDIR}/${NAME}.rpi.zip ${OUTDIR}/${NAME}.rpi 
	zip ${OUTDIR}/${NAME}.lin.zip ${OUTDIR}/${NAME}.lin
	zip ${OUTDIR}/${NAME}.mac.zip ${OUTDIR}/${NAME}.mac

alpine:
	GOOS=linux ${GOBUILD}  ${LDFLAGS} -o ${OUTDIR}/${NAME}.lin ${MAIN}

test:
	go test ./... -v

image:
	sudo docker build -t thiagozs/challenge-certi .

rundocker:
	sudo docker run --rm --name=challenge-certi --publish=8080:8080 thiagozs/challenge-certi:latest

