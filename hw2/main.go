package main

import (
	"fmt"
	"hw2/calculator"
	"hw2/file"
	"hw2/shapes"
	"strconv"
	"strings"
)

func main() {

	var filePath string
	fmt.Println("Enter file path")
	_, scanErr := fmt.Scan(&filePath)
	if scanErr != nil {
		fmt.Println("scanning error\n", scanErr)
		return
	}

	fileText, fileErr := file.ReadFile(filePath)
	if fileErr != nil {
		fmt.Println("reading error\n", fileErr)
		return
	}

	arrayOfShapes := strings.Split(fileText, "\n")
	var arrayOfCircles []shapes.Circle
	var arrayOfRectangles []shapes.Rectangle

	for i := range arrayOfShapes {
		n := strings.Split(arrayOfShapes[i], " ")
		if n[0] == "circle" {
			par, err := strconv.ParseFloat(n[1], 64)
			if err != nil {
				fmt.Println("parse error\n", par)
				return
			}
			shape := shapes.Circle{Radius: par}
			arrayOfCircles = append(arrayOfCircles, shape)
		}
		if n[0] == "rectangle" {
			par1, err1 := strconv.ParseFloat(n[1], 64)
			if err1 != nil {
				fmt.Println("parse error\n", par1)
				return
			}
			par2, err2 := strconv.ParseFloat(n[2], 64)
			if err2 != nil {
				fmt.Println("parse error\n", par2)
				return
			}
			shape := shapes.Rectangle{Width: par1, Height: par2}
			arrayOfRectangles = append(arrayOfRectangles, shape)
		}
	}

	totalArea, totalAreaErr := calculator.TotalArea(arrayOfCircles, arrayOfRectangles)
	if totalAreaErr != nil {
		fmt.Println(totalAreaErr)
		return
	}
	fmt.Println(totalArea)
}
