# ss

[![Build Status](https://travis-ci.org/ricallinson/ss.svg)](http://travis-ci.org/ricallinson/ss)

Command line tool for performing __simple searches__ over files and directories.

    ss -c /dir 'foo NOT bar'

## Install

You must have [go](http://golang.org/) installed to build __ss__.

    go get github.com/ricallinson/ss

## Usage

    ss
    Usage: [options] filename|directory "query"
      ss -c /dir "a b"
      ss -c /dir "a b NOT y z"
      ss -c /dir "a OR b"

    ss -h
    Usage of ss:
      -c=false: output the count of the number of matches found
      -o=0: output the matches found, use -1 for all
      -t=false: output the time taken in seconds
      -version=false: output the version information

    ss -c /dir 'PUT'
    5

## Query Syntax

Exact positive matching of __'a' AND 'b'__.

    'a b'

Exact negative matching __NOT 'a' AND NOT 'b'__.

    'NOT a b'

Exact positive and negative matching __'a' AND 'b' NOT 'c' AND NOT 'd'__.

    'a b NOT c d'

Exact positive OR matching __'a' OR 'b'__.

    'a OR b'

Exact positive OR group matching __'a' OR 'b' AND 'c' OR 'd' OR 'e'__.

    'a OR b c OR d OR e'

Exact negative OR matching __NOT 'a' OR NOT 'b'__.

    'NOT a OR b'

Exact positive and negative OR matching __'a' OR 'b' NOT 'c' OR 'd' OR 'e'__.

    'a OR b NOT c OR d OR e'

Complete example of positive and negative using OR matching __'a' AND 'b' AND 'c' OR 'd' NOT 'e' AND NOT 'f' AND NOT 'g' OR NOT 'h' OR NOT 'i'__.

    'a b c OR d NOT e f g OR h OR i'

## Why?

You can use `cat`, `grep` and `wc` to achieve the same result without needing to install yet another program.

    cat /dir/**/* | grep PUT | wc -l
           5

But if like me regex is not your first language then simple queries like the following become a time sink.

    ss -c /dir 'home NOT 302'
    5

Performance can also become issue when using `cat`, `grep`, `wc` over several gigabytes worth of files.

## Ideas

### Count items over some time period

How do you specify the field to represent time?

    ss -tp field,1d /dir 'query'
    20, 30, 50, 30, 20

## Test

Install the coverage tool `go get code.google.com/p/go.tools/cmd/cover`.

    go test

## Generate Code Coverage

Get a quick coverage number.

    go test -cover

View the lines covered by the tests.

    go test -coverprofile=coverage.out; go tool cover -html=coverage.out
