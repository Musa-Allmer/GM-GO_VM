package vm

const (
	OP_PUSH byte = iota // Push a constant onto the stack
	OP_ADD              // Pop two, add, push result
	OP_SUB              // Pop two, subtract
	OP_MUL
	OP_DIV
	OP_PRINT // Pop and print top
	OP_HALT  // Stop execution
)
