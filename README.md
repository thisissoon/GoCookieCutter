# SOON\_ Go Cookie Cutter

A Go project cookie cutter that gives us our standard Go project setup.

## Usage

1. Create your `GOPATH` directory and `src` directory: `mkdir -p ~/MyGoProject/src`
2. `cd` into the `src` directory you jsut created: `cd ~/MyGoProject/src`
3. Now use `cookiecutter`: `cookiecutter https://github.com/thisissoon/GoCookieCutter.git`
4. Fill in the information `cookiecutter` asks you
5. The project will now exist in the `src` directory, `cd` into it
6. Install `dep`: https://github.com/golang/dep/releases
7. Install dependencies (these live in the `vendor/` directory: `dep ensure`
8. Read the `README.md`
