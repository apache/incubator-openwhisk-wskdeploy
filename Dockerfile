FROM golang:1.7

# Install zip
RUN apt-get -y update && \
    apt-get -y install zip emacs

ENV GOPATH=/

# Download and install tools
RUN echo "Installing the godep tool"
RUN go get github.com/tools/godep


RUN echo "Installing the go-bindata tool"
RUN go get github.com/jteeuwen/go-bindata/...



ADD . /src/github.com/

RUN go get github.com/openwhisk/openwhisk-client-go/whisk

RUN cd /src/github.com/openwhisk/openwhisk-client-go/whisk && git fetch && git checkout i18n

RUN git clone http://github.com/openwhisk/openwhisk-wskdeploy /src/github.com/openwhisk/openwhisk-wskdeploy

RUN cd /src/github.com/openwhisk/openwhisk-wskdeploy && git checkout master

# Load all of the dependencies from the previously generated/saved godep generated godeps.json file
RUN echo "Restoring Go dependencies"
RUN cd /src/github.com && /bin/godep restore -v

# Collect all translated strings into single .go module
RUN echo "Generating i18n Go module"
RUN cd /src/github.com && /bin/go-bindata -pkg wski18n -o wski18n/i18n_resources.go wski18n/resources

# Generate a Go package dependency list
# NOTE: Currently, the 'go list' command will not work against the current Go CLI non-standard file structure
#RUN cd /src/github.com/go-whisk-cli && go list -f '{{join .Deps "\n"}}' > ../../../wsk.deps.out
RUN cd /src/github.com/ && echo "Dependencies list requires restructuring the GO CLI packages" > wskdeploy.deps.out


# All of the Go CLI binaries will be placed under a build folder
RUN mkdir /src/github.com/build

ARG WSKDEPLOY_OS
ARG WSKDEPLOY_ARCH

# Build the Go wsk CLI binaries and compress resultant binaries
RUN chmod +x /src/github.com/build.sh
RUN cd /src/github.com/ && ./build.sh
