package main 

import "fmt"

func main()  {
	s1 := []int{1, 2, 3}

	//[0,2)
	s2 := s1[0:2]

	fmt.Printf("len = %d, s1 = %v \n", len(s1), s1)

	fmt.Printf("len = %d, s2 = %v \n", len(s2), s2)

	s2[0] = 100

	fmt.Printf("len = %d, s1 = %v \n", len(s1), s1)

	fmt.Printf("len = %d, s2 = %v \n", len(s2), s2)

	//深拷贝
	s3 := make([]int, 3)
	copy(s3, s2)
	s3[0] = 101
	fmt.Printf("len = %d, s3 = %v \n", len(s3), s3)



	fmt.Printf("len = %d, s1 = %v \n", len(s1), s1)

	fmt.Printf("len = %d, s2 = %v \n", len(s2), s2)

}