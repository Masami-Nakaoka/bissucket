# bissucket

bissucket is a tool to manipulate Bitbucket Issue from the CLI.

## Usage

```shell
NAME:
   bissucket - bissucket is a tool to manipulate Bitbucket Issue from the CLI.
    First from [bissucket sync] please.

USAGE:
   bissucket [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
   sync     Synchronize with Bitbucket\'s repository and issue.
   list     Display Issue and list of repositories. Display a list of defaultRepository if no options are given.
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
go get -u github.com/Masami-Nakaoka/bissucket
```

## After installation

When you execute the command for the first time,  
you need to enter the user name and password of Bitbucket.

Next, execute the following command to obtain a list of repositories.

```shell
bissucket sync --repository
```

The list of repositories is stored in the following location:

```shell
$HOME/.bissucket.repositoriescache.json
```

## Todo

I will investigate whether it can be realized.

- Command add
    - issue create
    - issue complete
    - add comment
    - etc...

## Author

Masami-Nakaoka

## License

[MIT](https://opensource.org/licenses/MIT)
