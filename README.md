# goflow-builder

A Go-based tool to generate custom GitHub Actions workflows.

## About

This repository contains a Go tool  designed to build and generate custom GitHub Actions workflows. The tool allows you to automate the creation of workflow files tailored to your project’s testing, deployment, or other CI/CD needs.

## Prerequisites

- Go (version 1.22.2 or later) installed on your system
- Git installed to manage the repository
- Access to a terminal or command-line interface

## Installation

### Download Using `go install`
Install the latest zip manually or via the following github client script (Recommended):
```bash
§ gh release download -R ronthesoul/goflow-builder --pattern 'goflow-builder-{Insert here the wanted version}.zip'
unzip goflow-builder-{Insert here the wanted version}.zip
```

Install the `goflow-builder` tool directly from the source using Go:

```bash
go install github.com/ronthesoul/goflow-builder@v1.0.49
```

This command downloads and installs version `v1.0.49` of the tool, making the `goflow-builder` binary available in your `$GOPATH/bin` directory (ensure `$GOPATH/bin` is in your `PATH`).

### Alternative: Clone and Build
Clone the repository and build the tool manually:

```bash
git clone https://github.com/ronthesoul/goflow-builder.git
cd goflow-builder
go build -o goflow-builder main.go
```

## Usage

### Generate a Workflow
Use the `generate` command to create a custom workflow file:

```bash
goflow-builder generate [flags]
```

The generated workflow is saved as `workflow.yml` by default. Customize the output using the available flags.

### Available Flags for `generate`
The following table lists all supported flags for the `generate` command:

| Flag          | Description                                      | Default Value    | Example Usage             |
|---------------|--------------------------------------------------|------------------|---------------------------|
| `-d`, `--dry-run` | Output YAML to console instead of writing to file | `false`          | `-d`                     |
| `-f`, `--file` | Output file path for the generated workflow      | `workflow.yml`   | `-f custom-workflow.yml`  |
| `-h`, `--help` | Display help information for the command         | N/A              | `-h`                     |
| `-n`, `--name` | Name of the workflow                             | `ga-workflow`    | `-n my-workflow`          |
| `-y`, `--notify` | Enable Slack notifications                       | `false`          | `-y`                     |
| `-r`, `--runner` | Choose runner OS                                 | `ubuntu-latest`  | `-r windows-latest`       |
| `-s`, `--steps` | Number of steps to generate                      | `1`              | `-s 3`                   |

- Example with multiple flags:
  ```bash
  goflow-builder generate -d -f my-workflow.yml -n test-workflow -r ubuntu-22.04 -s 2 -y
  ```

### Customize the Workflow
- Modify the `main.go` source code to adjust the workflow generation logic (e.g., add specific steps or triggers).
- Rebuild the tool after changes:
  ```bash
  go build -o goflow-builder main.go
  ```
## Contributing
Fork the repository, make changes, and submit a pull request. Ensure the Go script builds without errors.

## 👤 Author

Written by Ron Negrov
