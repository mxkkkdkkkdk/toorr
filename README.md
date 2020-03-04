<div align="center">
  <br>
  <img alt="TorCloud" src="https://i.imgur.com/fPbQU5K.png" width="200px">
  <h1>☁️ TorCloud ☁️</h1>
  <strong>Torrent on Cloud</strong>
</div>
<br>
<div align="center">
  <a href="https://github.com/yohix/torcloud/actions">
    <img src="https://github.com/yohix/torcloud/workflows/CI/badge.svg" alt="CI">
  </a>
  <a href="https://golang.org">
    <img src="https://img.shields.io/badge/Go-v1.12.7-green.svg" alt="Go version">
  </a>
</div>

## Pitch

**TorCloud** is a cloud & self-hosted torrent downloader. On providing the magnet link files are downloaded from the distributed torrent and stored in the server, which can be streamed or downloaded via HTTP/HTTPS.

## Install

##### Source

*[Go](https://golang.org/dl/) is required to install from source*

``` sh
go get -v github.com/yohix/torcloud
```

## Setting up

```sh
sudo apt update
sudo apt install build-essential
sudo apt install golang-go
export GOPATH=$HOME
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
mkdir -p $GOPATH/src/github.com/yohix
cd src/github.com/yohix
git clone https://github.com/yohix/torcloud.git
```

##### Running TorCloud as Service

```sh
cp torcloud/scripts/service/torcloud.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable torcloud.service
sudo systemctl restart torcloud.service
sudo systemctl status torcloud.service
```

## License

MIT © [Yohix](https://yohix.github.io/mit)
