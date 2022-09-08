# patterns

String generator from a regex-like pattern.

[![Go Documentation](https://godocs.io/git.sr.ht/~mendelmaleh/patterns?status.svg)](https://godocs.io/git.sr.ht/~mendelmaleh/patterns)

## usage

```bash
$ go install git.sr.ht/~mendelmaleh/patterns/cmd/patterns@latest
$ patterns -h
Usage of patterns:
  -d	debug mode
  -i	interactive mode
  -x int
    	number of strings to generate (default 3)
$ patterns '0x[0-9a-f]{8}'
0x1e47bc0c
0x6abafc21
0x274e8542
```

## 0xdeadbeef

``` bash
$ hyperfine 'patterns -x 999999999999 0x[a-f]{8} | grep -m1 deadbeef'
Benchmark 1: patterns -x 999999999999 0x[a-f]{8} | grep -m1 deadbeef
  Time (mean ± σ):      4.537 s ±  4.220 s    [User: 4.145 s, System: 2.086 s]
  Range (min … max):    0.094 s … 12.529 s    10 runs
```
