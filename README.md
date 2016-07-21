# SOON\_ Go Cookie Cutter

A Go project cookie cutter that gives us our standard Go project setup.

## Features

* Built in CLI framework foundation using `cobra`
* Built in `config` package foundation using `viper`
* Easy to configure centralised `logger` package based on `logrus`
* Included `Dockerfile`
* Ready to go `Makefile`
* Dependencies managed by `govendor`
* Test coverage reporting
* Git Lab CI ready
* `89%` Test Coverage

## Usage

1. Create your `GOPATH` directory and `src` directory: `mkdir -p ~/MyGoProject/src`
2. `cd` into the `src` directory you jsut created: `cd ~/MyGoProject/src`
3. Now use `cookiecutter`: `cookiecutter https://github.com/thisissoon/GoCookieCutter.git`
4. Fill in the information `cookiecutter` asks you
5. The project will now exist in the `src` directory, `cd` into it
6. Follow the `README.md` in project to complete the setup
