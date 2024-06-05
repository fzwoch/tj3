# Go JPEG Encode and Decode via TurboJPEG v3

Simple JPEG Encoder and Decoder.

Replace
```Go
import (
    "image/jpeg"
)
```

with

```Go
import (
    jpeg "github.com/fzwoch/tj3"
)
```

or, if you only want the transparent decoding feature:

```Go
import (
    "image"
    _ "image/jpeg"
)
```

with

```Go
import (
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

For example, if your system has `pkg-config` here is how to run tests.

```shell
CGO_CFLAGS=`pkg-config --cflags libturbojpeg` CGO_LDFLAGS=`pkg-config --libs libturbojpeg` go test -v
```

## Benchmark

```
BenchmarkDecodeYUV420_TurboJPEG-16    	    9301	    148467 ns/op
BenchmarkDecodeYUV420_GoJPEG-16       	    1293	    818696 ns/op
BenchmarkEncodeYUV420_TurboJPEG-16    	    7362	    155416 ns/op
BenchmarkEncodeYUV420_GoJPEG-16       	    1129	    972425 ns/op
```
