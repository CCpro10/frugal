// +build go1.16,!go1.17

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package atm

import (
    `github.com/chenzhuoyu/iasm/x86_64`
)

/** Stack Checking **/

const (
    _M_memcpyargs  = 24
    _G_stackguard0 = 0x10
)

var _S_getg = []byte {
    0x65, 0x48, 0x8b, 0x0c, 0x25, 0x30, 0x00, 0x00, 0x00,   // MOVQ %gs:0x30, %rcx
}

func (self *CodeGen) abiStackCheck(p *x86_64.Program, to *x86_64.Label, sp uintptr) {
    p.Data (_S_getg)
    p.LEAQ (Ptr(RSP, -self.ctxt.size() - int32(sp)), RAX)
    p.CMPQ (Ptr(RCX, _G_stackguard0), RAX)
    p.JBE  (to)
}

/** Efficient Block Copy Algorithm **/

func (self *CodeGen) abiBlockCopy(p *x86_64.Program, pd PointerRegister, ps PointerRegister, nb GenericRegister) {
    rd := self.r(pd)
    rs := self.r(ps)
    rl := self.r(nb)

    /* save all the registers, if they will be clobbered */
    for _, lr := range self.ctxt.regs {
        if rr := self.r(lr); R_memcpy[rr] {
            p.MOVQ(rr, self.ctxt.slot(lr))
        }
    }

    /* load the args and call the function */
    p.MOVQ(rd, Ptr(RSP, 0))
    p.MOVQ(rs, Ptr(RSP, 8))
    p.MOVQ(rl, Ptr(RSP, 16))
    p.MOVQ(F_memmove, RDI)
    p.CALLQ(RDI)

    /* restore all the registers, if they were clobbered */
    for _, lr := range self.ctxt.regs {
        if rr := self.r(lr); R_memcpy[rr] {
            p.MOVQ(self.ctxt.slot(lr), rr)
        }
    }
}
