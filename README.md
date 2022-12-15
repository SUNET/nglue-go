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


### LICESE: MIT

