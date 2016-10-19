# Drawer

## The simplest drawer for your bookmarks

### Table of contents
1. [Installation](#Installation)
2. [Quickstart](#quickstart)


# Installation

You need to have [Go](https://golang.org/) already installed on your
computer. If you don't have it, follow
[these](https://golang.org/doc/install) instructions for your
distribution. You can also compile it from source. If you already have
it, move to the next paragraph.

Installing Drawer itself is pretty easy. Simply run: 

`go get -u github.com/endi1/drawer`

After it finishes running, to install run:

`go install github.com/endi1/drawer`


# Quickstart

After installing drawer, you can then run:

`drawer`

in your home directory or wherever you want your bookmarks file to be.
After that, drawer will automatically create a new bookmarks file (by
default it will be ".mydrawer"). After the file is created you will be
asked if you want to add a new bookmark. If the answer is no, just
press "n" and then "Enter". Otherwise press "Enter".

From there follow the instructions to add your first bookmark to drawer.

To view all your bookmarks, run:

`drawer`

again in the same directory where the bookmarks file is (by default it
must be called ".drawer", otherwise you need to use the -f flag to
specify the name of the file).
