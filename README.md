# pastenull
(extremely) minimal pastebin service written in go. Runs in a docker container and outputs gopher:// and http:// links
Accepts multi-part form data with `curl -F` from file or stdin.

## server usage (docker)
1. Build the image
```
git clone https://github.com/jsav0/pastenull
cd pastenull
docker build . -t pastenull
```

With docker, I mount a directory already being served by my web server on my domain. ie. `-v /srv/www/darkhttpd/paste:/srv/paste`

2. Run it
```
docker run -it -p 1337:1337 -v /path:/srv/paste pastenull
```

It's also hosted on my docker hub so you can pull and run without building locally first:
```
docker run -it -p 1337:1337 -v /path:/srv/paste wfnintr/pastenull
```

Note: Right now i have my test domain hardcoded in main.go. If you run it locally, the outputted links will be wrong. Use localhost instead for testing.

---

## client usage (upload)
1. upload file
```
curl -F'file=@file' wfnintr.net:1337

```

2. upload data from stdin
```
curl -F'file=@-' wfnintr.net:1337 < file

```

