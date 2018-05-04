# encryplacer
[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build Status][5]][6]
[![codecov][7]][8]

[1]: https://godoc.org/github.com/atsushi-ishibashi/encryplacer?status.svg
[2]: https://godoc.org/github.com/atsushi-ishibashi/encryplacer
[3]: https://goreportcard.com/badge/github.com/atsushi-ishibashi/encryplacer
[4]: https://goreportcard.com/report/github.com/atsushi-ishibashi/encryplacer
[5]: https://travis-ci.org/atsushi-ishibashi/encryplacer.svg?branch=master
[6]: https://travis-ci.org/atsushi-ishibashi/encryplacer
[7]: https://codecov.io/gh/atsushi-ishibashi/encryplacer/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/atsushi-ishibashi/encryplacer

encryplacer is CLI to replace the encryption of S3 object with KMS encryption

## Installing
```
make build
```

## Get Started
```
$ encryplacer -h
Usage of encryplacer:
  -bucket string
    	bucket name
  -concurrent int
    	concurrent (default 3)
  -contain string
    	contain to filter objects
  -kms string
    	KMS Key ID for encryption
  -region string
    	AWS region. priority: arg>AWS_REGION>AWS_DEFAULT_REGION
  -suffix string
    	suffix to filter objects
```
```
$ encryplacer -bucket <bucket> -kms <key-id>
key1 success.
key2 success.
...
```

### Warning
Currently `encryplacer` upload object with the same name without backing up.

### TODO
- ReadCloser to ReadSeeker
- prefix arg
- use channel to pass key to `replaceEncryption`
