
# webm2mp4

Simple server to convert WebM video to MP4, as a small working sample for [goavcodec](https://github.com/otiai10/goavcodec).

# Development

```sh
% docker-compose up
% open http://localhost:8080
```

# Deploy

```sh
% heroku create
% heroku container:login # If needed
% heroku container:push web
% heroku open
```
