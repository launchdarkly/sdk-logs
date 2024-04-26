.PHONY: logdrkly

logdrkly:
	cd tools/logdrkly && go build .

logdrkly-fmt:
	cd tools/logdrkly && go fmt .

all: logdrkly
