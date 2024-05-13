# viduration

## Description

viduration is a command-line tool to recursively walk a directory and print out the
video-length of each video file and total length of all files.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Credits](#credits)
- [License](#license)

## Installation

```shell
$ go install github.com/mehdieidi/viduration@latest
```

## Usage

```shell
$ viduration -d ./movies
```

Use -e flag to exclude a directory from getting searched.

## Credits

[go-ffprobe](https://github.com/vansante/go-ffprobe)

## License

MIT
