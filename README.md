# quartilesolver

Simple brute-force Apple News Quartile puzzle solver. Apple doesn't publish their
dictionary, so it misses some and finds some that Apple doesn't recognize.

```
go build .
./quartilesolver -h
Usage of ./quartilesolver:
  -fragments value
        space-delineated list of fragments
  -in string
        input file containing line-delineated word fragments
  -words string
        file containing line-delineated dictionary of words (default "./assets/usa2.txt")
```

Example usage
```
./quartilesolver -fragments "fo da ery tish on ils nce et ath tan dev ffo re le bu per bre rma ss dil"
Fragments:  [fo da ery tish on ils nce et ath tan dev ffo re le bu per bre rma ss dil]
Total words found:  25
Expected score:  127
breath
breathless
buffoon
buffoonery
buss
daffodil
dale
dance
dare
daredevils
devils
dilettantish
foils
fore
fossils
leathery
leery
leper
less
lesson
on
per
performance
perils
tan
```