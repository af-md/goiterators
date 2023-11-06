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

Why there was a need for this feature?

Go doesn't have a built in iterators

Iterators are features that provide access to an element in a collection

Java version of iterators

C# version of iterators has LINQ: where, when, select and etc

For example you can use the yield() to keep passing the functionality to the next phase of the iteration

Its also a scaling issue too. For example the split function returns a slice, but should instead return an iterator.

It's better to return an iterator instead of returning a large slice if at the end you're going to range over it

Iterators are passed to functions because they give access to the elements. 

Iterators are just a way to loop over a collection and they can be passed to functions or returned back to function

Functions can use them to do an operation and then they can be returned to be accessible after the operation

