# gobazel

Example of using Go and Bazel

# Requirements

Install `bazelisk` (see [here](https://github.com/bazelbuild/bazelisk)). I think that easiest way is to execute:

```shell
go get github.com/bazelbuild/bazelisk
```


# Run server

```shell
bazelisk run server/cmd
```

# Fire requests

```shell
# Good
curl --request POST 'http://localhost:8080/compute' --header 'Content-Type: application/json' --data-raw '{"Op": "add","Lh": 1,"Rh": 5}'

# Bad
curl --request POST 'http://localhost:8080/compute' --header 'Content-Type: application/json' --data-raw '{"Op": "mul","Lh": 1,"Rh": 5}'
```