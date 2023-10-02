	.arch armv8-a
	.file	"mutex_example.c"
	.text
	.global	shared_variable
	.bss
	.align	2
	.type	shared_variable, %object
	.size	shared_variable, 4
shared_variable:
	.zero	4
	.global	mutex
	.align	3
	.type	mutex, %object
	.size	mutex, 48
mutex:
	.zero	48
	.text
	.align	2
	.global	thread_function
	.type	thread_function, %function
thread_function:
.LFB6:
	.cfi_startproc
	stp	x29, x30, [sp, -48]!
	.cfi_def_cfa_offset 48
	.cfi_offset 29, -48
	.cfi_offset 30, -40
	mov	x29, sp
	str	x0, [sp, 24]
	str	wzr, [sp, 44]
	b	.L2
.L3:
	adrp	x0, mutex
	add	x0, x0, :lo12:mutex
	bl	pthread_mutex_lock
	adrp	x0, shared_variable
	add	x0, x0, :lo12:shared_variable
	ldr	w0, [x0]
	add	w1, w0, 1
	adrp	x0, shared_variable
	add	x0, x0, :lo12:shared_variable
	str	w1, [x0]
	adrp	x0, mutex
	add	x0, x0, :lo12:mutex
	bl	pthread_mutex_unlock
	ldr	w0, [sp, 44]
	add	w0, w0, 1
	str	w0, [sp, 44]
.L2:
	ldr	w1, [sp, 44]
	mov	w0, 16959
	movk	w0, 0xf, lsl 16
	cmp	w1, w0
	ble	.L3
	mov	x0, 0
	ldp	x29, x30, [sp], 48
	.cfi_restore 30
	.cfi_restore 29
	.cfi_def_cfa_offset 0
	ret
	.cfi_endproc
.LFE6:
	.size	thread_function, .-thread_function
	.section	.rodata
	.align	3
.LC0:
	.string	"pthread_mutex_init"
	.align	3
.LC1:
	.string	"pthread_create"
	.align	3
.LC2:
	.string	"pthread_join"
	.align	3
.LC3:
	.string	"Final value of the shared variable: %d\n"
	.text
	.align	2
	.global	main
	.type	main, %function
main:
.LFB7:
	.cfi_startproc
	stp	x29, x30, [sp, -48]!
	.cfi_def_cfa_offset 48
	.cfi_offset 29, -48
	.cfi_offset 30, -40
	mov	x29, sp
	mov	x1, 0
	adrp	x0, mutex
	add	x0, x0, :lo12:mutex
	bl	pthread_mutex_init
	cmp	w0, 0
	beq	.L6
	adrp	x0, .LC0
	add	x0, x0, :lo12:.LC0
	bl	perror
	mov	w0, 1
	bl	exit
.L6:
	str	wzr, [sp, 44]
	b	.L7
.L9:
	add	x1, sp, 24
	ldrsw	x0, [sp, 44]
	lsl	x0, x0, 3
	add	x4, x1, x0
	mov	x3, 0
	adrp	x0, thread_function
	add	x2, x0, :lo12:thread_function
	mov	x1, 0
	mov	x0, x4
	bl	pthread_create
	cmp	w0, 0
	beq	.L8
	adrp	x0, .LC1
	add	x0, x0, :lo12:.LC1
	bl	perror
	mov	w0, 1
	bl	exit
.L8:
	ldr	w0, [sp, 44]
	add	w0, w0, 1
	str	w0, [sp, 44]
.L7:
	ldr	w0, [sp, 44]
	cmp	w0, 1
	ble	.L9
	str	wzr, [sp, 40]
	b	.L10
.L12:
	ldrsw	x0, [sp, 40]
	lsl	x0, x0, 3
	add	x1, sp, 24
	ldr	x0, [x1, x0]
	mov	x1, 0
	bl	pthread_join
	cmp	w0, 0
	beq	.L11
	adrp	x0, .LC2
	add	x0, x0, :lo12:.LC2
	bl	perror
	mov	w0, 1
	bl	exit
.L11:
	ldr	w0, [sp, 40]
	add	w0, w0, 1
	str	w0, [sp, 40]
.L10:
	ldr	w0, [sp, 40]
	cmp	w0, 1
	ble	.L12
	adrp	x0, shared_variable
	add	x0, x0, :lo12:shared_variable
	ldr	w0, [x0]
	mov	w1, w0
	adrp	x0, .LC3
	add	x0, x0, :lo12:.LC3
	bl	printf
	adrp	x0, mutex
	add	x0, x0, :lo12:mutex
	bl	pthread_mutex_destroy
	mov	w0, 0
	ldp	x29, x30, [sp], 48
	.cfi_restore 30
	.cfi_restore 29
	.cfi_def_cfa_offset 0
	ret
	.cfi_endproc
.LFE7:
	.size	main, .-main
	.ident	"GCC: (Debian 10.2.1-6) 10.2.1 20210110"
	.section	.note.GNU-stack,"",@progbits
