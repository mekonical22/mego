package mego

// Use this function if you've variable(s)
// that unused yet,
//
// so Go will not give error about unused variable.
//
// use:
//	x := 1  // x declared and not used
//	mego.UnusedYet(x)  // now x is used
func UnusedYet(a ...any) {}
