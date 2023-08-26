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
pkg: github.com/ludi317/arm-benchmarks
                                  │ armv5_1cpu_raw.txt │       armv7soft_1cpu_raw.txt       │
                                  │       sec/op       │   sec/op     vs base               │
Float32Arithmetic                          4.731µ ± 5%   4.681µ ± 0%   -1.05% (p=0.002 n=6)
Int32Arithmetic                            15.65n ± 0%   15.65n ± 0%        ~ (p=0.455 n=6)
Float64Arithmetic                          3.906µ ± 0%   3.876µ ± 0%   -0.78% (p=0.002 n=6)
Int64Arithmetic                            29.06n ± 0%   29.07n ± 0%        ~ (p=0.675 n=6)
ANDconstBICconst                           52.54n ± 0%   52.54n ± 0%        ~ (p=0.991 n=6)
Uint64Move                                 22.36n ± 0%   22.36n ± 0%        ~ (p=0.273 n=6)
ADD                                        1.049µ ± 0%   1.009µ ± 0%   -3.81% (p=0.002 n=6)
ADDBICconst                                20.12n ± 0%   19.00n ± 0%   -5.59% (p=0.002 n=6)
ADDBICconstInt64                           29.07n ± 0%   27.95n ± 0%   -3.85% (p=0.002 n=6)
WithMulDAndMulF                           1029.0n ± 0%   986.2n ± 0%   -4.16% (p=0.002 n=6)
BitwiseInt32                               8.944n ± 0%   8.942n ± 0%   -0.03% (p=0.006 n=6)
BitwiseInt64                               13.42n ± 0%   13.42n ± 0%        ~ (p=1.000 n=6)
TrailingZeros                              43.60n ± 0%   30.18n ± 0%  -30.78% (p=0.002 n=6)
ProducerConsumerBufferedCh                 3.903µ ± 0%   3.635µ ± 0%   -6.87% (p=0.002 n=6)
ProducerConsumerBufferedChInt64            3.954µ ± 0%   3.673µ ± 0%   -7.09% (p=0.002 n=6)
ProducerConsumerUnBufferedCh               5.134µ ± 0%   4.717µ ± 0%   -8.11% (p=0.002 n=6)
ProducerConsumerUnBufferedChInt64          5.117µ ± 0%   4.643µ ± 0%   -9.26% (p=0.002 n=6)
GetCntxct                                  3.878µ ± 0%   3.581µ ± 0%   -7.67% (p=0.002 n=6)
CASInt32                                   158.9n ± 0%   161.0n ± 0%   +1.29% (p=0.002 n=6)
CASInt64                                   502.1n ± 0%   152.2n ± 0%  -69.68% (p=0.002 n=6)
CASUint64                                  502.1n ± 0%   152.2n ± 0%  -69.69% (p=0.002 n=6)
CASUint32                                  158.9n ± 0%   161.0n ± 0%   +1.32% (p=0.002 n=6)
CASUintptr                                 158.9n ± 0%   167.8n ± 0%   +5.60% (p=0.002 n=6)
AtomicOperationsInt64                      931.0n ± 0%   268.6n ± 0%  -71.15% (p=0.002 n=6)
AtomicOperationsInt32                      306.6n ± 0%   297.6n ± 0%   -2.92% (p=0.002 n=6)
AtomicOperationsUint64                     928.8n ± 0%   270.9n ± 0%  -70.84% (p=0.002 n=6)
AtomicOperationsUint32                     306.5n ± 0%   297.6n ± 0%   -2.90% (p=0.002 n=6)
AtomicOperationsUintptr                    308.8n ± 0%   304.4n ± 0%   -1.42% (p=0.002 n=6)
AtomicOperationsBool                       537.0n ± 0%   494.7n ± 0%   -7.89% (p=0.002 n=6)
geomean                                    300.2n        244.9n       -18.43%
