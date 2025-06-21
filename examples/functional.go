// Map wendet fn auf jedes Element an und gibt neues Slice zurück.
// - Generics erlauben Map für []int, []string, custom types.
func Map[T any, R any](slice []T, fn func(T) R) []R {
    out := make([]R, len(slice))
    for i, v := range slice {
        out[i] = fn(v)
    }
    return out
}

// Filter behält nur Werte, bei denen keep true zurückgibt.
func Filter[T any](slice []T, keep func(T) bool) []T {
    var out []T
    for _, v := range slice {
        if keep(v) {
            out = append(out, v)
        }
    }
    return out
}

// Reduce faltet slice mit init-Wert und fn zusammen.
func Reduce[T any, R any](slice []T, init R, fn func(R, T) R) R {
    acc := init
    for _, v := range slice {
        acc = fn(acc, v)
    }
    return acc
}