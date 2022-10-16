package main

import (
  "fmt"
  "github.com/yyycccxxx/baby-naming/naming"
  naming_v2 "github.com/yyycccxxx/baby-naming/v2/naming"
  naming_v3 "github.com/yyycccxxx/baby-naming/v3/naming"
  naming_v4 "github.com/yyycccxxx/baby-naming/v4/naming"
  naming_v5 "github.com/yyycccxxx/baby-naming/v5/naming"
)

func main() {
  fmt.Println("\n-----------v1: CreateBabyName()----------------------------------------")
  fmt.Println(naming.CreateBabyName())

  fmt.Println("\n-----------subDirectory v2: CreateBabyName(male bool)------------------")
  fmt.Println("male: ", naming_v2.CreateBabyName(true))
  fmt.Println("female: ", naming_v2.CreateBabyName(false))
  
  fmt.Println("\n-----------sub directory v3: CreateBabyName(male bool, minLen int)-----")
  fmt.Println("male and minLen is 5: ", naming_v3.CreateBabyName(true, 5))
  fmt.Println("female and minLen is 5: ", naming_v3.CreateBabyName(false, 5))
  
  fmt.Println("\n-----------branch v4: CreateBabyName(male bool)------------------------")
  fmt.Println("male: ", naming_v4.CreateBabyName(true))
  fmt.Println("female: ", naming_v4.CreateBabyName(false))
  
  fmt.Println("\n------------branch v5: CreateBabyName(male bool, minLen int)-----------")
  fmt.Println("male and minLen is 3: ", naming_v5.CreateBabyName(true, 3))
  fmt.Println("female and minLen is 3: ", naming_v5.CreateBabyName(false, 3))
}
