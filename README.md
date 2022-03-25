# analyzeinator

The analyzeinator^TM is designed to consume JSON trade data via stdin, continously compute, and return results from the data set.


## Design considerations

This is the first Go app I've ever written, so a lot of discovery took place during the design process. I needed to ensure I didn't grab all the cool shiny tech off the shelves and try to cram it into this app. I wanted to keep it as simple as possible while focusing on performance. 



## Challenges

Hopefully I look back on these challenges in a few months and laugh at how trivial they were. In the moment I spinned my wheels for a while on the following issues:

__Dealing with different types of data on stdin__

When I started this project I focused all of my attention on how I was going to read and write data. I discovered that it's unsuprisingly much faster to decode JSON rather than unmarshal, or read bytes from stdin. I tried to take the approach of using a reader to look for BEGIN, swapping to json decode for the primary payload and attempting to switch back to read bytes at the end of the json payload. What I should probably do is just ReadBytes and call it a day. Can't let go of those crucial seconds of execution time.

__How to somehow use go routines__

After reading about what can be accomplished with go and messing around with setting up channels, I was hooked on the idea of making this application multi-threaded. In the end I ran into issues exceeding tresholds for memory pointers (or something to that effect). KEEP IT SIMPLE!

__Finding a good data structure for market data__

It took me a while to land on the data structure design that I've implemented, and it's probably horrible. I'm excited to review this app and determine the best way to organize this data.
