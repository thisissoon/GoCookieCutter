# {{ cookiecutter.name }}

{{ cookiecutter.description }}


## Development

 - Go 1.11+
 - Dependencies managed with `go mod`

### Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Clone the repository: `git clone {{cookiecutter.origin}} {{cookiecutter.name}}`
2. Build: `make build`
3. 🍻

### Dependencies

Dependencies are managed using `go mod` (introduced in 1.11), their versions
are tracked in `go.mod`.

To add a dependency:
```
go get url/to/origin
```

### Configuration

Configuration can be provided through a toml file, these are loaded
in order from:

{% if cookiecutter.project is not none -%}
- `/etc/{{cookiecutter.project}}/{{ cookiecutter.name }}.toml`
- `$HOME/.config/{{cookiecutter.project}}/{{ cookiecutter.name }}.toml`
{% else -%}
- `/etc/{{cookiecutter.name}}/{{ cookiecutter.name }}.toml`
- `$HOME/.config/{{ cookiecutter.name }}.toml`
{% endif -%}

Alternatively a config file path can be provided through the
-c/--config CLI flag.

#### Example {{ cookiecutter.name }}.toml
```toml
[log]
format = "json"  # [json|console|discard]
```
