# CERTI Challenge

## Stack of tecnology

Here is a list of stack of tech used in this project

* Golang version 1.14.1 linux/amd64
* Docker CE
* httpie
* golangci-lint
* Github CI

## Dependencies of project

You need install some libs for running this project.

### Httpie for curl sintact sugar

Linux based on **Debian**, you need first install **httpie**, so you can use `sudo apt install httpie`, and you need a **plugin** for **jwt** request, so you can install with this command `pip3 install -U httpie-jwt-auth`, in my case I'm using python3.

### Jq (JQuery for command line)

On Debin linux, you can install with `sudo apt install jq`

## Test project

In command line inside folder of project, run `go test ./... -v` or makefile `make test`

## Check lint of code

You need the `golangci-lint`, for the installation I recommend you see the documentation about that. Link for doc [Here](https://github.com/golangci/golangci-lint)

## Start project

Run `go run main.go` in command line.
Or you can make the binary files using `makefile build`, this command you going generate all binaries(***raw and zip***) files for **linux**, **macOs** and **raspBery Pie** in folder `out`.

## Generate docker image

For create a image with **Docker-CE** just running in folder of projet the command `docker build -t thiagozs/challenge-certi .`

Runing the image after build `sudo docker run --rm --name=challenge-certi --publish=8080:8080 thiagozs/challenge-certi:latest`

Or just run the `make image` for create a new one, and `make rundocker` for start execute the API

## Docker Healthcheck

On the construction of image we have a change to put a little **healtcheck** on API.

```sh
Step 17/17 : HEALTHCHECK --interval=5s --timeout=2s --start-period=2s --retries=5 CMD [ "curl", "--silent", "--fail", "http://localhost:8080//ping" ]
```

## Testing API with httpie

After you install(if you want) the httpie for a sugar sintax for **CURL**. Execute the command in your terminal.

### Consume information

In your terminal: `http get http://localhost:8080/12345678`

```sh
HTTP/1.1 200 OK
Content-Length: 92
Content-Type: application/json; charset=utf-8
Date: Mon, 18 May 2020 02:00:01 GMT

{
    "extenso": "doze milhões e trezentos e quarenta e cinco mil e seiscentos e setenta e oito"
}
```

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the [LICENSE](https://github.com/thiagozs/extensive/blob/master/LICENCE) file

**2020, thiagozs**
