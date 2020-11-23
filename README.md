# GoStream

Stream utilities for Go.

![Version](https://img.shields.io/badge/Release-v0.1.2-brightgreen.svg?style=flat-square)
[![Reference](https://img.shields.io/badge/Go-Reference-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/deadblue/gostream)
![License](https://img.shields.io/:License-MIT-green.svg?style=flat-square)

## Packages

### quietly

Wraps some functions, without returning its error. Caller should use them when he explicitly know there is no error, or he really does not care the error.

Example:

```go
file, err := os.Open("/path/to/file")
if err != nil {
    panic(err)
}
defer quietly.Close(file)
```

### chain

Links multiple io.Reader into one, and read them one by one.

Example:

```go
r1 := strings.Reader("Hello,")
r2 := strings.Reader("world!")

r := chain.JoinReader(r1, r2)
content, _ := ioutil.ReadAll(r)
```

### multipart

Provides a read-on-demand multipart form, which is always used to upload files through HTTP POST request.

Example:

```go
form := multipart.New().
    AddValue("foo", "1").
    AddValue("bar", "2")
file, err := os.Open("/path/to/file")
form.AddFile("upload", file)

req, err := multipart.NewRequest("http://server/upload", form)
if err != nil {
    panic(err)
}
resp, err := client.Do(req)
```

### observe

Provides observed `io.ReadCloser` and `io.WriteCloser` that caller can monitor the transfer progress.

Example:

```go
file, err := os.Open("/path/to/file")
if err != nil {
    panic(err)
}

r := observe.Reader(file, &YourObserver{})
ioutil.ReadAll(r)
defer quietly.Close(r)
```

### binary

Read/write binary data on a stream.

Example:

```go
import github.com/deadblue/gostream/binary

buf := bytes.NewReader([]byte{
    0x12, 0x34, 0x56, 0x78,
})
br := NewReader(buf, LittleEndian)
if u, err := br.ReadUint32(); err != nil {
    log.Fatal(err)
} else {
    log.Printf("Uint value: 0x%x", u)
    // Output: "Uint value: 0x78563412"
}
```

## License

MIT