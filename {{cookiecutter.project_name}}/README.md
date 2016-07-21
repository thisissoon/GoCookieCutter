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

The application has a comprohensive test suite. To run the test suite simply
call `make test`.

The `make test` command can also be customised with the following variable overrides:

* `TEST_VERBOSE`: This will print verbose test output: (`TEST_VERBOSE=1 make test`)
* `TEST_COVERAGE`: This will produce a coverae report (`TEST_COVERAGE=1 make test`)

These can also be combined: `TEST_COVERAGE=1 TEST_VERBOSE=1 make test`
