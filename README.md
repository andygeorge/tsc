# tsc

Tail Site Code

## Installation / Update

```sh
# install go
sh -c 'VERSION="v0.0.2"; GH="github.com/andygeorge/tsc"; GOPRIVATE=$GH go install -v $GH@$VERSION'
```

## Usage

```
tsc [-f] URL
  -f --follow: follow
```

eg:

```
$ tsc -f http://google.com
2023-08-02 10:34:41 CDT http://google.com: 200 OK
2023-08-02 10:34:41 CDT http://google.com: 200 OK
2023-08-02 10:34:42 CDT http://google.com: 200 OK
2023-08-02 10:34:43 CDT http://google.com: 200 OK
2023-08-02 10:34:43 CDT http://google.com: 200 OK
```
