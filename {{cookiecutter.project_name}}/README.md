# {{ cookiecutter.client_name }} {{ cookiecutter.project_human_name }}

{{ cookiecutter.description }}

## Development

This project is written in `golang`. Please read up on `golang` here: https://golang.org/doc/
This project requires `go 1.5` and above.

### Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Create your directory paths: `mkdir -p ~/{{ cookiecutter.project_name|lower }}/{src/{{ cookiecutter.project_name|lower }},pkg,bin}`
2. Set your `GOPATH`: `export GOPATH=~/{{ cookiecutter.project_name|lower }}`
3. Change directory to the `src` directory: `cd ~/{{ cookiecutter.project_name|lower }}/src/{{ cookiecutter.project_name|lower }}`
4. Clone the repository: `git clone {{ cookiecutter.git_clone_url }} .`
5. Install `govendor`: https://github.com/kardianos/govendor
6. Install dependencies (these live in the `vendor/` directory: `govendor sync`
8. Build: `go build`
9. 💥

### Testing

The application has a comprohensive test suite covering unit and integration tests.

#### Test Enviornment Variables

The test suite supports the following environment variables:

* `VERBOSE`: This will print verbose test output: (`VERBOSE=1 make test`)

#### Coverage Reports

Test coverage reports are generated in the `.cover` directory. Here you will find
`cover.out` containing the raw line count coverage report and `cover.html` containing
a `html` generated report which can be viewed in your browser. On OSX you can use the
`open` command:

```
open .cover/cover.html
```

You will also find each individual packages line count coverage file.

## Configuration

By default the application will look for configuration files in the following locations:

* `/etc/{{ cookiecutter.client_name|lower }}/{{ cookiecutter.project_name|lower }}/config.toml`
* `$HOME/.config/{{ cookiecutter.client_name|lower }}/{{ cookiecutter.project_name|lower }}/config.toml`

These configuration files **MUST** be in `toml` format.

Please read more about `toml` here: https://github.com/toml-lang/toml

### Specific Configuration File

You can also provide a specific absolute path to a configuration file via
the CLI `-c` or `--config` flag. For example:

```
{{ cookiecutter.project_name|lower }} -c /path/to/config.toml
```

Please note that if the config file is not found or malformed the application
will fail to start.
