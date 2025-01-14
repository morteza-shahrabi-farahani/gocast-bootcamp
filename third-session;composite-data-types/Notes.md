## Map

\* Zero value of map is nil.

\* For finding if a key exists in the map or not

```
value, ok := map['test']
```

Map is also stored like slice and there is a pointer which maps to the first element of an array.

## Functions

variadic functions are functions with infinity values.

```
func prt(values ...string) {
    for _, value := range values {
        fmt.Println(value)
    }
}
```

## Methods

By using structs, we can model the real world and real entities in our code. And we use methods when some functionalities are related to that type of entity. After defining methods for that type, by calling one variable of that type, we exactly know what functionalities are related to that type. Somehow we can say by using methods we assign behavior to our types (like object oriented programming langugages). 

Consider a scenario where we want to add new behavior to some internal types of go (like int, int32, float32, ...). We cannot add new method for that type because the methods for each type must be located in the package which the type is initialized. But by making new type which is alias for that type, we cann add new method to this alias type.

## Pointers

For optimizing usage of memory, it is better to pass pointers as the function parameters.

Zero value of pointers are also nil.

```
var p *int
i := 43

p = &i
fmt.Println("p", p)
fmt.Println("value that p references to: ", *p)

// p 0xc0000220c8
// value that p references to: 43

d := 44

p = &d

fmt.Println("p", p)
fmt.Println("value that p references to: ", *p)

// p 0xc00010c018
// value that p references to: 44

*p = i
fmt.Println("p", p)
fmt.Println("value that p references to: ", *p)

// p 0xc00010c018
// value that p references to: 43

// in the last part, the address that p is referencing is not changed. However, the value that is stored on that location is changed from 44 to 43 (in this scenario, the actual value of d also changed!)
```

Passing pointer as parameter to functions and methods when struggling with structs and ... is good. However, if we don't have composite data types and we have some sort of types like int, string, ... as far as code readability is concerned, it is better to pass value as function variable. So, passing pointers as parameters of functions is not good in all scenarios.

when we pass variable as value to the functions, the compiler makes a new variable in the memory and save the value of that variable to this new location (copy the variable) and then send this new variable to the function. So every changes that are applied to the input parameter of the function, would not change anything on the original variable.

## Slice 
If we assign a slice variable to a new variable, it assigns the pointer. So if you
change the new variable, the original slice is also changed. If you want this situation
not happened, you should use copy keyword.
