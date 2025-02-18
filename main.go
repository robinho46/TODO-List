package main

import ()

func main() {
	connectDB()
	defer db.Close()
	menu()
}
