package handler

import "runtime"

// ============================================================
// STACK TRACE
// ============================================================

func stackTrace() string {

	buf := make(
		[]byte,
		64*1024,
	)

	n := runtime.Stack(
		buf,
		false,
	)

	return string(
		buf[:n],
	)
}
