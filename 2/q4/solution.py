# Jonathan Guillotte-Blouin - 7900293

# as found here https://www.ncbi.nlm.nih.gov/Class/FieldGuide/BLOSUM62.txt
BLOSUM62 = {
  "A": {
    "A": 4,
    "C": 0,
    "G": 0,
    "T": 0
  },
  "C": {
    "A": 0,
    "C": 9,
    "G": -3,
    "T": -1
  },
  "G": {
    "A": 0,
    "C": -3,
    "G": 6,
    "T": -2
  },
  "T": {
    "A": 0,
    "C": -1,
    "G": -2,
    "T": 5
  }
}

# as found here http://www.cbcb.umd.edu/confcour/Fall2008/CMSC423-materials/BLOSUM80
BLOSUM80 = {
  "A": {
    "A": 7,
    "C": -1,
    "G": 0,
    "T": 0
  },
  "C": {
    "A": -1,
    "C": 13,
    "G": -6,
    "T": -2
  },
  "G": {
    "A": 0,
    "C": -6,
    "G": 9,
    "T": -3
  },
  "T": {
    "A": 0,
    "C": -2,
    "G": -3,
    "T": 8
  }
}

cBotulinum = "GTGAATCAGCACCTGGACTTTCAGATGAAAAATTAAATTTAACTATCCAAAATGATGCTTATATACCAAAATATGATTCTAATGGAACAAGTGATATAGAACAACATGATGTTAATGAACTTAATGTATTTTTCTATTTAGATGCACAGAAAGTGCCCGAAGGTGAAAATAATGTCAATCTCACCTCTTCAATTGATACAGCATTATTAGAACAACCTAAAATATATACATTTTTTTCATCAGAATTTATTAATAATGTCAATAAACCTGTGCAAGCAGC"
unknownA = "TCTATCAAGTAGATTATTAAATACTACTGCACAAAATGATTCTTACGTTCCAAAATATGATTCTAATGGTACAAGTGAAATAAAGGAATATACTGTTGATAAACTAAATGTATTTTTCTATTTATATGCACAAAAAGCTCCTGAAGGTGAAAGTGCTATAAGTTTAACTTCTTCAGTTAATACAGCATTATTAGATGCATCTAAAGTTTATACGTTTTTTTCTTCAGATTTTATTAATAC"
unknownB = "TCCTGGCTCAGGACGAACGCTGGCGGCGTGTGCTTAACACATGCAAGTCGAGCGATGAAGCTTCCTTCGGGAAGTGGATTAGCGGCGGACGGGTGAGTAACACGTGGGTAACCTGCCTCAAAGTGGGGGATAGCCTTCCGAAAGGAAGATTAATACCGCATAACATAAGAGAATCGCATGATTTTCTTATCAAAGATTTATTGCTTTGAGATGGACCCGCGGCGCATTAGCTAGTTGGTA"

# get total scores with BLOSUM62, then print the results
alignmentScoreA, alignmentScoreB = 0, 0
for i in range(len(unknownA)):
  alignmentScoreA += BLOSUM62[unknownA[i]][cBotulinum[i]]
for j in range(len(unknownB)):
  alignmentScoreB += BLOSUM62[unknownB[j]][cBotulinum[j]]
print('BLOSUM62:\nAlignment score of "unknownA" is {}\nAlignment score "unknownB" is {}\n'.format(alignmentScoreA, alignmentScoreB))

# get total scores with BLOSUM80, then print the results
alignmentScoreA, alignmentScoreB = 0, 0
for i in range(len(unknownA)):
  alignmentScoreA += BLOSUM80[unknownA[i]][cBotulinum[i]]
for j in range(len(unknownB)):
  alignmentScoreB += BLOSUM80[unknownB[j]][cBotulinum[j]]
print('BLOSUM80:\nAlignment score of "unknownA" is {}\nAlignment score "unknownB" is {}'.format(alignmentScoreA, alignmentScoreB))
