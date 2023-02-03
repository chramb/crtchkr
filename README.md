
# crtchkr - Projekt
[![](https://img.shields.io/badge/license-BEER--WARE-critical)](https://github.com/chramb/crtchkr/blob/master/LICENSE)
![](https://img.shields.io/github/go-mod/go-version/chramb/crtchkr)
![](https://img.shields.io/badge/platform-Linux-yellow)

> Proste narzędzie monitorujące ważność certyfikatów - lokalnych i zdalnych z systemem powiadomień


## Dokumentacja 
Instrukcje korzystania poza Dokumentacją można znaleźć w [man/](/man)
- [`crtchkr(1)`](/man/crtchkr.1.md)
- [`crtchkr-check(1)`](/man/crtchkr-check.1.md)
- [`crtchkr-message(1)`](/man/crtchkr-message.1.md)
- [`crtchkr-print(1)`](/man/crtchkr-print.1.md)
- [`crtchkr.toml(5)`](/man/crtchkr.toml.5.md)

## Kompilacja programu

Wymagany jest kompilator języka Golang

znajdując się w głównym katalogu projektu należ wykonać komendę:
### Make
```shell
$ make
```
lub
### Ręczna Kompilacja
```shell
$ go build
```

## Generowanie Dokumentacji

### Man
Wymagany jest program [go-md2man](https://github.com/cpuguy83/go-md2man)
```
$ make man
```
### PDF
Wymagany jest język R wraz z: pandoc, Rmarkdown, oraz Xelatex
```
$ make pdf
```
