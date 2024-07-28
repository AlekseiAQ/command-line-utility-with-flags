# Command Line Utility With Flags - Directory Size Calculator

This is a command-line utility written in Go that calculates the size of directories.
It supports `-recursive` size calculation and can display sizes in a human-readable format (`-human`).

## Requirements

- Go 1.22.5 or later installed

## Usage

### Build

```bash
go build
```

### Run

```bash
./command-line-utility-with-flags [flags] <directory>
```

### Examples

#### Basic Usage

```bash
./command-line-utility-with-flags /path/to/directory
```

#### Recursive Size Calculation

```bash
./command-line-utility-with-flags -recursive /path/to/directory
```

#### Human Readable Output

```bash
./command-line-utility-with-flags -human /path/to/directory
```

#### Combining Flags

```bash
./command-line-utility-with-flags -recursive -human /path/to/directory
```

#### Help

```bash
./command-line-utility-with-flags -help
```
