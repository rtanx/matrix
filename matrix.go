package matrix

import (
	"fmt"
	"strings"
)

// Matrix store matrix elements and all the needed information
type Matrix struct {
	Element  [][]float64
	rows     int
	cols     int
	isSquare bool
}

// type Matrix interface {
// 	New(cols, rows int, element [][]float64) matrix
// }

func (m Matrix) String() string {

	var sraw string
	for _, el := range m.Element {
		sraw += fmt.Sprintf("%v\n", el)
	}
	return strings.TrimSuffix(sraw, "\n") //fmt.Sprintf("%v", m.Element)
}

// New initializes the matrix according to the given element and return as a new matrix
func New(rows, cols int, element [][]float64) *Matrix {
	m := Matrix{
		cols:     cols,
		rows:     rows,
		Element:  element,
		isSquare: (cols == rows),
	}
	m.validateCols()
	m.validateRows()
	return &m
}

// NewIdentity initializes and return a new identity matrix
func NewIdentity(size int) *Matrix {
	elem := make([][]float64, size)
	for r, el := range elem {
		for i := 0; i < size; i++ {
			if r == i {
				el = append(el, 1)
			} else {
				el = append(el, 0)
			}
		}
		elem[r] = el
	}
	return New(size, size, elem)
}

// Rows return the size Row of the matrix
func (m *Matrix) Rows() int {
	return m.rows
}

// Cols return the size Row of the matrix
func (m *Matrix) Cols() int {
	return m.cols
}

// CheckSameDimension check all given matrices that they all have the same dimensions,
// returning false if there are one or more that are different from other and vice versa
func CheckSameDimension(matrices ...*Matrix) bool {
	var same bool
	for i, m := range matrices {
		if (i + 1) < len(matrices) {
			same = (m.rows == matrices[i+1].rows) && (m.cols == matrices[i+1].cols)
			if !same {
				break
			}
		}
	}
	return same
}

// IsSquare check if the matrix is a square matrix
func (m *Matrix) IsSquare() bool {
	return m.cols == m.rows
}

// validateRows validate whether the size of the row given in the argument matches the number of rows in the element
func (m *Matrix) validateRows() {
	if m.rows != len(m.Element) {
		panic(errRowMissMatch())
	}
}

// validateCols validate whether the size of the column given in the argument matches the number of column in the element
func (m *Matrix) validateCols() {
	for _, col := range m.Element {
		if len(col) != m.cols {
			panic(errColsMissMatch())
		}
	}
}

// Copy return the copy of the matrix
func Copy(m *Matrix) *Matrix {
	el := make([][]float64, m.rows)
	for i := range m.Element {
		el[i] = make([]float64, m.cols)
		copy(el[i], m.Element[i])
	}
	return New(m.rows, m.cols, el)
}

// Det return determinant of the given matrix
func (m *Matrix) Det() float64 {
	if !m.isSquare {
		panic("not square matrix")
	}
	// if then matrix has 2x2 dimension,multiply top-left to bottom-right diagonal,
	// then subtract the product bottom-left to top-right diagonal directly
	if (m.cols == 2) && (m.rows == 2) {
		return (m.Element[0][0] * m.Element[1][1]) - (m.Element[0][1] * m.Element[1][0])
	}
	var sum float64

	cofct := make([]float64, m.rows)
	copy(cofct, m.Element[0])

	for i := 0; i < m.rows; i++ {
		if i%2 != 0 {
			cofct[i] *= -1
		}
		var rowMinor [][]float64
		for j, A := range m.Element {
			if j == 0 {
				continue
			}
			var colMinor []float64
			for k, B := range A {
				if k == i {
					continue
				}
				colMinor = append(colMinor, B)
			}
			rowMinor = append(rowMinor, colMinor)
		}
		minorSize := len(rowMinor)
		sum += cofct[i] * New(minorSize, minorSize, rowMinor).Det()
	}
	return sum
}

// Transpose the given matrix
func (m *Matrix) Transpose() {

}

//Sum summing n matrices and returning the summed matrices as new matrix
func Sum(matrices ...*Matrix) (*Matrix, error) {
	if !CheckSameDimension(matrices...) {
		return nil, errDiffDimensions()
	}
	var elem [][]float64
	for _, m := range matrices {
		for j, row := range m.Element {
			if !(len(elem) > j) {
				elem = append(elem, []float64{})
			}
			for k, el := range row {
				if !(len(elem[j]) > k) {
					elem[j] = append(elem[j], 0)
				}
				elem[j][k] += el
			}
		}
	}
	return New(len(elem), len(elem[0]), elem), nil
}
