package aviary

// Idiot :: a -> a
//
// Identity combinator
func Idiot[T any](a T) T {
	return a
}