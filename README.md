
# go-ticket

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/xavier268/go-ticket/master/LICENSE) [![Build Status](https://travis-ci.org/xavier268/go-ticket.svg?branch=master)](https://travis-ci.org/xavier268/go-ticket)  [![GoDoc](https://godoc.org/github.com/xavier268/go-ticket?status.svg)](https://godoc.org/github.com/xavier268/go-ticket)

Secure and flexible (template and config driven) access control solution for events.

## Use cases

Tickets are created and the ticket url is sent to the user.

Staff is authorized by site admin :

* Site administrator log to the /admin page using the superuser password
* Page provides "one-off" authorization "qr-code cards" that can be sent or printed for the event staff, to activate specific roles
* Staff will scan one of the card to get the role associated with it on their smart phone (cookie life set in configuration)
* Currently, for security reasons, role activation cards can only be used once. Admin can generate as many unique activation cards as needed.

Roles include :

* None (default, typically the event visitor)
* Entry : access control, check for ticket validitry and mark as used for entry
* Exit (optionnal) : mark as exited, allowing further reentry
* Review : control current status of ticket 
* Admin, Super : idem, displaying administrative or debugging information.

Scanning a ticket "qr-code card" has various effects depending on the role associated with the device that scans it :

* Role None : diplay ticket in graphical form, te be printed or shown unpon entry
* Other staff role : process ticket and displays its validity, according to the above roles.

## Configuration parameters

* See the Conf in the conf package
* Configuration is ste using default, then file, then env, then flags

## Conventions for using templates

* Templates should reside in the same directory, as configured in the config file.
* Template names are the file names (including extension).

## implementation considérations

* Key requirement was to support any generic (free) qr-code or barcode readers on both android or IOS smart phones (for instance, the www.scan.me apps are working perfectly)
* Current implementation stores everything in memory, but architecture makes it easy to implement database access
* Default is to use qr-code, but datamatrix is also available from configuration setting.

