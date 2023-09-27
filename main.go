package main

func main() {
	server := newAPIserver(":3000")
	server.Run
	//fmt.Println("Yeah Cedric!")
}
