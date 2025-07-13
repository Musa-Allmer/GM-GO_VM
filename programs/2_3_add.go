package programs

import "GM/vm"

func Add() {
	virtualM := &vm.VM{
		Code: []byte{
			vm.OP_PUSH, 2,
			vm.OP_PUSH, 3,
			vm.OP_ADD,
			vm.OP_PRINT,
			vm.OP_HALT,
		},
	}
	virtualM.Run()
	// Output: 5
}
