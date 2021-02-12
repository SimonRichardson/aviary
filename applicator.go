package aviary

// Applicator :: (a -> b) -> a -> b
//
// A/apply combinator
func Applicator[A, B any](f func(A) B) func(A) B {
	return func(a A) B {
		return f(a)
	}
}
