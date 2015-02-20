# url

[![Build Status](https://travis-ci.org/mong-go/url.svg?branch=master)][ci]
[![GoDoc](https://godoc.org/github.com/mong-go/url.v2?status.svg)][gd]

  [ci]: https://travis-ci.org/mong-go/url
  [gd]: http://godoc.org/github.com/mong-go/url.v2

MongoDB URL parser

## Install

    go get gopkg.in/mong-go/url.v1

## Usage

      u := url.Parse("mongodb://user:secret@example.com:12345/database")

      // u.Database() => "database"
      // u.User()     => *url.UserInfo
      // u.Scheme()   => "mongodb"
      // u.Host()     => "example.com"
      // u.Port()     => "12345"


## License

MIT
