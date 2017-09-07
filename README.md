
# webm2mp4

Simple server to convert WebM video to MP4, as a small working sample for [goavcodec](https://github.com/otiai10/goavcodec).

Try now here https://webm2mp4.herokuapp.com/, and deploy your own now.

# Qick Start

## Local Development

with [docker](https://www.docker.com/products/docker-toolbox) and [docker-compose](https://www.docker.com/products/docker-toolbox) required

```sh
% docker-compose up
# open http://localhost:8080
```

## Deploy to Heroku

with also [heroku cli](https://devcenter.heroku.com/articles/heroku-cli#download-and-install) required

```sh
% heroku create
% heroku container:login # If needed
% heroku container:push web
# heroku open
```
