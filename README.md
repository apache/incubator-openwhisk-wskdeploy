# Whisk Deploy `wskdeploy`

`wskdeploy` is a utility to help you create and deploy OpenWhisk projects. Deploy all your actions, triggers, and rules using a single command! You can use this in addition to the OpenWhisk CLI.

`wskdeploy` is currenty under development and in its very early stages.  Check back often to see how its progressing.

# How to use
`wskdeploy` is written in Go. You can invoke it as a Go program, or run its binary file `wskdeploy` directly after building.

Using command
```
$ go run main.go --help
```
or
```
$ ./wskdeploy --help
```
, you can get the detail usage of this tool.


For example,
```
$ go run main.go -m /tests/testcases/triggerrule/manifest.yml -d /tests/testcases/triggerrule/deployment.yml
```
or
```
$ ./wskdeploy -m /tests/testcases/triggerrule/manifest.yml -d /tests/testcases/triggerrule/deployment.yml
```
will deploy the `triggerrule` test case.

# How to Build on local host.
`wskdeploy` can be built with Go tool.

1. Setup your [Go development environment](https://golang.org/doc/code.html).

2.  `wskdeploy` depends on the `github.com/openwhisk/openwhisk-client-go/whisk` . To install:

``` go get github.com/openwhisk/openwhisk-client-go/whisk ```

3. Clone this repo into `$GOPATH/src/github.com/openwhisk`, which should have been created by Step #2.

```
$ cd $GOPATH/src/github.com/openwhisk
$ git clone http://github.com/openwhisk/wskdeploy
```

4. Tagged releases are in master. The latest build is always in the development branch. Inside `$GOPATH/src/github.com/openwhisk/wskdeply`:

```
$ git checkout development   ## or skip this step and just build master
$ go get -t -d ./...
$ go build

5. If you want to build with the godep tool, please execute the following commands.

```
$ go get github.com/tools/godep ## Install the godep tool.
$ godep get                     ## Download and install packages with specified dependencies.
$ godep go build                ## build the wskdeploy tool.
```

Note: we have no releases yet so you should build the `development` branch.

# How to Build with Docker.
If you don't want to bother with go installation, build, git clone etc, you can do it with Docker, then
you can run wskdeploy tool in your container.

1. First you need a docker daemon running locally on your machine or whatever in a VM etc.

2. Get the Docker file.
 ```
 wget -O Dockerfile https://raw.githubusercontent.com/openwhisk/wskdeploy/master/Dockerfile
 ```

3. Build and tag a docker image.
```
docker build -f Dockerfile .  -t openwhisk/wskdeploy
```

4. Bring up the docker container.
```
docker run -ti openwhisk/wskdeploy
```
5. Inside the container, run `wskdeploy` and have fun.

Note: Based on user role, you may need add sudo before your command to run as root.
