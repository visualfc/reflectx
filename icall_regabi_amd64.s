//go:build go1.17 && goexperiment.regabireflect
// +build go1.17,goexperiment.regabireflect

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"
#include "go_asm.h"

// The frames of each of the two functions below contain two locals, at offsets
// that are known to the runtime.
//
// The first local is a bool called retValid with a whole pointer-word reserved
// for it on the stack. The purpose of this word is so that the runtime knows
// whether the stack-allocated return space contains valid values for stack
// scanning.
//
// The second local is an abi.RegArgs value whose offset is also known to the
// runtime, so that a stack map for it can be constructed, since it contains
// pointers visible to the GC.
#define LOCAL_RETVALID 32
#define LOCAL_REGARGS 40

// spillArgs stores return values from registers to a *internal/abi.RegArgs in R12.
TEXT ·spillArgs(SB),NOSPLIT,$0-0
	MOVQ AX, 0(R12)
	MOVQ BX, 8(R12)
	MOVQ CX, 16(R12)
	MOVQ DI, 24(R12)
	MOVQ SI, 32(R12)
	MOVQ R8, 40(R12)
	MOVQ R9, 48(R12)
	MOVQ R10, 56(R12)
	MOVQ R11, 64(R12)
	MOVQ X0, 72(R12)
	MOVQ X1, 80(R12)
	MOVQ X2, 88(R12)
	MOVQ X3, 96(R12)
	MOVQ X4, 104(R12)
	MOVQ X5, 112(R12)
	MOVQ X6, 120(R12)
	MOVQ X7, 128(R12)
	MOVQ X8, 136(R12)
	MOVQ X9, 144(R12)
	MOVQ X10, 152(R12)
	MOVQ X11, 160(R12)
	MOVQ X12, 168(R12)
	MOVQ X13, 176(R12)
	MOVQ X14, 184(R12)
	RET

// unspillArgs loads args into registers from a *internal/abi.RegArgs in R12.
TEXT ·unspillArgs(SB),NOSPLIT,$0-0
	MOVQ 0(R12), AX
	MOVQ 8(R12), BX
	MOVQ 16(R12), CX
	MOVQ 24(R12), DI
	MOVQ 32(R12), SI
	MOVQ 40(R12), R8
	MOVQ 48(R12), R9
	MOVQ 56(R12), R10
	MOVQ 64(R12), R11
	MOVQ 72(R12), X0
	MOVQ 80(R12), X1
	MOVQ 88(R12), X2
	MOVQ 96(R12), X3
	MOVQ 104(R12), X4
	MOVQ 112(R12), X5
	MOVQ 120(R12), X6
	MOVQ 128(R12), X7
	MOVQ 136(R12), X8
	MOVQ 144(R12), X9
	MOVQ 152(R12), X10
	MOVQ 160(R12), X11
	MOVQ 168(R12), X12
	MOVQ 176(R12), X13
	MOVQ 184(R12), X14
	RET

// makeFuncStub is the code half of the function returned by MakeFunc.
// See the comment on the declaration of makeFuncStub in makefunc.go
// for more details.
// No arg size here; runtime pulls arg map out of the func value.
// This frame contains two locals. See the comment above LOCAL_RETVALID.
#define MAKE_FUNC_FN(NAME,INDEX)		\
TEXT NAME(SB),(NOSPLIT|WRAPPER),$312		\
	NO_LOCAL_POINTERS		\
	LEAQ	LOCAL_REGARGS(SP), R12		\
	CALL	·spillArgs(SB)		\
	MOVQ	24(SP), DX		\
	MOVQ	DX, 0(SP)		\
	LEAQ	argframe+0(FP), CX		\
	MOVQ	CX, 8(SP)		\
	MOVB	$0, LOCAL_RETVALID(SP)		\
	LEAQ	LOCAL_RETVALID(SP), AX		\
	MOVQ	AX, 16(SP)		\
	LEAQ	LOCAL_REGARGS(SP), AX		\
	MOVQ	AX, 24(SP)		\
	MOVQ	$INDEX, AX		\
	MOVQ	AX, 32(SP)		\
	CALL	·i_x(SB)		\
	LEAQ	LOCAL_REGARGS(SP), R12		\
	CALL	·unspillArgs(SB)		\
	RET

MAKE_FUNC_FN(·f0,0)
MAKE_FUNC_FN(·f1,1)
MAKE_FUNC_FN(·f2,2)
MAKE_FUNC_FN(·f3,3)
MAKE_FUNC_FN(·f4,4)
MAKE_FUNC_FN(·f5,5)
MAKE_FUNC_FN(·f6,6)
MAKE_FUNC_FN(·f7,7)
MAKE_FUNC_FN(·f8,8)
MAKE_FUNC_FN(·f9,9)
MAKE_FUNC_FN(·f10,10)
MAKE_FUNC_FN(·f11,11)
MAKE_FUNC_FN(·f12,12)
MAKE_FUNC_FN(·f13,13)
MAKE_FUNC_FN(·f14,14)
MAKE_FUNC_FN(·f15,15)
MAKE_FUNC_FN(·f16,16)
MAKE_FUNC_FN(·f17,17)
MAKE_FUNC_FN(·f18,18)
MAKE_FUNC_FN(·f19,19)
MAKE_FUNC_FN(·f20,20)
MAKE_FUNC_FN(·f21,21)
MAKE_FUNC_FN(·f22,22)
MAKE_FUNC_FN(·f23,23)
MAKE_FUNC_FN(·f24,24)
MAKE_FUNC_FN(·f25,25)
MAKE_FUNC_FN(·f26,26)
MAKE_FUNC_FN(·f27,27)
MAKE_FUNC_FN(·f28,28)
MAKE_FUNC_FN(·f29,29)
MAKE_FUNC_FN(·f30,30)
MAKE_FUNC_FN(·f31,31)
MAKE_FUNC_FN(·f32,32)
MAKE_FUNC_FN(·f33,33)
MAKE_FUNC_FN(·f34,34)
MAKE_FUNC_FN(·f35,35)
MAKE_FUNC_FN(·f36,36)
MAKE_FUNC_FN(·f37,37)
MAKE_FUNC_FN(·f38,38)
MAKE_FUNC_FN(·f39,39)
MAKE_FUNC_FN(·f40,40)
MAKE_FUNC_FN(·f41,41)
MAKE_FUNC_FN(·f42,42)
MAKE_FUNC_FN(·f43,43)
MAKE_FUNC_FN(·f44,44)
MAKE_FUNC_FN(·f45,45)
MAKE_FUNC_FN(·f46,46)
MAKE_FUNC_FN(·f47,47)
MAKE_FUNC_FN(·f48,48)
MAKE_FUNC_FN(·f49,49)
MAKE_FUNC_FN(·f50,50)
MAKE_FUNC_FN(·f51,51)
MAKE_FUNC_FN(·f52,52)
MAKE_FUNC_FN(·f53,53)
MAKE_FUNC_FN(·f54,54)
MAKE_FUNC_FN(·f55,55)
MAKE_FUNC_FN(·f56,56)
MAKE_FUNC_FN(·f57,57)
MAKE_FUNC_FN(·f58,58)
MAKE_FUNC_FN(·f59,59)
MAKE_FUNC_FN(·f60,60)
MAKE_FUNC_FN(·f61,61)
MAKE_FUNC_FN(·f62,62)
MAKE_FUNC_FN(·f63,63)
MAKE_FUNC_FN(·f64,64)
MAKE_FUNC_FN(·f65,65)
MAKE_FUNC_FN(·f66,66)
MAKE_FUNC_FN(·f67,67)
MAKE_FUNC_FN(·f68,68)
MAKE_FUNC_FN(·f69,69)
MAKE_FUNC_FN(·f70,70)
MAKE_FUNC_FN(·f71,71)
MAKE_FUNC_FN(·f72,72)
MAKE_FUNC_FN(·f73,73)
MAKE_FUNC_FN(·f74,74)
MAKE_FUNC_FN(·f75,75)
MAKE_FUNC_FN(·f76,76)
MAKE_FUNC_FN(·f77,77)
MAKE_FUNC_FN(·f78,78)
MAKE_FUNC_FN(·f79,79)
MAKE_FUNC_FN(·f80,80)
MAKE_FUNC_FN(·f81,81)
MAKE_FUNC_FN(·f82,82)
MAKE_FUNC_FN(·f83,83)
MAKE_FUNC_FN(·f84,84)
MAKE_FUNC_FN(·f85,85)
MAKE_FUNC_FN(·f86,86)
MAKE_FUNC_FN(·f87,87)
MAKE_FUNC_FN(·f88,88)
MAKE_FUNC_FN(·f89,89)
MAKE_FUNC_FN(·f90,90)
MAKE_FUNC_FN(·f91,91)
MAKE_FUNC_FN(·f92,92)
MAKE_FUNC_FN(·f93,93)
MAKE_FUNC_FN(·f94,94)
MAKE_FUNC_FN(·f95,95)
MAKE_FUNC_FN(·f96,96)
MAKE_FUNC_FN(·f97,97)
MAKE_FUNC_FN(·f98,98)
MAKE_FUNC_FN(·f99,99)
MAKE_FUNC_FN(·f100,100)
MAKE_FUNC_FN(·f101,101)
MAKE_FUNC_FN(·f102,102)
MAKE_FUNC_FN(·f103,103)
MAKE_FUNC_FN(·f104,104)
MAKE_FUNC_FN(·f105,105)
MAKE_FUNC_FN(·f106,106)
MAKE_FUNC_FN(·f107,107)
MAKE_FUNC_FN(·f108,108)
MAKE_FUNC_FN(·f109,109)
MAKE_FUNC_FN(·f110,110)
MAKE_FUNC_FN(·f111,111)
MAKE_FUNC_FN(·f112,112)
MAKE_FUNC_FN(·f113,113)
MAKE_FUNC_FN(·f114,114)
MAKE_FUNC_FN(·f115,115)
MAKE_FUNC_FN(·f116,116)
MAKE_FUNC_FN(·f117,117)
MAKE_FUNC_FN(·f118,118)
MAKE_FUNC_FN(·f119,119)
MAKE_FUNC_FN(·f120,120)
MAKE_FUNC_FN(·f121,121)
MAKE_FUNC_FN(·f122,122)
MAKE_FUNC_FN(·f123,123)
MAKE_FUNC_FN(·f124,124)
MAKE_FUNC_FN(·f125,125)
MAKE_FUNC_FN(·f126,126)
MAKE_FUNC_FN(·f127,127)
MAKE_FUNC_FN(·f128,128)
MAKE_FUNC_FN(·f129,129)
MAKE_FUNC_FN(·f130,130)
MAKE_FUNC_FN(·f131,131)
MAKE_FUNC_FN(·f132,132)
MAKE_FUNC_FN(·f133,133)
MAKE_FUNC_FN(·f134,134)
MAKE_FUNC_FN(·f135,135)
MAKE_FUNC_FN(·f136,136)
MAKE_FUNC_FN(·f137,137)
MAKE_FUNC_FN(·f138,138)
MAKE_FUNC_FN(·f139,139)
MAKE_FUNC_FN(·f140,140)
MAKE_FUNC_FN(·f141,141)
MAKE_FUNC_FN(·f142,142)
MAKE_FUNC_FN(·f143,143)
MAKE_FUNC_FN(·f144,144)
MAKE_FUNC_FN(·f145,145)
MAKE_FUNC_FN(·f146,146)
MAKE_FUNC_FN(·f147,147)
MAKE_FUNC_FN(·f148,148)
MAKE_FUNC_FN(·f149,149)
MAKE_FUNC_FN(·f150,150)
MAKE_FUNC_FN(·f151,151)
MAKE_FUNC_FN(·f152,152)
MAKE_FUNC_FN(·f153,153)
MAKE_FUNC_FN(·f154,154)
MAKE_FUNC_FN(·f155,155)
MAKE_FUNC_FN(·f156,156)
MAKE_FUNC_FN(·f157,157)
MAKE_FUNC_FN(·f158,158)
MAKE_FUNC_FN(·f159,159)
MAKE_FUNC_FN(·f160,160)
MAKE_FUNC_FN(·f161,161)
MAKE_FUNC_FN(·f162,162)
MAKE_FUNC_FN(·f163,163)
MAKE_FUNC_FN(·f164,164)
MAKE_FUNC_FN(·f165,165)
MAKE_FUNC_FN(·f166,166)
MAKE_FUNC_FN(·f167,167)
MAKE_FUNC_FN(·f168,168)
MAKE_FUNC_FN(·f169,169)
MAKE_FUNC_FN(·f170,170)
MAKE_FUNC_FN(·f171,171)
MAKE_FUNC_FN(·f172,172)
MAKE_FUNC_FN(·f173,173)
MAKE_FUNC_FN(·f174,174)
MAKE_FUNC_FN(·f175,175)
MAKE_FUNC_FN(·f176,176)
MAKE_FUNC_FN(·f177,177)
MAKE_FUNC_FN(·f178,178)
MAKE_FUNC_FN(·f179,179)
MAKE_FUNC_FN(·f180,180)
MAKE_FUNC_FN(·f181,181)
MAKE_FUNC_FN(·f182,182)
MAKE_FUNC_FN(·f183,183)
MAKE_FUNC_FN(·f184,184)
MAKE_FUNC_FN(·f185,185)
MAKE_FUNC_FN(·f186,186)
MAKE_FUNC_FN(·f187,187)
MAKE_FUNC_FN(·f188,188)
MAKE_FUNC_FN(·f189,189)
MAKE_FUNC_FN(·f190,190)
MAKE_FUNC_FN(·f191,191)
MAKE_FUNC_FN(·f192,192)
MAKE_FUNC_FN(·f193,193)
MAKE_FUNC_FN(·f194,194)
MAKE_FUNC_FN(·f195,195)
MAKE_FUNC_FN(·f196,196)
MAKE_FUNC_FN(·f197,197)
MAKE_FUNC_FN(·f198,198)
MAKE_FUNC_FN(·f199,199)
MAKE_FUNC_FN(·f200,200)
MAKE_FUNC_FN(·f201,201)
MAKE_FUNC_FN(·f202,202)
MAKE_FUNC_FN(·f203,203)
MAKE_FUNC_FN(·f204,204)
MAKE_FUNC_FN(·f205,205)
MAKE_FUNC_FN(·f206,206)
MAKE_FUNC_FN(·f207,207)
MAKE_FUNC_FN(·f208,208)
MAKE_FUNC_FN(·f209,209)
MAKE_FUNC_FN(·f210,210)
MAKE_FUNC_FN(·f211,211)
MAKE_FUNC_FN(·f212,212)
MAKE_FUNC_FN(·f213,213)
MAKE_FUNC_FN(·f214,214)
MAKE_FUNC_FN(·f215,215)
MAKE_FUNC_FN(·f216,216)
MAKE_FUNC_FN(·f217,217)
MAKE_FUNC_FN(·f218,218)
MAKE_FUNC_FN(·f219,219)
MAKE_FUNC_FN(·f220,220)
MAKE_FUNC_FN(·f221,221)
MAKE_FUNC_FN(·f222,222)
MAKE_FUNC_FN(·f223,223)
MAKE_FUNC_FN(·f224,224)
MAKE_FUNC_FN(·f225,225)
MAKE_FUNC_FN(·f226,226)
MAKE_FUNC_FN(·f227,227)
MAKE_FUNC_FN(·f228,228)
MAKE_FUNC_FN(·f229,229)
MAKE_FUNC_FN(·f230,230)
MAKE_FUNC_FN(·f231,231)
MAKE_FUNC_FN(·f232,232)
MAKE_FUNC_FN(·f233,233)
MAKE_FUNC_FN(·f234,234)
MAKE_FUNC_FN(·f235,235)
MAKE_FUNC_FN(·f236,236)
MAKE_FUNC_FN(·f237,237)
MAKE_FUNC_FN(·f238,238)
MAKE_FUNC_FN(·f239,239)
MAKE_FUNC_FN(·f240,240)
MAKE_FUNC_FN(·f241,241)
MAKE_FUNC_FN(·f242,242)
MAKE_FUNC_FN(·f243,243)
MAKE_FUNC_FN(·f244,244)
MAKE_FUNC_FN(·f245,245)
MAKE_FUNC_FN(·f246,246)
MAKE_FUNC_FN(·f247,247)
MAKE_FUNC_FN(·f248,248)
MAKE_FUNC_FN(·f249,249)
MAKE_FUNC_FN(·f250,250)
MAKE_FUNC_FN(·f251,251)
MAKE_FUNC_FN(·f252,252)
MAKE_FUNC_FN(·f253,253)
MAKE_FUNC_FN(·f254,254)
MAKE_FUNC_FN(·f255,255)
