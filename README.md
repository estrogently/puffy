![](https://github.com/estrogently/puffy/workflows/Test/badge.svg)

PUFFY (Pledge & Unveil Facilitated For You) wraps Go's `Pledge`,
`PledgePromises`, `PledgeExecPromises`, `Unveil`, and `UnveilBlock` functions,
making them available on platforms other than OpenBSD. Rather than cause your
build to fail, on those platforms they will simply be no-ops.

This relieves you from having to write OS-specific files each time you want to
use pledge or unveil in a Go project—PUFFY does the work for you.

For further documentation, see golang/x/sys/unix's
[`pledge_openbsd.go`](https://github.com/golang/sys/blob/master/unix/pledge_openbsd.go),
[`unveil_openbsd.go`](https://github.com/golang/sys/blob/master/unix/unveil_openbsd.go),
as well as OpenBSD's [`pledge(2)`](https://man.openbsd.org/pledge.2) and
[`unveil(2)`](https://man.openbsd.org/unveil.2).
