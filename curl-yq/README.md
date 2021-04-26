# Curl-yq

Small docker image based off the [curl official image](https://hub.docker.com/r/curlimages/curl) that includes yq to enable json/yml parsing.

Mainly used as a base image to run other scripts

## Usage

```bash
docker pull nfowl/curl-yq
docker run -it nfowl/curl-yq /bin/sh
```