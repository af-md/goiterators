# GO ITERATORS

This repo contains an example implementation of iterators using the newly available feature "range over func"

To track this feature: <https://github.com/golang/go/issues/61405>

How to install this experimental feature:
```
go install golang.org/dl/gotip@latest
gotip download 510541   # download and build CL 510541
gotip version  # should say "(w/ rangefunc)" or just output a hash
```

How to run the code with the feature:
```
GOEXPERIMENT=range gotip test main_test.go  -bench .
```
