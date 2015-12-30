# Strange issue with go + CGo + freebsd file

The go compiler doesn't seem to properly compile a file with the `freebsd`
suffix if it contains `import "C"`.

```
GOOS=freebsd GOARCH=amd64 go run ./main.go
main.go:3:8: no buildable Go source files in /Users/john/go/src/github.com/j16r/cgo_platform_file_issue/pkg
make: *** [default] Error 1
```
