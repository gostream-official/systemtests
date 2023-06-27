# systemtests

---

*systemtests* is part of the *gostream* project. *gostream* is simple music database. *systemtests* provides executable system tests for *gostream* written in *Go*.

Features:

- quickly test the *gostream* system
- automate system tests for *gostream*

---

## Usage

To execute any system test, use *Go*:

```sh
$ go run cmd/tracks/create/run.go \
  --server-url http://localhost:9999 \
  --mongo-uri mongodb://root:example@127.0.0.1:27017
```
