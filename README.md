# Mp4-CLI

CLI Tool for video mp4

## Features

- List videos
- Get info file
- Split video

## Requirements

- [Go](https://golang.org/doc/install)
- [FFmpeg](https://ffmpeg.org/download.html)

## Installation

### Installing Go

Provide instructions or a link to the Go installation guide: [Go Installation Guide](https://golang.org/doc/install).

### Installing FFmpeg

#### Windows

1. Open Terminal with Run as Administrator and run `choco install ffmpeg`

#### macOS

1. Open Terminal and run `brew install ffmpeg`.

#### Linux

1. Open Terminal and run `sudo apt-get install ffmpeg` (for Ubuntu/Debian) or the appropriate command for your distro.

## Usage

```
usage: mp4-cli [options]
```

## Examples

`mp4-cli split xxx.mp4 00:00:10 00:01:00`
