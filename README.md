# ghstar [![Build Status](https://travis-ci.org/monmaru/ghstar.svg?branch=master)](https://travis-ci.org/monmaru/ghstar)
CLI tool to list Github's starred repository.

## Installation
```
go get github.com/monmaru/ghstar
```

## Usage
```
NAME:
   ghstar - A new cli application

USAGE:
   $ ghstar <GitHub User Name>

VERSION:
   1.0

AUTHOR:
   monmaru

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --lang value, -l value       Filter the repository by the name of the programming language
   --sort value, -s value       Sort. You can specify either created, updated or pushed (default: "created")
   --direction value, -d value  Sorting direction. You can specify either desc or asc. (default: "desc")
   --help, -h                   show help
   --version, -v                print the version
```

## License
Licensed under the [MIT](LICENSE) License.