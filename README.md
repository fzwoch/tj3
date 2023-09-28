# Go JPEG Encode and Decode via TurboJPEG v3

Simple JPEG Encoder and Decoder.

Replace
```Go
import(
    "image/jpeg"
)
```

with

```Go
import(
    jpeg "github.com/fzwoch/tj3"
)
```

or, if you only want the transparent decoding feature:

```Go
import(
    "image"
    _ "image/jpeg"
)
```

with

```Go
import(
    "image"
    _ "github.com/fzwoch/tj3"
)
```

## Cgo

At build time of your application you will need to tell the Go compiler where to find the TurboJPEG header and library files. Check Go's [cgo](https://pkg.go.dev/cmd/cgo) documentation for more details.

```shell
export CGO_CFLAGS="-I /path/to/libjpeg-turbo/header-files"
export CGO_LDFLAGS="-L /path/to/libjpeg-turbo/library"
```
