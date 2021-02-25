# file-mod

File mod is a command line go utility to modify files. Below you can find the the availability and usage.

### Installation

You can install the file-mod component for a specific OS at https://github.com/chef/file-mod/releases

Currently available for:
OS | ARCH
--- | ---
Darwin | x86_64 (amd64), arm64
Linux | x86_64 (amd64), arm64
Windows | x86-64 (amd64)

### Build

You can optionally build file-mod via the [goreleaser](https://goreleaser.com/) utility. Goreleaser uses a `.goreleaser.yml` to define how you would like to build and release your go binary(ies).

### `file-mod`

<!-- stdout "./build/linux/file-mod --help" -->
```
Command line utility to modify files.

Usage:
  file-mod [command]

Available Commands:
  append-if-missing Append STRING to FILE if not already there.
  find-and-replace  Replace REGEX_STR with STRING in FILE. Supports multiline replace.
  help              Help about any command

Flags:
  -h, --help   help for file-mod

Use "file-mod [command] --help" for more information about a command.
```
<!-- stdout -->
