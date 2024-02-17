<h1><img align="center" src="https://github.com/dockit-dev/cli/assets/26047023/7ed3a6a6-c09c-408c-89f5-a8733e7ad0b8" width="70"> dockit</h1>

[![Test](https://github.com/dockit-dev/cli/actions/workflows/makefile.yml/badge.svg)](https://github.com/dockit-dev/cli/actions)
[![codecov](https://codecov.io/gh/dockit-dev/cli/graph/badge.svg?token=IAQXVDRKDL)](https://codecov.io/gh/dockit-dev/cli)

The command-line tool provides a convenient way to set up access to remote Docker servers hosted by [Dockit](https://dockit.dev).

## Installation

### Via Homebrew

You can install the CLI using Homebrew. First, tap into this repository:

```bash
brew tap dockit-dev/dockit
```

Then, install the CLI tool:

```bash
brew install dockit
```

### Via Go

If you have Go installed, you can also install the CLI tool directly using `go get`:

```bash
go install github.com/dockit-dev/dockit@latest
```

Make sure your Go bin directory is added to your system's PATH.

### Releases

You can download pre-built binaries from the [releases](https://github.com/dockit-dev/dockit/releases) page on GitHub. Choose the appropriate binary for your system and architecture, download it, and place it in a directory included in your system's PATH.

### Verifying Installation

To verify the installation, run:

```bash
dockit --version
```

## Command Reference

### configure
The configure command sets up access to a remote Docker server by providing the Dockit configuration file. Additionally, it creates a new Docker context and sets it as active, enabling seamless interaction with the remote Docker server.

```bash
dockit configure [path]
```

`[path]`: path to the Dockit configuration file (required).

<b>Example</b>

```bash
dockit configure /path/to/dockit_configuration.tar.gz
```

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the Apache License - see the [LICENSE](https://github.com/dockit-dev/cli/blob/master/LICENSE) file for details.
