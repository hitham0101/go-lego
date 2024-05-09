package main

// func main() {

// 	fmt.Println("hey man")

// 	var a, b = 1, 2
// 	var c , d = 3 ,4

// 	if a < b {
// 		fmt.Println("A is bigger man")
// 	}
// 	if d < c {
// 		fmt.Println("d is bigger man")
// 	}else{

// 		fmt.Println("you are wrong")
// 	}

// 	if d> a || c <b {

// 		fmt.Println("Done")
// 	}

// 	if d<a{
// 		fmt.Println("d smaller than c")

// 	}else if b>a{

// 		fmt.Println("b bigger than a")
// 	}else{
// 		fmt.Println("this is the default")
// 	}

// }

// func main() {

// 		day := "sunday"

// 	switch day {
// 	case "friday":
// 		fmt.Println("ohhhhh friday")
// 	case "saterday":
// 		fmt.Println("ohhhhh satalana")
// 	case "sunday":
// 		fmt.Println("outttting")
// 	default:
// 		fmt.Println("DON'T KNOW THAT")

// 	}

// }

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "channel 1"
// 	}()

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch2 <- "channel 2"
// 	}()

// 	for i := 0; i < 2; i++ {

// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("Received from", msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println("Received from", msg2)
// 		}
// 	}
// }

// func main() {

// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 	}()
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c2 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)
// 		}
// 	}
// }

// func doWork(ctx context.Context, mssg string) {

// 	time.Sleep(2 * time.Second)
// 	select {
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Operation completed within the timeout")
// 	case <-ctx.Done():
// 		fmt.Println("Operation timed out")
// 	}

// }

// func main() {

// 	deadline := time.Now().Add(5 * time.Second)

// 	fmt.Println(deadline)

// 	fmt.Println("wasssup")

// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	doWork(ctx, "MEOW")

// }

// func main() {

// 	ctx := context.WithValue(context.Background(), "requestID", "12345")

// 	v := ctx.Value("requestID")
// 	fmt.Println(v)
// }
