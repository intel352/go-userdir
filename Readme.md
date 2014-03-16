# go-userdir [![Build Status](https://travis-ci.org/wilmoore/go-userdir.png?branch=master)](https://travis-ci.org/wilmoore/go-userdir)

  Platform agnostic access to common user file/directory paths for [Go].

## Home Directory

    userdir.Home()
    //=> /Users/wilmoore

## Authorized Keys File

    userdir.AuthorizedKeys()
    //=> /Users/wilmoore/.ssh/authorized_keys

## Installation

    % go get github.com/wilmoore/go-userdir

## Import

    import "github.com/wilmoore/go-userdir"

## License

  MIT

[Go]:  http://golang.org/

