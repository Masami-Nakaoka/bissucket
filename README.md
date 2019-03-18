# bissucket

bissucket is a tool to manipulate Bitbucket Issue from the CLI.

## Usage

```shell
NAME:
   bissucket - bissucket is a tool to manipulate Bitbucket Issue from the CLI.

USAGE:
   bissucket [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
   list     Display issues for the specified repository.
   config   Command to set bissucket related operations. If there is no argument, display a list of settings.
   show     Display Issue details of defaultRepository.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Install

If you have installed Golang:

```shell
go get -u github.com/namahu/bissucket
cd $GOPATH/src/github.com/namahu/bissucket
go install
```

## After installation

When you execute the command for the first time,  
you need to enter the user name and password of Bitbucket.

## Todo

- Command add
  - issue create
  - issue complete
  - add comment
  - etc...

## Author

namahu

## License

[MIT](https://opensource.org/licenses/MIT)
