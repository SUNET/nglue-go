## nglue implementation in golang

Required golang version: `1.19`. You can [download](https://go.dev/dl/) from Upstream.

### How to build?

```
make build
```

This will create a binary called `nglue` in the current directory. By default
it tries to connect to `http://localhost:8000/`, but you can point it to
another location using environment variable `NGLUE_SERVER`.


### LICESE: MIT

