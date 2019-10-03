package main

func main()  {
	t := triangle{}
	t.height = 10
	t.base = 5

	sq := square{}
	sq.sideLength = 7
	printArea(t)
	printArea(sq)
}

