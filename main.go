package main
import "fmt"
type shaper interface {
  Area( float32)
}

type square struct {
side float32
}
func (sq *squar) float32 {
return sq.side * sq.side
}
func main() {
sq1.side = 5
// var areaintf shaper
//areaintf = sq1horter
// shorter =sq1
// shorter, without seperate declaration:
// areaintf := shaper(sq1)
// or even:
areaintf :=sq1
fmt.Fprintf("The square has area: %f\n", areaintf,Area()

}
