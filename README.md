# bissucket

bissucket is a tool to manipulate Bitbucket Issue from the CLI.

## Usage

```
NAME:
   bissucket - bissucket is a tool to manipulate Bitbucket Issue from the CLI.
    First from [bissucket sync] please.

USAGE:
   bissucket [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
     repository, repo  Repository related operations. Currently only list view.
     issue, i          Display the issue of a specific repository.
     sync              Get your repository from Bitbucket.
     help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Install

If you have installed Golang:

```
go get -u bitbucket.org/Masami_Nakaoka/bissucket
```

Or download from [here](https://bitbucket.org/Masami_Nakaoka/bissucket/downloads/).

## After installation

When you execute the command for the first time, you need to enter the user name and password of Bitbucket.

Next, execute the following command to obtain a list of repositories.

```
bissucket sync
```

The list of repositories is stored in the following location:

```
$HOME/.bissucket.repositoriescache.json
```

---

(From here down is under construction)
