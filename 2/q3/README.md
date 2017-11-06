Jonathan Guillotte-Blouin — 7900293

To run this program, you need to install external dependencies (as plotting is not part of the core language).
This can be done like this:
```
pip3 install numpy && pip3 install matplotlib
```

If that's the case, simply run `python3 solution.py --plot` — the program runs in around ~1min.

I've provided in this current folder 4 images of successive runs of the algorithm; the distribution is pretty stable around 8, with some variations (sometimes 13 is the maximum and sometimes not, for example).

To answer question 3, which asked:
> Does it look like any distribution that you know of?

I would say that it looks like a Extreme Value Distribution (EVD), more precisely a Gumbel (or type I) distribution. Otherwise, it may be a slightly asymetric normal distribution.
