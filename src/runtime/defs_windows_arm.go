// created by cgo -godefs and then converted to Go
// cgo -godefs defs_windows.go
// this file was manually crafted

package runtime

const (
	_PROT_NONE  = 0
	_PROT_READ  = 1
	_PROT_WRITE = 2
	_PROT_EXEC  = 4

	_MAP_ANON    = 1
	_MAP_PRIVATE = 2

	_DUPLICATE_SAME_ACCESS   = 0x2
	_THREAD_PRIORITY_HIGHEST = 0x2

	_SIGINT           = 0x2
	_CTRL_C_EVENT     = 0x0
	_CTRL_BREAK_EVENT = 0x1

	_CONTEXT_CONTROL = 0x10001
	_CONTEXT_FULL    = 0x10007

	_EXCEPTION_ACCESS_VIOLATION     = 0xc0000005
	_EXCEPTION_BREAKPOINT           = 0x80000003
	_EXCEPTION_FLT_DENORMAL_OPERAND = 0xc000008d
	_EXCEPTION_FLT_DIVIDE_BY_ZERO   = 0xc000008e
	_EXCEPTION_FLT_INEXACT_RESULT   = 0xc000008f
	_EXCEPTION_FLT_OVERFLOW         = 0xc0000091
	_EXCEPTION_FLT_UNDERFLOW        = 0xc0000093
	_EXCEPTION_INT_DIVIDE_BY_ZERO   = 0xc0000094
	_EXCEPTION_INT_OVERFLOW         = 0xc0000095

	_INFINITE     = 0xffffffff
	_WAIT_TIMEOUT = 0x102

	_EXCEPTION_CONTINUE_EXECUTION = -0x1
	_EXCEPTION_CONTINUE_SEARCH    = 0x0
)

type systeminfo struct {
	anon0                       [4]byte
	dwpagesize                  uint32
	lpminimumapplicationaddress *byte
	lpmaximumapplicationaddress *byte
	dwactiveprocessormask       uint32
	dwnumberofprocessors        uint32
	dwprocessortype             uint32
	dwallocationgranularity     uint32
	wprocessorlevel             uint16
	wprocessorrevision          uint16
}

type exceptionrecord struct {
	exceptioncode        uint32
	exceptionflags       uint32
	exceptionrecord      *exceptionrecord
	exceptionaddress     *byte
	numberparameters     uint32
	exceptioninformation [15]uint32
}

type neon128 struct {
    low  uint64
    high int64
}

type context struct {
	contextflags      uint32
	r0                uint32
	r1                uint32
	r2                uint32
	r3                uint32
	r4                uint32
	r5                uint32
	r6                uint32
	r7                uint32
	r8                uint32
	r9                uint32
	r10               uint32
	r11               uint32
	r12               uint32

	spr               uint32
	r14               uint32
	pc                uint32
	cpsr              uint32

	fpscr             uint32
	padding           uint32

    floatNeon         [64]neon128

    bvr               [8]uint32
    bcr               [8]uint32
    wvr               [1]uint32
    wcr               [1]uint32
}

func (c *context) q(i uint32) neon128 {
    return c.floatNeon[i]
}

func (c *context) d(i uint32) uint64 {
    if i%2 == 0 {
        return c.floatNeon[i/2].low
    } else {
        return uint64(c.floatNeon[i/2].high)
    }   
}

func (c *context) s(i uint32) uint32 {
    o := i%4
    switch o {
        default: 
            return 0xffffffff
        case 0: 
            return uint32(c.floatNeon[i/2].low)
        case 1: 
            return uint32(c.floatNeon[i/2].low >> 32)
        case 2: 
            return uint32(c.floatNeon[i/2].high)
        case 3: 
            return uint32(c.floatNeon[i/2].high >> 32)
    }
}

func (c *context) ip() uintptr { return uintptr(c.pc) }
func (c *context) sp() uintptr { return uintptr(c.spr) }
func (c *context) lr() uintptr { return uintptr(c.r14) }

func (c *context) setip(x uintptr) { c.pc = uint32(x) }
func (c *context) setsp(x uintptr) { c.spr = uint32(x) }

func dumpregs(r *context) {
	print("r0   ", hex(r.r0),   "\n")
	print("r1   ", hex(r.r1),   "\n")
	print("r2   ", hex(r.r2),   "\n")
	print("r3   ", hex(r.r3),   "\n")
	print("r4   ", hex(r.r4),   "\n")
	print("r5   ", hex(r.r5),   "\n")
	print("r6   ", hex(r.r6),   "\n")
	print("r7   ", hex(r.r7),   "\n")
	print("r8   ", hex(r.r8),   "\n")
	print("r9   ", hex(r.r9),   "\n")
	print("r10  ", hex(r.r10),  "\n")
	print("r11  ", hex(r.r11),  "\n")
	print("r12  ", hex(r.r12),  "\n")
	print("sp   ", hex(r.spr),  "\n")
	print("lr   ", hex(r.r14),   "\n")
	print("pc   ", hex(r.pc),   "\n")
	print("cpsr ", hex(r.cpsr), "\n")
}

type overlapped struct {
	internal     uint32
	internalhigh uint32
	anon0        [8]byte
	hevent       *byte
}