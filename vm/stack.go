package vm

func (vm *VM) pop() int {
	val := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return val
}

func (vm *VM) push(val int) {
	vm.stack = append(vm.stack, val)
}
