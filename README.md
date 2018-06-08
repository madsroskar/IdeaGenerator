# IdeaGenerator

![travis](https://travis-ci.com/madsroskar/IdeaGenerator.svg?branch=master)

This is a project I've created to mess around with Go and Docker. It generates "ideas" formatted as:

> {Existing service}, but for {target audience or platform}.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

To run this program you need:

* [Go](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/install/)

### Installing

Assuming the prerequisites are met, installation is as simple as cloning this repository, building the docker image, and running it:

Clone the repository:

```
$ git clone git@github.com:madsroskar/IdeaGenerator.git
$ cd IdeaGenerator
```

Build the docker image:

```
$ docker build ./
```

Get the Docker image ID:

```
$ docker images
```

Run the image in development mode:

```
$ docker run -it -p 8000:8000 -v $(pwd)/src:/go/src/github.com/madsroskar/IdeaGenerator/src <Image ID>
```

On changes, the code will hot relaunch, so you won't have to build the image every time there's a change.

## Running the tests

The tests can be run with the shell script [test.sh](./test.sh).g

## Built With

* [gorilla/mux](https://github.com/gorilla/mux) - Multilpexer
* [Docker](https://www.docker.com/) - Building and running the code
* [fresh](https://github.com/pilu/fresh) - Hot reloading of code

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/madsroskar/IdeaGenerator/tags). 

## Authors

* **Mads Røskar** - [madsroskar](https://github.com/madsroskar)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
