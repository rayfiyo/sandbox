package main

import(
  "github.com/YosCiiCable/sandbox/gosand/goImportTest/targetPKG"
  "fmt"
)


func main(){
  p1 := targetPKG.Person{Name: "Go太郎", Age: 10}
  fmt.Println(p1.Name, p1.Age)
  fmt.Println(targetPKG.Add(3,5))
  fmt.Println(targetPKG.TAX)
}
