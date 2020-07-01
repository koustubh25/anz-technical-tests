# ANZ Technical Tests

The solution to both the tests is included in this repo.

## I. Test 1

The solution to this is present in the folder `test-1`.
To run the `Dockerfile`, you can run this command
```
cd test-1 
docker build --no-cache -t anz-test-1  .  && docker run --rm -it  --name anz-test-1 -p 8000:8000  anz-test-1:latest
```

This will build and run the program in `main.go`. Once you see `starting http server` in the logs, the http server is ready to serve.

The following three endpoints are available:
```
1. http://localhost:8000/
2. http://localhost:8000/go
3. http://localhost:8000/opt
```
You can hit any one of the them to see the output

```
curl http://localhost:8000/
Hello, world.

curl http://localhost:8000/go
Don't communicate by sharing memory, share memory by communicating.

curl http://localhost:8000/opt
Don't communicate by sharing memory, share memory by communicating.
```

#### Changes

##### 1. Dockerfile

The Dockerfile includes the following changes
1. **Multi-stage docker build**
This allows us to build a go binary in an environment with the necessary tools for static compilation and then simply copy this binary in a `scratch` container which is the bare minimum. We can then simply distribute this `scratch` container which contains the binary.
2. **Use the `scratch` container in the final stage** 
Since the scratch container by itself is empty, the final image size and the attack surface area is greatly reduced.
e.g the final image size is just 8.2 MB which is mainly the binary.
3. **Removal of unnecessary packages (e.g. git)**
Tools like `git` are not needed for the functionality that we want. So, it has been removed
4. **Run as a non-privileged user**
This is important so that anyone with access to the container can't run anything undesirable in it.

##### 2. main.go
This is one small change `Addr:         ":8000",` so that the web server is accessible outside the docker network.

## II. Test 2

The application as well as the CI pipeline configuration code is placed in the `test-2` repo. The choice of tool for CI is [Concourse](https://concourse-ci.org/) and the tool used for deployment is [Flux](https://github.com/fluxcd/flux) which allows to use [GitOps](https://www.weave.works/technologies/gitops/).

**APP**

The app code is written in the `version.go` file. The version of the app is available through the `/version` endpoint. The app simply displays the contents of the `resource/version.json` file. It is the responsiblity of the CI pipeline to keep track of the appropriate versions and make modifications to the `version.json` file inside the docker container.

Once the docker container is ready. The CD tool will automatically deploy the docker image in Kubernetes cluster.

**Resources**

Public API implementation - http://130.211.117.181/version
Github repository (app) - https://github.com/koustubh25/anz-technical-tests/tree/master/test-2
Github repository (releases) - https://github.com/koustubh25/anz-technical-tests-releases


The flow is as shown below:

![logo](test-2/anz.svg)


1. As soon as a developer makes an update to the `anz-technical-tests` repo, a  concourse pipeline is triggered. You can see the status changed on the commits page on github https://github.com/koustubh25/anz-technical-tests/commits/master like this 

![status](test-2/images/github-status.png)

2. This pipeline 