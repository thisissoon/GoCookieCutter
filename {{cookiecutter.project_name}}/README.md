# {{ cookiecutter.project_human_name }}

{{ cookiecutter.description }}

## Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Create your directory paths: `mkdir -p ~/{{cookiecutter.project_name}}/{src/{{cookiecutter.project_name}},pkg,bin}`
2. Set your `GOPATH`: `export GOPATH=~/{{cookiecutter.project_name}}`
3. Change directory to the `src` directory: `cd ~/{{cookiecutter.project_name}}/src/{{cookiecutter.project_name}}`
4. Clone the repository: `git clone {REPO_PATH} .`
5. Install `govendor`: https://github.com/kardianos/govendor
6. Install dependencies (these live in the `vendor/` directory: `govendor sync`
8. Build: `go build`
9. 🍻

## Testing

Run the test suite by calling the `make test` command. A coverage profile can be generated and
viewed in your browser by running `make coverage`.
