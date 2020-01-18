//
// Copyright (c) 2020 estrogently <41487185+estrogently@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package puffy

import "golang.org/x/sys/unix"

func Pledge(promises, execpromises string) error {
	return unix.Pledge(promises, execpromises)
}

func PledgePromises(promises string) error {
	return unix.PledgePromises(promises)
}

func PledgeExecpromises(execpromises string) error {
	return unix.PledgeExecpromises(execpromises)
}

func Unveil(path string, flags string) error {
	return unix.Unveil(path, flags)
}

func UnveilBlock() error {
	return unix.UnveilBlock()
}
