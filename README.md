# Risuto

Simple wishlist web application

## Technologies

- [Golang](https://golang.org/)
- [Vue.js](https://vuejs.org/)
- [Bulma](https://bulma.io/) (Formally [Buefy](https://buefy.github.io) wrapper)

## Requirements

- Golang 1.9.x

## Installation

```bash
$ go get -u github.com/gobuffalo/packr/...
$ go get github.com/Masterminds/glide
$ glide install
```

## Usage

- **development**

```bash
# Download assets
$ go run risuto.go fetch

# Run app server
$ make live-reload
# old fashion
$ go run risuto.go server -p 5000

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

$
$ packr build risuto.go
$ ./risuto server -p 5000
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
