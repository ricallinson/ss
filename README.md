# ss

Command line tool for performing __simple searches__ over files and directories.

## Usage

    ss
    Usage: [options] path 'query'
      rsa -c ./path 'a b'
      rsa -c ./path 'a b NOT y z'
      rsa -c ./path 'a OR b'

    ss -h
    Usage of ss:
      -c=false: output the count of the number of matches found
      -o=0: output the matches found, use -1 for all
      -t=false: output the time taken in seconds
      -version=false: output the version information

    ss -c ./fixtures 'PUT'
    5

# Why?

You can use `cat`, `grep` and `wc` to achieve the same result without needing to install yet another program.

    cat ./fixtures/**/* | grep PUT | wc -l
    5

But if like me regex is not your first language then simple queries like the following become a time sink.

    ss -c ./fixtures 'home NOT 302'
    5

Performance can also become issue when using `cat` and `grep` over several gigabytes worth of files.

# Ideas

## Count items over some time period

How do you specify the field to represent time?

    ss -tp field,1d ./path 'query'
    20, 30, 50, 30, 20

# Test

Install the coverage tool `go get code.google.com/p/go.tools/cmd/cover`.

    go test

# Code Coverage

Get a quick coverage number.

    go test -cover

Viewing the lines covered.

    go test -coverprofile=coverage.out; go tool cover -html=coverage.out
