## Background
These benchmarks were written in support of Go proposal [#61588](https://github.com/golang/go/issues/61588), titled "proposal: runtime: software floating point for GOARM=6, 7 (not only GOARM=5)". A Change List (CL) linked in the proposal introduces new soft float options for ARM targets: `GOARM=6,softfloat`and `GOARM=7,softfloat`.

The benchmarks measure the performance of various basic operations, including arithmetic, bitwise, addition, channel-based, and atomic operations. Performance was compared between benchmarks compiled with `GOARM=7,softfloat` versus `GOARM=5`. As of now, `GOARM=5` remains the only officially supported way to get soft float on ARM targets.
### Round 1
Benchmarks were run on a Raspberry Pi 2 with the following system details: Linux rpi2-20230613 6.1.0-11-armmp #1 SMP Debian 6.1.38-4 (2023-08-08) armv7l.
The option `-test.cpu 1` was used because the intended target has a single CPU. The intended target, and inspiration for this proposal, is an armv7 device without a hardware FPU.
#### Commands
```bash
# GOARM=5 benchmarks
GOARCH=arm GOARM=5 GOOS=linux go test -c -o armv5.test
./armv5.test -test.bench . -test.v -test.count 6 -test.cpu 1 | tee armv5_1cpu_raw.txt

# GOARM=7,softfloat benchmarks
GOARCH=arm GOARM=7,softfloat GOOS=linux go test -c -o armv7soft.test
./armv7soft.test -test.bench . -test.v -test.count 6 -test.cpu 1 | tee armv7soft_1cpu_raw.txt

benchstat armv5_1cpu_raw.txt armv7soft_1cpu_raw.txt
```

#### Results

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
### Round 2
The same Raspberry Pi 2 also measured benchmarks on `GOARM=5+keith` which is `GOARM=5` with  [Keith Randall's CL](https://go-review.googlesource.com/c/go/+/525637) applied.

The commit prior to Keith's change was `a2647f08f0c4e540540a7ae1b9ba7e668e6fed80`, so the `GOARM=7,softfloat` patch was rebased on top of that commit too, and rerun, for parity.

System details given by `uname -a`: `Linux raspberrypi 6.1.0-rpi4-rpi-v7 #1 SMP Raspbian 1:6.1.54-1+rpt2 (2023-10-05) armv7l GNU/Linux`. Its OS has been updated since Round 1.
#### Commands
```bash
# GOARM=5+keith benchmarks
# build go with keith's CL applied on top of a2647f08f0c4e540540a7ae1b9ba7e668e6fed80
GOARCH=arm GOARM=5 GOOS=linux go test -c -o armv5keith.test
./armv5keith.test -test.bench . -test.v -test.count 6 -test.cpu 1 | tee armv5keith_1cpu_raw.txt

# GOARM=7,softfloat benchmarks
# build go with GOARM=7,softfloat patch applied on top of a2647f08f0c4e540540a7ae1b9ba7e668e6fed80
GOARCH=arm GOARM=7,softfloat GOOS=linux go test -c -o armv7soft.test
./armv7soft.test -test.bench . -test.v -test.count 6 -test.cpu 1 | tee armv7soft_1cpu_raw.txt

benchstat raw/round2/armv5keith_1cpu_raw.txt raw/round2/armv7soft_1cpu_raw.txt
```

#### Results
```
goos: linux
goarch: arm
pkg: github.com/ludi317/arm-wrestle
cpu: ARMv7 Processor rev 5 (v7l)
                                  │ raw/round2/armv5keith_1cpu_raw.txt │ raw/round2/armv7soft_1cpu_raw.txt  │
                                  │               sec/op               │   sec/op     vs base               │
Float32Arithmetic                                          4.761µ ± 0%   4.715µ ± 0%   -0.98% (p=0.002 n=6)
Int32Arithmetic                                            15.70n ± 0%   15.74n ± 0%   +0.25% (p=0.002 n=6)
Float64Arithmetic                                          3.923µ ± 0%   3.898µ ± 0%   -0.64% (p=0.002 n=6)
Int64Arithmetic                                            29.15n ± 0%   29.15n ± 0%        ~ (p=0.396 n=6)
ANDconstBICconst                                           52.69n ± 0%   52.65n ± 0%        ~ (p=0.266 n=6)
Uint64Move                                                 22.42n ± 0%   22.42n ± 0%        ~ (p=1.000 n=6)
ADD                                                        1.060µ ± 0%   1.007µ ± 0%   -5.09% (p=0.002 n=6)
ADDBICconst                                                20.18n ± 0%   18.95n ± 0%   -6.12% (p=0.002 n=6)
ADDBICconstInt64                                           29.15n ± 0%   27.88n ± 0%   -4.36% (p=0.002 n=6)
WithMulDAndMulF                                           1040.5n ± 0%   984.7n ± 0%   -5.36% (p=0.002 n=6)
BitwiseInt32                                               8.968n ± 0%   8.919n ± 0%   -0.55% (p=0.002 n=6)
BitwiseInt64                                               13.46n ± 0%   13.38n ± 0%   -0.56% (p=0.002 n=6)
TrailingZeros                                              43.72n ± 0%   30.10n ± 0%  -31.15% (p=0.002 n=6)
LeadingZeros                                               42.60n ± 0%   39.14n ± 0%   -8.13% (p=0.002 n=6)
RotateLeft                                                 114.4n ± 0%   111.0n ± 0%   -2.97% (p=0.002 n=6)
OnesCount                                                  150.2n ± 0%   144.7n ± 0%   -3.73% (p=0.002 n=6)
ProducerConsumerBufferedCh                                 3.602µ ± 0%   3.480µ ± 0%   -3.39% (p=0.002 n=6)
ProducerConsumerBufferedChInt64                            3.687µ ± 0%   3.529µ ± 0%   -4.29% (p=0.002 n=6)
ProducerConsumerUnBufferedCh                               4.683µ ± 0%   4.492µ ± 0%   -4.09% (p=0.002 n=6)
ProducerConsumerUnBufferedChInt64                          4.653µ ± 0%   4.490µ ± 0%   -3.50% (p=0.002 n=6)
GetCntxct                                                  3.580µ ± 1%   3.496µ ± 1%   -2.36% (p=0.002 n=6)
CASInt32                                                   158.4n ± 0%   160.5n ± 0%   +1.33% (p=0.002 n=6)
CASInt64                                                   152.6n ± 2%   154.8n ± 0%   +1.44% (p=0.015 n=6)
CASUint64                                                  152.6n ± 1%   151.8n ± 1%        ~ (p=0.117 n=6)
CASUint32                                                  159.3n ± 0%   160.3n ± 0%   +0.63% (p=0.002 n=6)
CASUintptr                                                 159.3n ± 0%   164.8n ± 0%   +3.45% (p=0.002 n=6)
AtomicOperationsInt64                                      269.4n ± 0%   270.0n ± 0%   +0.24% (p=0.002 n=6)
AtomicOperationsInt32                                      307.5n ± 0%   296.8n ± 0%   -3.50% (p=0.002 n=6)
AtomicOperationsUint64                                     269.3n ± 0%   267.7n ± 0%   -0.59% (p=0.002 n=6)
AtomicOperationsUint32                                     307.5n ± 0%   296.9n ± 0%   -3.46% (p=0.002 n=6)
AtomicOperationsUintptr                                    309.7n ± 0%   301.2n ± 1%   -2.74% (p=0.002 n=6)
AtomicOperationsBool                                       539.7n ± 0%   498.2n ± 0%   -7.68% (p=0.002 n=6)
Mutex                                                      45.09µ ± 0%   22.68µ ± 0%  -49.69% (p=0.002 n=6)
RWMutex_Read                                               45.13µ ± 0%   22.71µ ± 1%  -49.69% (p=0.002 n=6)
RWMutex_Write                                              45.37µ ± 0%   22.90µ ± 0%  -49.52% (p=0.002 n=6)
WaitGroup                                                  77.64m ± 6%   77.01m ± 5%        ~ (p=0.394 n=6)
Channel                                                    8.816m ± 0%   8.344m ± 1%   -5.35% (p=0.002 n=6)
AtomicAdd                                                  71.80n ± 0%   73.67n ± 0%   +2.61% (p=0.002 n=6)
Once                                                       67.31n ± 0%   64.98n ± 0%   -3.45% (p=0.002 n=6)
Cond                                                      10.316µ ± 0%   9.796µ ± 0%   -5.04% (p=0.002 n=6)
Pool                                                       787.1n ± 1%   741.6n ± 2%   -5.79% (p=0.002 n=6)
MutexContended                                             233.6n ± 1%   235.1n ± 0%   +0.64% (p=0.002 n=6)
RWMutexContendedRead                                       276.3n ± 0%   278.6n ± 0%   +0.80% (p=0.002 n=6)
RWMutexContendedWrite                                      499.2n ± 0%   521.5n ± 0%   +4.48% (p=0.002 n=6)
Semaphore                                                  970.6n ± 0%   967.6n ± 0%   -0.30% (p=0.002 n=6)
Mutex2                                                     213.1n ± 0%   210.9n ± 0%   -1.03% (p=0.002 n=6)
RWMutex                                                    253.5n ± 0%   242.2n ± 0%   -4.46% (p=0.002 n=6)
Channel2                                                   970.7n ± 0%   968.0n ± 0%   -0.28% (p=0.002 n=6)
MapRWMutex/Write                                           807.4n ± 0%   790.3n ± 0%   -2.12% (p=0.002 n=6)
MapRWMutex/Read                                            335.1n ± 0%   334.7n ± 0%   -0.10% (p=0.002 n=6)
MapMutex                                                   503.9n ± 0%   520.6n ± 0%   +3.32% (p=0.002 n=6)
geomean                                                    586.8n        550.0n        -6.27%

```
