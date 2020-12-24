package matrix

import (
	"fmt"
)

func errRowMissMatch() error {
	return fmt.Errorf("Matrix Error: The given rows statement does not match the actual given element rows")
}
func errColsMissMatch() error {
	return fmt.Errorf("Matrix Error: The given columns statement does not match the actual given element columns")
}
func errNotSquare() error {
	return fmt.Errorf("Matrix Error: The given columns statement does not match the actual given element columns")
}
func errDiffDimensions() error {
	return fmt.Errorf("Matrix Error: cannot summed two or more matrices if the dimensions are different from each other")
}
