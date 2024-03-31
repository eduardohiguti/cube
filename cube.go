package main

import (
	"fmt"
	"math"
)

var xp, yp, idx int
var A, B, C, x, y, z, ooz float64
var cubeWidth, K1, incrementSpeed = 20.0, 40.0, 0.6
var width, height, distanceFromCam = 160, 44, 100
var zBuffer = make([]float64, 160*44*4)
var buffer = make([]byte, 160*44)
var backgroundASCIICode = ' '

func calculateX(i float64, j float64, k float64) float64 {
	return float64(j)*math.Sin(A)*math.Sin(B)*math.Cos(C) - float64(k)*math.Cos(A)*math.Sin(B)*math.Cos(C) + float64(j)*math.Cos(A)*math.Sin(C) + float64(k)*math.Sin(A)*math.Sin(C) + float64(i)*math.Cos(B)*math.Cos(C)
}

func calculateY(i float64, j float64, k float64) float64 {
	return float64(j)*math.Cos(A)*math.Cos(C) + float64(k)*math.Sin(A)*math.Cos(C) - float64(j)*math.Sin(A)*math.Sin(B)*math.Sin(C) + float64(k)*math.Cos(A)*math.Sin(B)*math.Sin(C) - float64(i)*math.Cos(B)*math.Sin(C)
}

func calculateZ(i float64, j float64, k float64) float64 {
	return float64(k)*math.Cos(A)*math.Cos(B) - float64(j)*math.Sin(A)*math.Cos(B) + float64(i)*math.Sin(B)
}

func calculateForSurface(cubeX float64, cubeY float64, cubeZ float64, ch byte) {
	x = calculateX(cubeX, cubeY, cubeZ)
	y = calculateY(cubeX, cubeY, cubeZ)
	z = calculateZ(cubeX, cubeY, cubeZ) + float64(distanceFromCam)
	ooz = 1 / z
	xp = width/2 + int(K1*ooz*x*2)
	yp = height/2 + int(K1*ooz*y)
	idx = xp + yp*width
	if idx >= 0 && idx < width*height {
		if ooz > zBuffer[idx] {
			zBuffer[idx] = ooz
			buffer[idx] = ch
		}
	}
}

func main() {
	fmt.Printf("\x1b[2J")
	for {
		for i := range buffer {
			buffer[i] = byte(backgroundASCIICode)
		}
		for i := range zBuffer {
			zBuffer[i] = 0
		}
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateForSurface(cubeX, cubeY, -cubeWidth, '.')
				calculateForSurface(cubeWidth, cubeY, cubeX, '$')
				calculateForSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateForSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateForSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateForSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}
		fmt.Printf("\x1b[H")
		for k := 0; k < width*height; k++ {
			if k%width != 0 {
				fmt.Printf("%c", buffer[k])
			} else {
				fmt.Printf("%c", 10)
			}
		}
		A += 0.005
		B += 0.005
	}
}
