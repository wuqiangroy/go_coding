package main

import (
	"fmt"
	//"runtime"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OS X.\n")
	case "linux":
		fmt.Print("Linux.\n")
	case "windows":
		fmt.Print("Windows.\n")
	default:
		fmt.Printf("%s.", os)
	}
	i := 2
		fmt.Printf("%v as: ", i)
	switch i {
	case 1:
		fmt.Print("No\n")
	case 2:
		fmt.Print("Yes\n")
	default:
		fmt.Print("Yes or No\n")
	}
	fmt.Print("DEBUG##########", time.Now().Weekday().String())
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Printf("today is %T, It's weekend\n", day.String())
	default:
		fmt.Printf("today is %T, it's weekday\n", day.String())
	}

	t := time.Now()
	//没有条件的switch相当于switcch true
	switch {
	case t.Hour() < 12:
		fmt.Printf("Now is %T, It's before noon\n", t.String())
	case t.Hour() > 12:
		fmt.Printf("Now is %T, It's after noon\n", t.String())
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Print("I am a bool\n")
		case int:
			fmt.Print("I am a int\n")
		default:
			fmt.Printf("Don't know type: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("this is my dream")

	fmt.Print("When's Saturday?\n")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Print("today")
	case today + 1:
		fmt.Print("tomorrow")
	case today + 2:
		fmt.Print("In 2 days")
	default:
		fmt.Print("Too far to welcome saturday")
	}
}
