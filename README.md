# aws-kms

```
Usage: aws-kms [global options] <verb> [verb options]


Global options:
        -h, --help   Show this help

Verbs:
    decrypt:
        -i, --input  Input data (*)
    encrypt:
        -k, --key-id Key ID to encrypt with (*)
        -i, --input  Input data (*)
```

## Examples

*Encrypting data, e.g:*
```bash
aws-kms encrypt --input "some secret infos" --key-id alias/MyMasterKey
aws-kms encrypt -i "some secret infos" -k 12345-123456-1234567
 ```
 
*Decrypting data, e.g:*

```bash
aws-kms decrypt -i "AQECAHgBe/Bo6N6P3DZu2j8rSbOworqekgIES5evTHicjBhp6A=="
```

## Building
```bash
glide install --strip-vendor
go build
```
