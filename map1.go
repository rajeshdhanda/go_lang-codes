package main
import "fmt"
func main() {
	elements := map[string]map[string]string{        
		"H": map[string]string{            
			"name":"Hydrogen",            
			"state":"gas",        
			},        
		"He": map[string]string{            
			"name":"Helium",            
			"state":"gas",        
			},        
		"Li": map[string]string{            
			"name":"Lithium",            
			"state":"solid",        
			},        
		"Be": map[string]string{            
			"name":"Beryllium",            
			"state":"solid",        
			},        
		"B":  map[string]string{            
			"name":"Boron",            
			"state":"solid",        
			},        
		"C":  map[string]string{            
			"name":"Carbon",            
			"state":"solid",        
			},        
		"N":  map[string]string{
			"name": "Nitrogen",
			"state": "gas",
			},
	}
	fmt.Println(elements)
}