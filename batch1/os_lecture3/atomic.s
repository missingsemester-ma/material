	.arch armv8-a
	.file	"atomic_example.c"
// GNU C17 (Debian 10.2.1-6) version 10.2.1 20210110 (aarch64-linux-gnu)
//	compiled by GNU C version 10.2.1 20210110, GMP version 6.2.1, MPFR version 4.1.0, MPC version 1.2.0, isl version isl-0.23-GMP

// GGC heuristics: --param ggc-min-expand=100 --param ggc-min-heapsize=131072
// options passed:  -imultiarch aarch64-linux-gnu -D_REENTRANT
// atomic_example.c -mlittle-endian -mabi=lp64 -auxbase-strip atomic.s
// -fverbose-asm -fasynchronous-unwind-tables
// options enabled:  -fPIC -fPIE -faggressive-loop-optimizations
// -fallocation-dce -fasynchronous-unwind-tables -fauto-inc-dec
// -fdelete-null-pointer-checks -fdwarf2-cfi-asm -fearly-inlining
// -feliminate-unused-debug-symbols -feliminate-unused-debug-types
// -ffp-int-builtin-inexact -ffunction-cse -fgcse-lm -fgnu-unique -fident
// -finline-atomics -fipa-stack-alignment -fira-hoist-pressure
// -fira-share-save-slots -fira-share-spill-slots -fivopts
// -fkeep-static-consts -fleading-underscore -flifetime-dse -fmath-errno
// -fmerge-debug-strings -fomit-frame-pointer -fpeephole -fplt
// -fprefetch-loop-arrays -freg-struct-return
// -fsched-critical-path-heuristic -fsched-dep-count-heuristic
// -fsched-group-heuristic -fsched-interblock -fsched-last-insn-heuristic
// -fsched-rank-heuristic -fsched-spec -fsched-spec-insn-heuristic
// -fsched-stalled-insns-dep -fschedule-fusion -fsemantic-interposition
// -fshow-column -fshrink-wrap-separate -fsigned-zeros
// -fsplit-ivs-in-unroller -fssa-backprop -fstdarg-opt
// -fstrict-volatile-bitfields -fsync-libcalls -ftrapping-math
// -ftree-cselim -ftree-forwprop -ftree-loop-if-convert -ftree-loop-im
// -ftree-loop-ivcanon -ftree-loop-optimize -ftree-parallelize-loops=
// -ftree-phiprop -ftree-reassoc -ftree-scev-cprop -funit-at-a-time
// -funwind-tables -fverbose-asm -fzero-initialized-in-bss
// -mfix-cortex-a53-835769 -mfix-cortex-a53-843419 -mglibc -mlittle-endian
// -momit-leaf-frame-pointer -moutline-atomics -mpc-relative-literal-loads

	.text
	.global	shared_variable
	.bss
	.align	2
	.type	shared_variable, %object
	.size	shared_variable, 4
shared_variable:
	.zero	4
	.global	__aarch64_ldadd4_acq_rel
	.text
	.align	2
	.global	thread_function
	.type	thread_function, %function
thread_function:
.LFB6:
	.cfi_startproc
	stp	x29, x30, [sp, -48]!	//,,,
	.cfi_def_cfa_offset 48
	.cfi_offset 29, -48
	.cfi_offset 30, -40
	mov	x29, sp	//,
	str	x0, [sp, 24]	// arg, arg
// atomic_example.c:11:   for (int i = 0; i < ITERATIONS; ++i) {
	str	wzr, [sp, 44]	//, i
// atomic_example.c:11:   for (int i = 0; i < ITERATIONS; ++i) {
	b	.L2		//
.L3:
// atomic_example.c:13:     __atomic_add_fetch(&shared_variable, 1, __ATOMIC_SEQ_CST);
	adrp	x0, shared_variable	// tmp95,
	add	x0, x0, :lo12:shared_variable	// tmp94, tmp95,
	mov	x1, x0	//, tmp94
	mov	w0, 1	//,
	bl	__aarch64_ldadd4_acq_rel		//
// atomic_example.c:11:   for (int i = 0; i < ITERATIONS; ++i) {
	ldr	w0, [sp, 44]	// tmp97, i
	add	w0, w0, 1	// tmp96, tmp97,
	str	w0, [sp, 44]	// tmp96, i
.L2:
// atomic_example.c:11:   for (int i = 0; i < ITERATIONS; ++i) {
	ldr	w1, [sp, 44]	// tmp98, i
	mov	w0, 34463	// tmp99,
	movk	w0, 0x1, lsl 16	// tmp99,,
	cmp	w1, w0	// tmp98, tmp99
	ble	.L3		//,
// atomic_example.c:17:   return NULL;
	mov	x0, 0	// _5,
// atomic_example.c:18: }
	ldp	x29, x30, [sp], 48	//,,,
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
	.string	"pthread_create"
	.align	3
.LC1:
	.string	"pthread_join"
	.align	3
.LC2:
	.string	"Final value of the shared variable: %d\n"
	.text
	.align	2
	.global	main
	.type	main, %function
main:
.LFB7:
	.cfi_startproc
	stp	x29, x30, [sp, -48]!	//,,,
	.cfi_def_cfa_offset 48
	.cfi_offset 29, -48
	.cfi_offset 30, -40
	mov	x29, sp	//,
// atomic_example.c:23:   for (int i = 0; i < NUM_THREADS; ++i) {
	str	wzr, [sp, 44]	//, i
// atomic_example.c:23:   for (int i = 0; i < NUM_THREADS; ++i) {
	b	.L6		//
.L8:
// atomic_example.c:24:     if (pthread_create(&threads[i], NULL, thread_function, NULL) != 0) {
	add	x1, sp, 24	// tmp99,,
	ldrsw	x0, [sp, 44]	// tmp100, i
	lsl	x0, x0, 3	// tmp101, tmp100,
	add	x4, x1, x0	// _1, tmp99, tmp101
	mov	x3, 0	//,
	adrp	x0, thread_function	// tmp102,
	add	x2, x0, :lo12:thread_function	//, tmp102,
	mov	x1, 0	//,
	mov	x0, x4	//, _1
	bl	pthread_create		//
// atomic_example.c:24:     if (pthread_create(&threads[i], NULL, thread_function, NULL) != 0) {
	cmp	w0, 0	// _2,
	beq	.L7		//,
// atomic_example.c:25:       perror("pthread_create");
	adrp	x0, .LC0	// tmp103,
	add	x0, x0, :lo12:.LC0	//, tmp103,
	bl	perror		//
// atomic_example.c:26:       exit(EXIT_FAILURE);
	mov	w0, 1	//,
	bl	exit		//
.L7:
// atomic_example.c:23:   for (int i = 0; i < NUM_THREADS; ++i) {
	ldr	w0, [sp, 44]	// tmp105, i
	add	w0, w0, 1	// tmp104, tmp105,
	str	w0, [sp, 44]	// tmp104, i
.L6:
// atomic_example.c:23:   for (int i = 0; i < NUM_THREADS; ++i) {
	ldr	w0, [sp, 44]	// tmp106, i
	cmp	w0, 1	// tmp106,
	ble	.L8		//,
// atomic_example.c:30:   for (int i = 0; i < NUM_THREADS; ++i) {
	str	wzr, [sp, 40]	//, i
// atomic_example.c:30:   for (int i = 0; i < NUM_THREADS; ++i) {
	b	.L9		//
.L11:
// atomic_example.c:31:     if (pthread_join(threads[i], NULL) != 0) {
	ldrsw	x0, [sp, 40]	// tmp107, i
	lsl	x0, x0, 3	// tmp108, tmp107,
	add	x1, sp, 24	// tmp109,,
	ldr	x0, [x1, x0]	// _3, threads[i_7]
	mov	x1, 0	//,
	bl	pthread_join		//
// atomic_example.c:31:     if (pthread_join(threads[i], NULL) != 0) {
	cmp	w0, 0	// _4,
	beq	.L10		//,
// atomic_example.c:32:       perror("pthread_join");
	adrp	x0, .LC1	// tmp110,
	add	x0, x0, :lo12:.LC1	//, tmp110,
	bl	perror		//
// atomic_example.c:33:       exit(EXIT_FAILURE);
	mov	w0, 1	//,
	bl	exit		//
.L10:
// atomic_example.c:30:   for (int i = 0; i < NUM_THREADS; ++i) {
	ldr	w0, [sp, 40]	// tmp112, i
	add	w0, w0, 1	// tmp111, tmp112,
	str	w0, [sp, 40]	// tmp111, i
.L9:
// atomic_example.c:30:   for (int i = 0; i < NUM_THREADS; ++i) {
	ldr	w0, [sp, 40]	// tmp113, i
	cmp	w0, 1	// tmp113,
	ble	.L11		//,
// atomic_example.c:37:   printf("Final value of the shared variable: %d\n", shared_variable);
	adrp	x0, shared_variable	// tmp115,
	add	x0, x0, :lo12:shared_variable	// tmp114, tmp115,
	ldr	w0, [x0]	// shared_variable.0_5, shared_variable
	mov	w1, w0	//, shared_variable.0_5
	adrp	x0, .LC2	// tmp116,
	add	x0, x0, :lo12:.LC2	//, tmp116,
	bl	printf		//
// atomic_example.c:39:   return 0;
	mov	w0, 0	// _14,
// atomic_example.c:40: }
	ldp	x29, x30, [sp], 48	//,,,
	.cfi_restore 30
	.cfi_restore 29
	.cfi_def_cfa_offset 0
	ret	
	.cfi_endproc
.LFE7:
	.size	main, .-main
	.ident	"GCC: (Debian 10.2.1-6) 10.2.1 20210110"
	.section	.note.GNU-stack,"",@progbits
