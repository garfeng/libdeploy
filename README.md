# LibDeploy

# Desc

Copy all dependencies to the dir where a program built by gcc/g++ stays.

So that you can package them and run them on another computer that may not have the dependencies installed.

# Run

``` shell
libdeploy -no=Windows -no=system32 helloworld.exe
libdeploy -no=Windows testlib.dll
```

# Build

```shell
go get github.com/garfeng/libdeploy
go install github.com/garfeng/libdeploy
```
