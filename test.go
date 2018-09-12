package main

import (
	"fmt"
	"github.com/o1egl/govatar"
)

func main() {
	
	getgirl := govatar.GenerateFile(govatar.FEMALE, "girlavatar.jpg")
	getboy := govatar.GenerateFile(govatar.MALE, "boyavatar.jpg")
	
	//get one of each avatar
	fmt.Println(getboy,getgirl)
	fmt.Println("Enjoy your avatars.")
	
}
