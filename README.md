# validate

[![Build Status](https://travis-ci.org/northbright/validate.svg?branch=master)](https://travis-ci.org/northbright/validate)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/validate)](https://goreportcard.com/report/github.com/northbright/validate)
[![GoDoc](https://godoc.org/github.com/northbright/validate?status.svg)](https://godoc.org/github.com/northbright/validate)

Package validate provides functions to validate phone number, username and password for Chinese users.

#### Username validation
  * Chinese Characters supported.
  
      Each Chinese character's length is recognized as 2.
      Because the font width of Chinese character is 2x than Latin(number) in most case.
      e.g.
      "中文汉字" -> 4 Chinese Characters: display width = 8
      "abcd1234" -> 8 Latin chars and numbers mixed: display width = 8.5      

  * Min / Max length. Default: 6 - 16.
  * If can have numbers. Default: true.
  * If can have hyphens(`-`). Default: true.
  * If can have underscore(`_`). Default: true.

#### Password validation
  * Min / Max length. Default: 8 - 64.
  * If have at least one number. Default: true.
  * If have at least one upper-case letter. Default: true.
  * If have at least one lower-case letter. Default: true.
  * If have at least one special character(one symbol or one puncuation). Default: true.

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/validate)

#### License
* [MIT License](LICENSE)
