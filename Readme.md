# Testing Google cloud platform framework

Running using docker:

```bash
docker run -it -v $(pwd):/go/src -p 8080:8080 golang:latest bash
```

Running:

```bash
cd /go/src && go run cmd/main.go
```

Checking if it works:

```bash
wget http://localhost:8080/
```

Using .gcloudignore file not to push some files to the cloud while publishing. Won't use pack.

## Publishing

```shell
docker run --rm --volumes-from gcloud-config -v $(pwd):/src gcr.io/google.com/cloudsdktool/cloud-sdk:latest gcloud functions deploy HelloWorld --runtime go113 --trigger-http --allow-unauthenticated --source /src/.
```

Check the URL with:

```shell
docker run --rm --volumes-from gcloud-config gcr.io/google.com/cloudsdktool/cloud-sdk:latest gcloud functions describe CpfFull
```

## Google Example

https://github.com/GoogleCloudPlatform/functions-framework-go
