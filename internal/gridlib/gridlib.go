package gridlib

// Mod is a modulo function for integers that will not return negative values.
func Mod(d, m int) (rem int) {
    rem = d % m
    if rem < 0 {
        rem += m
    }
    return
}
