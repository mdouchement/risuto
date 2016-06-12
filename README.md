# Risuto

Simple wishlist web application

## Requirements

- Golang 1.6.x

## Installation

```bash
$ go get github.com/Masterminds/glide
$ glide install
```

## Usage

- **development**

```bash
$ go run ritsu.go -p 5000

# Before pushing to Github
$ find . -name '*.go' -not -path './vendor*' -exec go fmt {} \;
```

- **production**

```bash
# Using Docker
$ docker run -p 5000:5000 -d mdouchement/risuto
# or
$ docker run -v /data:/data -p 5000:5000 -d mdouchement/risuto

# Baremetal

$ go build risuto.go
$ ./risuto -p 5000
```
> Environment variables https://github.com/mdouchement/risuto/blob/master/config/config.go


## Licence

MIT. See the [LICENSE](https://github.com/mdouchement/risuto/blob/master/LICENSE) for more details.

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request
