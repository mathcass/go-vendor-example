# Golang Vendor Example for Appengine

Here's an example of a simple Google Appengine app written in Go that has vendored dependencies. 

To replicate it, clone down the repository and try to deploy it to appengine using the `gcloud` tool. You'll need to make sure you've installed the `app-engine-go`

``` bash
$ gcloud app deploy

[...output redacted...]

ERROR: (gcloud.app.deploy) Error Response: [9] Deployment contains files that cannot be compiled: Compile failed:
2017/05/11 19:07:24 go-app-builder: Failed parsing input: parser: bad import "unsafe" in vendor/golang.org/x/net/ipv6/bpfopt_linux.go

```


Here's information about Gcloud version: 

``` bash
$ gcloud version
Google Cloud SDK 154.0.1
app-engine-go 
app-engine-python 1.9.52
bq 2.0.24
core 2017.05.04
gcloud 
gsutil 4.25

```
