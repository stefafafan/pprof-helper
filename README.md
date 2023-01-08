# pprof-helper
A helper for using Go's [pprof package](https://pkg.go.dev/net/http/pprof) with web applications. For now, it just prints out pseudocode.

## Installation
```sh
go install github.com/stefafafan/pprof-helper
```

## Usage
Go to the directory with `go.mod` and call the helper. The helper will go through your `go.mod` and for now print out pseudocode for 

```sh
$ cd your-go-project
$ pprof-helper
For goji, update the code like this:
import (
	"net/http/pprof"
)

func main() {
	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/debug/pprof/"), pprof.Index)
	mux.HandleFunc(pat.Get("/debug/pprof/cmdline"), pprof.Cmdline)
	mux.HandleFunc(pat.Get("/debug/pprof/profile"), pprof.Profile)
	mux.HandleFunc(pat.Get("/debug/pprof/symbol"), pprof.Symbol)
}
```

You can also specify `--go-mod-file` option, which is the path of the `go.mod` instead of moving to that directory.
```sh
$ pprof-helper --go-mod-file=./my-go-project/go.mod
```
