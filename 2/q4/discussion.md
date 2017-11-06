Jonathan Guillotte-Blouin - 7900293

*NB: this program requires to have python 3*

After running `solution.py`, it seems that both *Unknown A* and *Unknown B* are related to *Clostridium botulinum*. After summing **BLOSUM62** scores at every index in *Unknown A* with *C. botulinum*, and then summing scores for every index in *Unknown B* with *C. botulinum*, both total scores were vastly positive. We repeat this process with **BLOSUM80**, again getting vastly positive scores.

```
BLOSUM62:
Alignment score of "unknownA" is 242
Alignment score "unknownB" is 199

BLOSUM80:
Alignment score of "unknownA" is 351
Alignment score "unknownB" is 268
```

*Unknown A* seems however to be more closely related to *C. botulinum* than *Unknown B*.

I chose **BLOSUM62** because it is a very popular substitution matrix (it is the default matrix for **BLAST**). I then computed with **BLOSUM80** to see if it would go against the previous results, but it in fact strengthened them.