// package main

// import "fmt"


// func celsiusToFahrenheit (celsius int) float64{
// 	result := float64(celsius)*9/5 + 32
// 	return float64(result)
// }

// func checkTrafficLight(color string) string{
// 	switch color {
// 	case "red":
// 		return "থামুন! সামনে যাওয়া নিষেধ।"
// 	case "yellow":
// 		return "প্রস্তুত হোন, কিছুক্ষণের মধ্যে লাইট বদলাবে।"
// 	case "green":
// 		return "নিরাপদে সামনে এগিয়ে যান"
// 	default:
// 		return "ভুল ট্রাফিক সিগন্যাল!" 
// 	}
// }

// func evaluateResult(mark int) string {
// 	if mark >80{
// 		return "you got A+"
// 	}else if mark >= 70 && mark <= 80{
// 		return "you got a"
// 	}else if mark >= 60 && mark < 70{
// 		return "you got b"
// 	}else if mark >= 40 && mark < 60{
// 		return "you got c"
// 	}else if mark >= 33 && mark < 40{
// 		return "you got d"
// 	}else{
// 		return "you failed"
// 	}
// }
// func main() {
// 	fmt.Println(celsiusToFahrenheit(105))
// 	fmt.Println(checkTrafficLight("red"))
// 	fmt.Println(evaluateResult(85))
// 	printSomething()
// 	hello("Alice")
// }

// func printSomething(){
// 	fmt.Println("education must be free for everyone")
// }

// func hello( name string){
// fmt.Println("hello", name)
// }






// package main
// import "fmt"

// var target = 10

// func main() {
//     if true {
//         target := 20
//         _ = target 
//     }
//     fmt.Println(target)
// }


// package main

// import "fmt"


// func main() {
// 	// স্টেপ ১: ফার্স্ট ওয়েলকাম মেসেজ
// 	fmt.Println("=== স্বাগতম রকিব ভাইর গো ল্যাঙ্গুয়েজ প্রজেক্টে! ===")

// 	// স্টেপ ২: নাম এন্টার করতে বলবে
// 	var name string
// 	fmt.Print("আপনার নাম এন্টার করুন: ")
// 	fmt.Scanln(&name) // ইউজারের ইনপুট ভ্যারিয়েবলে স্টোর করবে

// 	// স্টেপ ৩: এন্টার দিলে নাম দেখাবে
// 	fmt.Printf("হ্যালো %s, চলুন এবার একটা যোগ অংক করি!\n", name)
// 	fmt.Println("----------------------------------------")

// 	// স্টেপ ৪: ২টা নাম্বার দিতে বলবে (Number 1 and Number 2)
// 	var num1, num2 int

// 	fmt.Print("প্রথম নাম্বারটি (Number 1) দিন: ")
// 	fmt.Scanln(&num1)

// 	fmt.Print("দ্বিতীয় নাম্বারটি (Number 2) দিন: ")
// 	fmt.Scanln(&num2)

// 	// স্টেপ ৫: ২টার যোগফল দেখাবে
// 	result := num1 + num2
// 	fmt.Printf("ফলাফল: %d এবং %d এর যোগফল হলো = %d\n", num1, num2, result)
// 	fmt.Println("========================================")
// }