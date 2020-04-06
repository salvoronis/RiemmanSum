package main

import(
  "math"
  "fmt"
  "errors"
)

func square(x float64) float64 {
  return x*x
}

func cube(x float64) float64 {
  return x*x*x
}

func calculateError(a,b,e float64) (int, float64, error) {
  if e == 0.0 {
    return 0, 0, errors.New("Деление на 0")
  }
  floatN := (b-a)/math.Pow(e, 0.25)
  n := int(math.Ceil(floatN))
  newerr := math.Pow(((b-a)/float64(n)),4)
  return n, newerr, nil
}

func integrate(a,b,side float64, n int, f func(float64) float64) (float64, float64) {
  h := (b-a)/float64(n)
  dh := (b-a)/float64(n*2)
  var result float64 = 0
  var dresult float64 = 0
  for i := 0; i < n; i++ {
    result += f(a+h*(float64(i)+side))
  }
  result *= h
  for i := 0; i < 2*n; i++ {
    dresult += f(a+dh*(float64(i)+side))
  }
  dresult *= dh
  runge := (math.Abs(result-dresult))/2
  return result, runge
}

func main(){
  var a, b, e, sideInt float64
  var mathFunc, side string
  var f func(float64) float64
  fmt.Println("Выберите функцию (default: square)")
  fmt.Println("square: x^2")
  fmt.Println("cube: x^3")
  fmt.Println("cosinus: cos(x)")
  fmt.Println("sinus: sin(x)")
  fmt.Print("Ваш выбор: ")
  fmt.Scan(&mathFunc)
  switch mathFunc {
  case "cube":
    f = cube
  case "square":
    f = square
  case "cosinus":
    f = math.Cos
  case "sinus":
    f = math.Sin
  default:
    f = square
  }
  fmt.Println("Выберите сторону (default: center)")
  fmt.Println("l: left")
  fmt.Println("r: right")
  fmt.Println("c: center")
  fmt.Print("Ваш выбор: ")
  fmt.Scan(&side)
  switch side {
  case "l":
    sideInt = 0
  case "c":
    sideInt = 0.5
  case "r":
    sideInt = 1
  default:
    sideInt = 0.5
  }
  fmt.Print("Введите нижний и верхний порог: ")
  fmt.Scan(&a, &b)
  fmt.Print("Введите погрешность: ")
  fmt.Scan(&e)
  n, newerr, err := calculateError(a,b,e)
  if err != nil {
    panic(err)
  }
  fmt.Println()
  result, runge := integrate(a,b,sideInt,n,f)
  if math.IsNaN(result) {
    result, newerr, runge = 0,0,0
  }
  fmt.Printf("Результат: %f\n", result)
  fmt.Printf("Количество разбиений: %d\n", n)
  fmt.Printf("Погрешность: %f\n", newerr)
  fmt.Printf("Условие Рунге: %f ", runge)
  if runge < math.E {
    fmt.Println("< e")
  } else {
    fmt.Println("> e")
  }
}
