package main

import "fmt"

//for -> Only construct in go for looping
func main(){


	//while loop
	i:=1
	for(i<=3){
		fmt.Println(i)
		i =i+1;
	}

	fmt.Println("**********************Classic For loop**********")
	//classic for loop
	for j:=0; j<3;j++ {
		fmt.Println(j)
		
	}

	fmt.Println("**********************Classic For loop break**********")
	//classic for loop
	for j:=0; j<3;j++ {
		
		if(j==1){
			break
		}
		fmt.Println(j)
		
	}

	fmt.Println("**********************Classic For loop continue **********")
	//classic for loop
	for j:=0; j<3;j++ {
		
		if(j==1){
			continue
		}
			
		fmt.Println(j)
	}

	fmt.Println("**********************Classic For loop range **********")
	for j:=range 10 {
					
		fmt.Println(j)
	}

}