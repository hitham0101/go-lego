package main

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	go performTask(ctx)

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Task timed out")
// 	}
// }

// func performTask(ctx context.Context) {
// 	select {
// 	case <-time.After(5 * time.Second):
// 		fmt.Println("Task completed successfully")
// 	}
// }

// var weeg sync.WaitGroup

// func main() {
// 	weeg.Add(1)
// 	ctx := context.Background()

// 	ctx = context.WithValue(ctx, "UserID", 123)

// 	go performTask(ctx)

// 	// Continue with other operations
// 	weeg.Wait()
// }

// func performTask(ctx context.Context) {
// 	userID := ctx.Value("UserID")
// 	fmt.Println("User ID:", userID)
// 	weeg.Done()
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go performTask(ctx)

// 	time.Sleep(2 * time.Second)
// 	cancel()

// 	time.Sleep(1 * time.Second)
// }

// func performTask(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Task cancelled")
// 			return
// 		default:
// 			// Perform task operation
// 			fmt.Println("Performing task...")
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }
