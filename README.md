# confc

[![Build Status](https://travis-ci.org/researchlab/confc.svg?branch=master)](https://travis-ci.org/researchlab/confc)

`confc` is a lightweight configuration management tool focused on:

* keeping local configuration files up-to-date using data stored in each env config file itself
* reloading applications to pick up new config file changes


## Building

Go 1.12 is required to build confc, which uses the new vendor directory.

```
$ mkdir -p $GOPATH/src/github.com/researchlab
$ git clone https://github.com/researchlab/confc.git $GOPATH/src/github.com/researchlab/confc
$ cd $GOPATH/src/github.com/researchlab/confc
$ make
```

You should now have confc in your `bin/` directory:

```
$ ls bin/
confc
```

## Getting Started

Before we begin be sure to [download and install confc](docs/installation.md).

* [quick start guide](docs/quick-start-guide.md)

## Next steps

Check out the [docs directory](docs) for more docs.

