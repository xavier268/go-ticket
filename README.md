
# go-ticket

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/xavier268/go-ticket/master/LICENSE) [![Build Status](https://travis-ci.org/xavier268/go-ticket.svg?branch=master)](https://travis-ci.org/xavier268/go-ticket)  [![GoDoc](https://godoc.org/github.com/xavier268/go-ticket?status.svg)](https://godoc.org/github.com/xavier268/go-ticket)

Secure and flexible (template and config driven) access control solution for events.

**Work in progress ...**

## Conventions for using templates

* Templates should reside in the same directory, as configured in the config file.
* Template names are the file names (including extension).
* Templates are expecting a Session object (which incorporates a pointer to the Conf object) or a Conf object.
* Nested templates are passed the full data set.
