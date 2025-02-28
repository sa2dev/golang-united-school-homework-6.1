package golang_united_school_homework

import (
	"errors"
	"fmt"
	"strings"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		make([]Shape, 0, shapesCapacity),
		shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of range")
	}
	if b.shapes[i] == nil {
		return nil, errors.New("shape not found")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes[i] = shape

	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var perimeter float64
	for _, v := range b.shapes {
		perimeter += v.CalcPerimeter()
	}
	return perimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64
	for _, v := range b.shapes {
		area += v.CalcArea()
	}
	return area

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	exist := false
	shapeTypes := ""
	var resultShapes []Shape
	for _, v := range b.shapes {
		shapeTypeString := strings.Split(fmt.Sprintf("%T", v), ".")[1]
		shapeTypes += shapeTypeString + ";"

		if shapeTypeString == "Circle" {
			exist = true
		} else {
			resultShapes = append(resultShapes, v)
		}
	}
	if !exist {
		return fmt.Errorf("circles not found in box: %s", shapeTypes)
	}
	b.shapes = resultShapes
	return nil
}
