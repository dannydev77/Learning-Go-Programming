# Variables
- Redeclarartion and shadowing 
- Visibility
- Naming Conventions
- Type Conversions

<!-- method 1 -->
var a int
	a = 45
	a = 56
- This format can be used when we declare a variable but not ready to use it yet
<!-- method 2 -->
var a int = 55
- if Go doesn't have enough info to assign the type that you actually want...susceptible to change
```
var a float32 = 55
fmt.Printf("%v, %T", a,a)
```
- To have more control over the type since go can't easily infer float32 


<!-- Method 3 -->
a := 55

Go will infer automactically
examples: 
    - a := 44 -> int
    - a := 44. -> float64

<!-- method 4 -->
```

// Declaring at the package level
// you have to use the full declaration syntax

var a int = 42

func main() {
	fmt.Println("===================================")
	fmt.Println(a)
	fmt.Println("===================================")
}
```
## shadowing
a variable at the inner scope takes precedence more... called shadowing
```
var a int = 42

func main() {
    var a int = 77
	fmt.Println("===================================")
	fmt.Println(a)
	fmt.Println("===================================")
}
output: 77
```

# Variables naming Conventions
#### at the package level
1. if at the package level and is: 
- lower case variables are scoped to the package: any file in the package can access it
- Upper case variables are exported to other packages and is globally visible
- at block scope: only visible to the block say in a function block
2. Rules of naming
- length of the variable name should reflect its life
- used for a long time, use long names e.g seasonNumber
3. Acronyms variables
- keep them all upper case 
var theHTTP string = "https://google.com"