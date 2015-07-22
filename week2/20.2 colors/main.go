

package main

"fmt"
"github.com/ttacon/chalk"
"examples/week2/day1/converters"  //need to be made

func main() {
  fmt.Println(chalk.Blue, converters.FeetToKM(23),chalk.ResetColor)
}
