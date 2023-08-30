## Background
These benchmarks were written in support of Go proposal [#61588](https://github.com/golang/go/issues/61588), titled "proposal: runtime: software floating point for GOARM=6, 7 (not only GOARM=5)". A Change List (CL) linked in the proposal introduces new soft float options for ARM targets: `GOARM=6,softfloat`and `GOARM=7,softfloat`.

The benchmarks measure the performance of various basic operations, including arithmetic, bitwise, addition, channel-based, and atomic operations. Performance was compared between benchmarks compiled with `GOARM=7,softfloat` versus `GOARM=5`. As of now, `GOARM=5` remains the only officially supported way to get soft float on ARM targets.

Benchmarks were run on a Raspberry Pi 2 with the following system details: Linux rpi2-20230613 6.1.0-11-armmp #1 SMP Debian 6.1.38-4 (2023-08-08) armv7l.
The option `-test.cpu 1` was used because the intended target has a single CPU. The intended target, and inspiration for this proposal, is an armv7 device without a hardware FPU.
## Commands
```bash
# GOARM=5 benchmarks
GOARCH=arm GOARM=5 GOOS=linux go test -c -o armv5.test
./armv5.test -test.bench . -test.v -test.count 6 -test.cpu 1 | tee armv5_1cpu_raw.txt

# GOARM=7,softfloat benchmarks
GOARCH=arm GOARM=7,softfloat GOOS=linux go test -c -o armv7soft.test
./armv7soft.test -test.bench . -test.v -test.count 6 -test.cpu 1 | tee armv7soft_1cpu_raw.txt

benchstat armv5_1cpu_raw.txt armv7soft_1cpu_raw.txt
```

## Results

```
goos: linux
goarch: arm
pkg: github.com/ludi317/arm-wrestle
                                  │ armv5_1cpu_raw.txt │       armv7soft_1cpu_raw.txt       │
                                  │       sec/op       │   sec/op     vs base               │
Float32Arithmetic                          4.728µ ± 0%   4.689µ ± 0%   -0.80% (p=0.002 n=6)
Int32Arithmetic                            15.65n ± 0%   15.69n ± 0%   +0.26% (p=0.002 n=6)
Float64Arithmetic                          3.907µ ± 0%   3.877µ ± 0%   -0.77% (p=0.002 n=6)
Int64Arithmetic                            29.07n ± 0%   29.07n ± 0%        ~ (p=0.740 n=6)
ANDconstBICconst                           52.55n ± 0%   52.55n ± 0%        ~ (p=0.974 n=6)
Uint64Move                                 22.36n ± 0%   22.36n ± 0%        ~ (p=0.803 n=6)
ADD                                        1.049µ ± 0%   1.008µ ± 0%   -3.91% (p=0.002 n=6)
ADDBICconst                                20.12n ± 0%   19.01n ± 0%   -5.52% (p=0.002 n=6)
ADDBICconstInt64                           29.07n ± 0%   27.95n ± 0%   -3.85% (p=0.002 n=6)
WithMulDAndMulF                           1029.0n ± 0%   988.8n ± 0%   -3.91% (p=0.002 n=6)
BitwiseInt32                               8.944n ± 0%   8.945n ± 0%        ~ (p=0.909 n=6)
BitwiseInt64                               13.41n ± 0%   13.42n ± 0%   +0.07% (p=0.013 n=6)
TrailingZeros                              43.59n ± 0%   30.19n ± 0%  -30.74% (p=0.002 n=6)
ProducerConsumerBufferedCh                 3.875µ ± 0%   3.494µ ± 0%   -9.83% (p=0.002 n=6)
ProducerConsumerBufferedChInt64            3.921µ ± 0%   3.582µ ± 0%   -8.66% (p=0.002 n=6)
ProducerConsumerUnBufferedCh               5.099µ ± 0%   4.549µ ± 0%  -10.80% (p=0.002 n=6)
ProducerConsumerUnBufferedChInt64          5.104µ ± 0%   4.538µ ± 0%  -11.08% (p=0.002 n=6)
GetCntxct                                  3.884µ ± 8%   3.479µ ± 7%  -10.42% (p=0.002 n=6)
CASInt32                                   158.9n ± 0%   161.0n ± 0%   +1.32% (p=0.002 n=6)
CASInt64                                   502.1n ± 1%   154.4n ± 0%  -69.24% (p=0.002 n=6)
CASUint64                                  502.1n ± 0%   152.2n ± 0%  -69.68% (p=0.002 n=6)
CASUint32                                  158.9n ± 0%   161.0n ± 0%   +1.32% (p=0.002 n=6)
CASUintptr                                 159.0n ± 0%   165.4n ± 0%   +4.03% (p=0.002 n=6)
AtomicOperationsInt64                      932.3n ± 0%   270.9n ± 0%  -70.94% (p=0.002 n=6)
AtomicOperationsInt32                      306.9n ± 0%   297.6n ± 0%   -3.01% (p=0.002 n=6)
AtomicOperationsUint64                     930.2n ± 0%   268.6n ± 0%  -71.12% (p=0.002 n=6)
AtomicOperationsUint32                     307.0n ± 0%   297.8n ± 0%   -3.00% (p=0.002 n=6)
AtomicOperationsUintptr                    308.9n ± 0%   302.2n ± 0%   -2.17% (p=0.002 n=6)
AtomicOperationsBool                       538.0n ± 0%   496.8n ± 0%   -7.65% (p=0.002 n=6)
Mutex                                      44.94µ ± 0%   22.60µ ± 0%  -49.70% (p=0.002 n=6)
RWMutex_Read                               45.00µ ± 0%   22.65µ ± 0%  -49.67% (p=0.002 n=6)
RWMutex_Write                              45.22µ ± 0%   22.87µ ± 0%  -49.42% (p=0.002 n=6)
WaitGroup                                  90.63m ± 4%   77.13m ± 4%  -14.89% (p=0.002 n=6)
Channel                                    8.781m ± 0%   8.383m ± 0%   -4.54% (p=0.002 n=6)
AtomicAdd                                 259.40n ± 0%   73.86n ± 0%  -71.53% (p=0.002 n=6)
Once                                       67.11n ± 0%   64.87n ± 0%   -3.33% (p=0.002 n=6)
Cond                                      11.126µ ± 0%   9.781µ ± 0%  -12.09% (p=0.002 n=6)
Pool                                       774.5n ± 1%   723.2n ± 1%   -6.62% (p=0.002 n=6)
MutexContended                             233.4n ± 0%   233.0n ± 0%   -0.19% (p=0.002 n=6)
RWMutexContendedRead                       278.2n ± 0%   277.6n ± 0%   -0.22% (p=0.002 n=6)
RWMutexContendedWrite                      494.2n ± 0%   519.5n ± 0%   +5.12% (p=0.002 n=6)
Semaphore                                  964.6n ± 0%   955.6n ± 0%   -0.93% (p=0.002 n=6)
Mutex2                                     212.5n ± 0%   210.3n ± 0%   -1.04% (p=0.002 n=6)
RWMutex                                    254.9n ± 0%   243.8n ± 0%   -4.37% (p=0.002 n=6)
Channel2                                   964.5n ± 0%   955.6n ± 0%   -0.92% (p=0.002 n=6)
MapRWMutex/Write                           807.7n ± 0%   778.9n ± 0%   -3.56% (p=0.002 n=6)
MapRWMutex/Read                            332.0n ± 0%   331.5n ± 0%   -0.15% (p=0.002 n=6)
MapMutex                                   501.8n ± 0%   518.8n ± 0%   +3.37% (p=0.002 n=6)
geomean                                    757.6n        617.4n       -18.50%
```