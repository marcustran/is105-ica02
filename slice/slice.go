package slice

//import "fmt"

func AllocateVar(b int) []byte {
	var x []byte
	//x = append(x, b)
	x = make([]byte, b, b)
	return x
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte {
	y := make([]byte, b, b)
	return y
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	var s []byte = slc[lidx:uidx]

	return s
}

//func CpySlice()
