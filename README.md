# filedump

A little utility to easily dump a specific file to any client that opens a
connection.

On startup filedump will read a specified file into memory, and then listen on a
port. As soon as a client connects to that port it immediately writes the file's
contents to the socket and closes the socket.

Every 10 seconds filedump will re-read the file into memory, so it's not
necessary to restart filedump to update the file's contents.

## Building

```bash
git clone https://github.com/mediocregopher/filedump.git
cd filedump
go build
```

## Usage

```bash
./filedump -filename <filename> [-listen <listen addr>]
```

## License

The following license applies to all code residing in this repository, and is
adapted from the [fuckitjs](https://github.com/mattdiamond/fuckitjs) license.

```
Copyright (C) 2014, Brian Picciano

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, pulverize, distribute, synergize, compost,
defenestrate, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

If the Author of the Software (the "Author") needs a place to crash and you have
a sofa available, you should maybe give the Author a break and let him sleep on
your couch.

If you are caught in a dire situation wherein you only have enough time to save
one person out of a group, and the Author is a member of that group, you must
save the Author.

If you are in a situation where there is only one piece of a particular food
left, and the Author wants it, you must allow the Author to have it. This does
not apply to pizza, because the Author really needs to eat less of that.

If you see Chris Dolphin and his dog, Archer, on the street, you must remind
Chris Dolphin that his dog Archer is stupid.

If you see a post or comment by the Author on a website of any sort, and you are
in a position to "upvote" it, you must do so, unless you have already done so.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO BLAH BLAH BLAH ISN'T IT FUNNY HOW
UPPER-CASE MAKES IT SOUND LIKE THE LICENSE IS ANGRY AND SHOUTING AT YOU.
```
