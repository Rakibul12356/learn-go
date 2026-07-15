package main

import "fmt"


func celsiusToFahrenheit (celsius int) float64{
	result := float64(celsius)*9/5 + 32
	return float64(result)
}

func checkTrafficLight(color string) string{
	switch color {
	case "red":
		return "থামুন! সামনে যাওয়া নিষেধ।"
	case "yellow":
		return "প্রস্তুত হোন, কিছুক্ষণের মধ্যে লাইট বদলাবে।"
	case "green":
		return "নিরাপদে সামনে এগিয়ে যান"
	default:
		return "ভুল ট্রাফিক সিগন্যাল!" 
	}
}

func evaluateResult(mark int) string {
	if mark >80{
		return "you got A+"
	}else if mark >= 70 && mark <= 80{
		return "you got a"
	}else if mark >= 60 && mark < 70{
		return "you got b"
	}else if mark >= 40 && mark < 60{
		return "you got c"
	}else if mark >= 33 && mark < 40{
		return "you got d"
	}else{
		return "you failed"
	}
}
func main() {
	fmt.Println(celsiusToFahrenheit(105))
	fmt.Println(checkTrafficLight("red"))
	fmt.Println(evaluateResult(85))
}