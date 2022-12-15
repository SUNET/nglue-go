## nglue implementation in golang

Required golang version: `1.19`. You can [download](https://go.dev/dl/) from
Upstream.

As the code still uses `libc` from the system, please build it on the
OS/distribution you will use.

### How to build?

```
make build
```

This will create a binary called `nglue` in the current directory. By default
it tries to connect to `http://localhost:8000/`, but you can point it to
another location using environment variable `NGLUE_SERVER`.


### How to build in a container?

First let us create the container for building the project.

```
podman build -t local/nglue-builder .
```

Then build the binary

```
podman run --rm -it -v $PWD:/nglue-go:z local/nglue-builder bash
```

Or via docker:

```
docker run --rm -it -v $PWD:/nglue-go local/nglue-builder bash
```


### LICESE: MIT

