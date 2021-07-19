## Efficient Markets: The Space Shuttle *Challenger* Story

### Introduction

I wrote this thesis in the summer of 1992 after I had written a program in SAS to experiment with economic event studies. While experimenting with that program, I tried the day of the *Challenger* accident and discovered that Morton Thiokol was the main recipient of the market's wrath, so to speak. After I investigated a little further, I realized I had a great topic for my thesis (in economics).

### About the repo

This repo contains the updated thesis files that build with the latest version of LaTeX. I added the original Stata and SAS code I wrote to do event studies. It's probably so out-of-date that it should be considered *harmful* or *damaging*, maybe *toxic*.

I even corrected some typos and redrew the graphs after seeing the gnuplot output up close and cringing at ugly pixellated line graphs. So I wrote a little hack in Go to recreate the graph lines as polylines. Those fixed files are in the graphs-source directory, and they're [Affinity Designer](https://affinity.serif.com/en-us/designer/) files. The Go hack is included too.

Unfortunately I don't have any of the original data I used. It's a miracle that I still have this thesis after all of these years. 

I didn't keep track of all the LaTeX dependencies required to build the thesis (sorry, really, really sorry), and it doesn't have proper makefile, just two sad "scripts": cleanit and makeit. I'm embarrassed.

### Other Resources on *Challenger*

One of my thesis advisors, Mike Maloney, wrote a [paper](https://maloney.people.clemson.edu/challenger.pdf) that goes into much more detail about the intraday trading that affected Morton Thiokol's stock and a whole lot more. It's definitely worth a read.

The popular book [*The Wisdom of Crowds*](https://amzn.to/3BiR8MM) by James Surowiecki discusses Maloney's paper in Chapter 1, section II. I hate the idea that this research is being used to support the idea of crowds having some sort of wisdom when it's really individuals acting in their own self interest---there are no crowds here. Ideas like *wisdom of crowds* gives us bad stuff like team rooms, groupthink, and the *Challenger* accident. (By the way, that link to *The Wisdom of Crowds* is an affiliate link.)

And, finally, the Rogers Commission Report is available online [here](https://history.nasa.gov/rogersrep/genindex.htm).