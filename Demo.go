package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Student struct {
	name          string
	age           int
	dept          string
	contactnumber int
}

func main() {
	start := time.Now()
	var num int
	var name string
	var passcode string
	check := 1
	var stud Student
	//fmt.Print("Enter your name: ")
	//fmt.Scanf( "%d", &num)
	//fmt.Println("Hello", num)
	fmt.Println("Please enter your name")
	fmt.Scanf("%s", &name)
	if name == "Alex" {
		fmt.Printf("\nHi %s Welcome to Go Demo... \n", name)
		fmt.Printf("%s enter passcode to proceed : \n", name)
		fmt.Scan(&passcode)
		if passcode == "let_me_in" {
			for check == 1 {
				fmt.Printf("Enter the topic to view Demo\n\n")
				listoftopic()
				fmt.Scan(&num)
				var topic string = gettopicfromnum(num)
				fmt.Printf("\n%s you chose %s Demo\n", name, topic)

				switch num {
				case 1:
					myMessage()
					break
				case 2:
					returningvalue()
					break
				case 3:
					stud.name = "Tony Stark"
					stud.age = 24
					stud.dept = "Marketing"
					stud.contactnumber = 8888888888
					printStudent(stud)
					break
				case 4:
					array()
					break
				case 5:
					goroutine()
					break
				case 6:
					write()
					break
				default:
					fmt.Println("Enter a valid number to view Demo")
				}
				fmt.Println("Do you want to countinue Demo (Y/N)")
				var flag string
				fmt.Scan(&flag)
				if flag == "N" {
					check = 2
				}

			}
		} else {
			fmt.Printf("%s the passcode entered by you is incorrect\n", name)
		}
		fmt.Println("Demo Completed Successfully...!!!")
		fmt.Print("Total time duration of this Demo was ")
		fmt.Print(time.Since(start).Minutes())
		fmt.Print(" mins.")
	} else {
		fmt.Printf("Sorry %s !! You do not have access to this Demo", name)
	}
}

func myMessage() {
	fmt.Print("Hello World !! Welcome to Go Demo\n")
}

func vals() (int, int) {
	return 3, 7
}

func returningvalue() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func listoftopic() {

	fmt.Println("1 : Hello World")
	fmt.Println("2 : Functions")
	fmt.Println("3 : Structure")
	fmt.Println("4 : Arrays")
	fmt.Println("5 : Go Routine")
	fmt.Println("6 : Write File")
	//fmt.Println("\nTo Exit enter any other Number")

}
func printStudent(stud Student) {
	fmt.Println("Name: ", stud.name)
	fmt.Println("Age: ", stud.age)
	fmt.Println("Department: ", stud.dept)
	fmt.Println("Contact: ", stud.contactnumber)
}
func gettopicfromnum(num int) string {
	if num == 1 {
		return "Hello World"
	} else if num == 2 {
		return "Functions"
	} else if num == 3 {
		return "Structure"
	} else if num == 4 {
		return "Arrays"
	} else if num == 5 {
		return "Go Routine"
	} else if num == 6 {
		return "Write File"
	}
	return "a wrong input in"
}
func array() {

	var size int

	fmt.Print("Enter the Array Size = ")
	fmt.Scan(&size)

	numarray := make([]int, size)

	fmt.Print("Enter the Array Items = ")
	for i := 0; i < size; i++ {
		fmt.Scan(&numarray[i])
	}

	arrSum := 0

	for _, a := range numarray {
		arrSum = arrSum + a
	}
	fmt.Println("The Sum of Array Items = ", arrSum)
}
func goroutine() {
	go func() {
		fmt.Println("First")
	}()
	fmt.Println("Second")
	time.Sleep(2 * time.Second) // waiting time for 1 to complete
}
func write() {

	f, err := os.Create("demoSummary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Hi Alex...Demo was completed successfully\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Summary Generated")
}
