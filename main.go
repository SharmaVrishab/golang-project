package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type value struct {
	Value1 string 
	Value2 string 
} 

func New(value1,value2 string)(value,error){
	if value1 == "" || value2 == ""{
		return value{},errors.New("values are missing")
	}
	return value{value1,value2},nil
}
func main(){
	fmt.Println("starting point")
	input1  := getUserInput("pls enter your string ")
	input2 := getUserInput("pls enter your string ")
	c, err  := New(input1,input2)
	if err != nil {
		fmt.Println("Error creating value:", err)
		return
	}
	WrittingToFile(c)
	WrittingToFileJson(c)
	
}

func getUserInput(prompt string) string {
	fmt.Printf("%v",prompt);
	reader:= bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')
	if err != nil {
		return "error"
	}
	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")
	return text
}
func WrittingToFile(c value) {
	content := fmt.Sprintf("Value1: %s, Value2: %s", c.Value1, c.Value2)
	err := os.WriteFile("secret.txt",[]byte(content),0644);
	if err != nil {
		customErr := errors.New("error while writing to file: " + err.Error())
		fmt.Println(customErr)
		return
	}
	fmt.Println("writting to file was successful")
}
func WrittingToFileJson(c value) {
	// x:= "this is a value of string"
	json, err := json.Marshal(c)
	if err != nil {
		return 
	}
	err2 := os.WriteFile("secret.json",json,0644);
	if err2 != nil {
		customErr := errors.New("error while writing to file: " + err2.Error())
		fmt.Println(customErr)
		return
	}
	fmt.Println("writting to file was successful")
}