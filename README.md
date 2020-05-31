# The package replace vowels with numbers

Decodes and encodes a string in English replacing vowels to numbers and vice verse.
E.g. encodes "hello" to "h2ll4", "h3 th2r2" decodes to "hi there"

## Code Status

[![CircleCI](https://circleci.com/gh/evsyukovmv/encryption.svg?style=svg)](https://circleci.com/gh/evsyukovmv/encryption)

## How to use 

Encode
```
go run cli/main.go hello
```

Decode
```
go run cli/main.go -d h3 th2r2
```
