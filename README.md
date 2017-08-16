
# webm2mp4

Simple server to convert WebM video to MP4, as a small working sample for [goavcodec](https://github.com/otiai10/goavcodec).

# Deploy

## Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Docker

`docker run --rm -p 8080:8080 otiai10/webm2mp4`

## manual deploy

```sh
sudo apt-get install -y ffmpeg
go get github.com/otiai10/webm2mp4
$GOPATH/bin/webm2mp4
```

# API Endpoints

// TODO: Write something
