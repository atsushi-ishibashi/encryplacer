# encryplacer
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
