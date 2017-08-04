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

#### References
* [为什么淘宝和京东会员注册的时候用户名一个汉字算两个字符，有什么考虑吗？](https://www.zhihu.com/question/22295828/answer/82576462)

#### License
* [MIT License](LICENSE)

-------------------------

validate包提供了适用于中国用户的函数用于验证手机号码，用户名以及密码。

#### 用户名验证
  * 支持中文作为用户名。
      
          每个中文字符的长度将被计算成2。
          因为在通常情况下，中文字符的字体的显示宽度是拉丁字母（数字)的2倍。
          例如：
          "中文汉字" -> 4个中文字符:显示宽度 ＝ 8
          "abcd1234" -> 8个拉丁字母和数字组合: 显示宽度 = 8.5
  * 最短 ／ 最长长度。默认: 6 - 16。
  * 是否可以包含数字。默认: 是。
  * 是否可以包含连接符(`-`)。默认: 是。
  * 是否可以包含下划线(`_`)。默认: 是。

#### 密码验证
  * 最短 ／ 最长长度。默认: 8 - 64。
  * 是否至少包含1个数字。默认: 是。
  * 是否至少包含1个大写字母。默认: 是。
  * 是否至少包括1个小写字母。默认: 是。
  * 是否至少包括1个特殊字符(1个符号或者标点)。默认: 是。

#### 文档
* [API 文档](http://godoc.org/github.com/northbright/validate)

#### 参考资料
* [为什么淘宝和京东会员注册的时候用户名一个汉字算两个字符，有什么考虑吗？](https://www.zhihu.com/question/22295828/answer/82576462)

#### License
* [MIT License](LICENSE)
