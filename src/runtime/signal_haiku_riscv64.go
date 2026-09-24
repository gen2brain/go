// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

type sigctxt struct {
	info *siginfo
	ctxt unsafe.Pointer
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) regs() *mcontext {
	return &(*ucontext)(c.ctxt).uc_mcontext
}

func (c *sigctxt) ra() uint64  { return c.regs().x[0] }
func (c *sigctxt) sp() uint64  { return c.regs().x[1] }
func (c *sigctxt) gp() uint64  { return c.regs().x[2] }
func (c *sigctxt) tp() uint64  { return c.regs().x[3] }
func (c *sigctxt) t0() uint64  { return c.regs().x[4] }
func (c *sigctxt) t1() uint64  { return c.regs().x[5] }
func (c *sigctxt) t2() uint64  { return c.regs().x[6] }
func (c *sigctxt) s0() uint64  { return c.regs().x[7] }
func (c *sigctxt) s1() uint64  { return c.regs().x[8] }
func (c *sigctxt) a0() uint64  { return c.regs().x[9] }
func (c *sigctxt) a1() uint64  { return c.regs().x[10] }
func (c *sigctxt) a2() uint64  { return c.regs().x[11] }
func (c *sigctxt) a3() uint64  { return c.regs().x[12] }
func (c *sigctxt) a4() uint64  { return c.regs().x[13] }
func (c *sigctxt) a5() uint64  { return c.regs().x[14] }
func (c *sigctxt) a6() uint64  { return c.regs().x[15] }
func (c *sigctxt) a7() uint64  { return c.regs().x[16] }
func (c *sigctxt) s2() uint64  { return c.regs().x[17] }
func (c *sigctxt) s3() uint64  { return c.regs().x[18] }
func (c *sigctxt) s4() uint64  { return c.regs().x[19] }
func (c *sigctxt) s5() uint64  { return c.regs().x[20] }
func (c *sigctxt) s6() uint64  { return c.regs().x[21] }
func (c *sigctxt) s7() uint64  { return c.regs().x[22] }
func (c *sigctxt) s8() uint64  { return c.regs().x[23] }
func (c *sigctxt) s9() uint64  { return c.regs().x[24] }
func (c *sigctxt) s10() uint64 { return c.regs().x[25] }
func (c *sigctxt) s11() uint64 { return c.regs().x[26] }
func (c *sigctxt) t3() uint64  { return c.regs().x[27] }
func (c *sigctxt) t4() uint64  { return c.regs().x[28] }
func (c *sigctxt) t5() uint64  { return c.regs().x[29] }
func (c *sigctxt) t6() uint64  { return c.regs().x[30] }

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) pc() uint64 { return c.regs().pc }

func (c *sigctxt) sigcode() uint64 { return uint64(c.info.si_code) }
func (c *sigctxt) sigaddr() uint64 { return uint64(c.info.si_addr) }

func (c *sigctxt) set_pc(x uint64) { c.regs().pc = x }
func (c *sigctxt) set_ra(x uint64) { c.regs().x[0] = x }
func (c *sigctxt) set_sp(x uint64) { c.regs().x[1] = x }
func (c *sigctxt) set_gp(x uint64) { c.regs().x[2] = x }

func (c *sigctxt) set_sigcode(x uint32) { c.info.si_code = int32(x) }
func (c *sigctxt) set_sigaddr(x uint64) { c.info.si_addr = uintptr(x) }
