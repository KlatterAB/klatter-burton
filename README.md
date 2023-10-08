## Installation

> git clone https://github.com/KlatterAB/klatter-burton <br>
> cd butler-burton <br>
> make install

or download the binary and run it

## Configuration

Edit config-file in `$HOME/.config/klatter-burton`

Example-config:

```yaml
name: Klatter Burton
color: "#46D9FF"
notifications: true
id: kb
```

## Development

### Build docker image

```sh
docker build -t "imageName" .
```

### Start docker container

```sh
docker run -it "imageName" sh
```

### Get time sheet

```sh
docker ps (get containerId)
docker cp "containerId":/root/.butlerburton/"reportName.xlsx" .
```

### Generate man page from butlerburton.md

```sh
pandoc butlerburton.md -s -t man -o butler-burton.1
gzip butler-burton.1
```
