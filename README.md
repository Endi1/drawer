# Drawer

## The simplest drawer for your bookmarks

### Table of contents
1. [Installation](#installation)
2. [Quickstart](#quickstart)
3. [Adding bookmarks from the CLI](#adding-bookmarks-from-the-cli)
4. [Finding and displaying bookmarks](#finding-and-displaying-bookmarks)  
   - [By ID](#by-id)
   - [By tag](#by-tag)


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

From there follow the instructions to add your first bookmark to
drawer. The URL of your new bookmark _needs_ to start with the
protocol (either http:// or https://).

To view all your bookmarks, run:

`drawer`

again in the same directory where the bookmarks file is (by default it
must be called ".drawer", otherwise you need to use the -f flag to
specify the name of the file).


# Adding bookmarks from the CLI

To add a new bookmark run:

`drawer -a`

in the same directory where your bookmarks file is. 

You will then be prompted with these steps:

1. Add a new bookmark (Y/n)?

   Pretty self-explanatory. "n" for no and "y" or empty for yes.
   
2. Add a link:

   The bookmark URL. **Needs to start with the protocol (https:// or http://)**
   
3. Add a title for that link:

   The title for that bookmark. Can be empty
   
4. Add a comment: 

   In case you want to add a comment for that bookmark
   you can do so here. Otherwise leave empty.
   
5. Add a tag (empty to stop): 

   In case you want to add tags to separate
   bookmarks in different categories so that they can be searched
   easier, you can do so here. As long as you add a tag it will keep
   asking you for one (you can add more than one tag). If you are done
   adding tags, just press "Enter" without writing anything.

### Example

```
$ drawer -a
file:.mydrawer
> Add a new bookmark (Y/n)? y
> Add a link: https://google.com
> Add a title for that link: Google
> Add a comment: This is a search engine
> Add a tag (empty to stop): search engine
> Add a tag (empty to stop):
```

Now typing `drawer` in the terminal will yield:

```
file:.mydrawer
0. Google: https://google.com
// This is a search engine
#search engine
```

# Finding and displaying bookmarks

To display all your saved bookmarks you can run just `drawer` and it
will list all the bookmarks in the drawer file.

## By id

If you want to display only one bookmark of which you know the ID. You
can use the `-i` flag. From the previous example:

```
$ drawer -i 0

file:.mydrawer
0. Google: https://google.com
// This is a search engine
#search engine
```

## By tag

To display bookmarks that have a certain tag, you can use the `-t`
flag. From the previous example:

```
$ drawer -t search engine

file:.mydrawer
0. Google: https://google.com
// This is a search engine
#search engine
```

