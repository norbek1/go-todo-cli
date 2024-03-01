package main

import "fmt"

func main() {
	tasks := []string{""}

	for {
		fmt.Println("\n\n1.Add Tasks")
		fmt.Println("\n2.Show tasks")
		fmt.Println("\n3.Exit")

		var option int
		fmt.Print("\nSelect option: ")
		fmt.Scanf("%d", &option)

		if option == 1 {
			var txt string
			fmt.Print("\nEnter a task: ")
			fmt.Scanf("%s", &txt)
			tasks = append(tasks, txt)

		} else if option == 2 {
			fmt.Println(tasks)
			
		} else if option == 3{
			break
		}
		
		
		
	}
	
}
