package main

import (

  "fmt"
  "strconv"
  "strings"
  "math/rand"
  "time"

)

func main() {
//initialize variables

// constant comparison operator
  comparison := true

//variable to recieve keyboard input
  input := ""

//Determine if input value from keyboard is valid
  valid_inp := false
//Type of die to ve rolled
  type_die := 0
//Number of dies to be rolled
  num_die := 0
//Deterrmines if user is done with program
  finished := false
//stores result of number generation
  result := 0


  for finished != comparison {

      valid_inp = false
      fmt.Print("Enter the type of die to be rolled (Only enter numerical values, so a d20 would be entered as 20): ")

//Read in the type of die that the user wishes to roll in numerical value only
      for valid_inp != comparison {
        fmt.Scanln(&input)

        //Remove the newline from the read in string
        input = strings.Replace(input,"\n","",-1)

        //convert the string to int for number generation
        type_dice, err := strconv.Atoi(input)

        type_die = type_dice

        //If no error occurs in conversion then set valid_inp to true and break from loop
        if err == nil {
          valid_inp = true
        } else {
          fmt.Print("Invalid value input, please input a numerical value only for type of die to roll: ")
        }

      }

      //Set back to false for later use in a for loop
      valid_inp = false

//Read in the number of selected dies to be rolled
      fmt.Print("Enter the number of die's to be rolled: ")

      for valid_inp != comparison {
        fmt.Scanln(&input)

        //Remove the newline from the read in string
        input = strings.Replace(input,"\n","",-1)

        //convert the string to int for number generation
        num_dice, erro := strconv.Atoi(input)

        num_die = num_dice

        //If no error occurs in conversion then set valid_inp to true and break from loop
        if erro == nil {
          valid_inp = true
        } else {
          fmt.Print("Invalid value input, please input a numerical value only for number of die's to roll: ")
        }

      }

      //Start Number generation
      seed := rand.NewSource(time.Now().UnixNano())
      gen := rand.New(seed)

      for num_die > 0{
        result = gen.Intn(type_die) + 1
        fmt.Print("You have rolled a ")
        fmt.Print(result)
        fmt.Println()
        num_die = num_die - 1
      }

      fmt.Print("Would you like to roll another set of dice, enter y for yes, anything else will terminate program: ")
      fmt.Scanln(&input)

      if input != "y"{
        finished = true
      }


  }
}
