# {{ cookiecutter.name }}

{{ cookiecutter.description }}

## Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Create your directory paths: `mkdir -p ~/{{cookiecutter.name}}/src/{{cookiecutter.name}}`
2. Set your `GOPATH`: `export GOPATH=~/{{cookiecutter.name}}`
3. Change directory to the `src` directory: `cd ~/{{cookiecutter.name}}/src/{{cookiecutter.name}}`
4. Clone the repository: `git clone {{cookiecutter.origin}} .`
5. Install `dep`: https://github.com/golang/dep/releases
6. Install dependencies (these live in the `vendor/` directory: `dep ensure`
8. Build: `make build`
9. 🍻
