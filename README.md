<h1><img align="center" src="https://github.com/dockit-dev/cli/assets/26047023/7ed3a6a6-c09c-408c-89f5-a8733e7ad0b8" width="70"> dockit</h1>

[![Release](https://github.com/vektra/mockery/actions/workflows/release.yml/badge.svg)](https://github.com/dockit-dev/cli/actions)

The command-line tool provides a convenient way to set up access to remote Docker servers hosted by [Dockit](https://dockit.dev).

## Installation

To install dockit CLI, you can use the following command:

```bash
go get -u github.com/dockit-dev/cli
```

Make sure your Go bin directory is added to your system's PATH.

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
