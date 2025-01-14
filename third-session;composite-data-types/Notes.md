## map

\* Zero value of map is nil.

\* For finding if a key exists in the map or not

```
value, ok := map['test']
```

## functions

variadic functions are functions with infinity values.

```
func prt(values ...string) {
    for _, value := range values {
        fmt.Println(value)
    }
}
```

## methods

By using structs, we can model the real world and real entities in our code. And we use methods when some functionalities are related to that type of entity. After defining methods for that type, by calling one variable of that type, we exactly know what functionalities are related to that type. Somehow we can say by using methods we assign behavior to our types (like object oriented programming langugages).