package vm

import "fmt"

type VM struct {
	Code  []byte
	ip    int // instruction pointer
	stack []int
}

func (vm *VM) Run() {
	for {
		op := vm.Code[vm.ip]
		vm.ip++
		switch op {
		case OP_PUSH:
			val := int(vm.Code[vm.ip])
			vm.ip++
			vm.stack = append(vm.stack, val)
		case OP_ADD:
			b := vm.pop()
			a := vm.pop()
			vm.push(a + b)
		case OP_PRINT:
			fmt.Println(vm.pop())
		case OP_HALT:
			return
		default:
			panic("unhandled default case")
		}
	}
}
