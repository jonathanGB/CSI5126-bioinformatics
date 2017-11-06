import random
import sys

nucleotides = ["A", "C", "G", "T"]
defaultSubstitutionMatrix = {
  "A": {
    "A": 1,
    "C": -5,
    "G": -1,
    "T": -5
  },
  "C": {
    "A": -5,
    "C": 1,
    "G": -5,
    "T": -1
  },
  "G": {
    "A": -1,
    "C": -5,
    "G": 1,
    "T": -5
  },
  "T": {
    "A": -5,
    "C": -1,
    "G": -5,
    "T": 1
  }
}

# function to compute the local alignment score of strings mSeq and nSeq
# defaults to defaultSubstitutionMatrix and to an indelCost of 3
def localAlignmentScore(mSeq, nSeq, S=defaultSubstitutionMatrix, indelCost=3):
  m, n, maxScore = len(mSeq), len(nSeq), 0
  V = [[0 for _ in range(n+1)] for _ in range(m+1)] # initialize dynamic table to 0
  
  for i in range(1, m+1):
    for j in range (1, n+1):
      # get the max of all possible scores, and keep the maximum score of the matrix
      score = max(0,
                  V[i-1][j] - indelCost,
                  V[i][j-1] - indelCost,
                  V[i-1][j-1] + S[mSeq[i-1]][nSeq[j-1]])
      maxScore = max(maxScore, score)
      V[i][j] = score
  
  return maxScore


# generate `num` pair of sequences of length m and n
# defaults to 1000 sequences of length 280 and 240
def generateSequences(num=1000, m=280, n=240):
  sequences = []
  for _ in range(num):
    # make random sequence of length m
    first = [random.choice(nucleotides) for _ in range(m)]

    # make random sequence of length n
    second = [random.choice(nucleotides) for _ in range(n)]

    # push tuple in list of sequences
    sequences.append((first, second))

  return sequences


# script starts here (takes ~1min to run with default params)
if len(sys.argv) != 2 or sys.argv[1] != "--plot":
  sys.exit("flag '--plot' not provided. Providing this flag requires having numpy and matplotlib to be installed")

import numpy as np
import matplotlib.pyplot as plt

# store the score of all pair of sequences 
scores = [localAlignmentScore(first, second) for first, second in generateSequences()]
# get the min and max scores out of all scores to arrange the x-axis of the histogram
lowerBound, upperBound = min(scores), max(scores) 

# build the distribution graph and show it
bins = np.arange(lowerBound, upperBound+2, 1)
hist, _ = np.histogram(scores, bins)

freq = hist / hist.sum()

plt.bar(bins[:-1], freq, width=1, ec="k")
plt.show()
