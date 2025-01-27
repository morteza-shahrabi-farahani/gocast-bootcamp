\* In other languages, character is one byte. But is Go, as far as supporting all langugaes is concerned, rune is made and it is 4 bytes instead of one byte.

\* The result of 
```
fmt.Println(len("سلام"))
```
is 8! Because each persian character stored in two bytes not one byte. Because its utf-8 int for it is two bytes not one byte. 

\* ASCII is just English words but utf-8 is all languages.

\* Although the "سلام" length is 8, if you use for loop for it, you will get 4 values. Because it has 4 runes but it is 8 bytes. So the len, outputs the length of the bytes of the string, but when using for loop, it just give me the runes of the string not its bytes.

```
for index, value := range tempStr {
    fmt.Println(index, value)
}
```

\* When initializing an integer, if we don't specify its size, for example int8 or int32 or ... 
the compiler assign int type to it. And the size of int variable, depends on the cpu of the system 
which the program is running on it. If it is 32 bytes system, it would be int32 and otherwise, it
would be int64

\* use go mod init {name_of_the_package} to create go.mod for your package.

\* Every go program, need a main package and a main function for running

\* Dgraph => graph database

\* GOROOT is an environment variable that defines the location of your Go SDK. 
It means you can find the compiler, go tools or, standard libraries in this directory.
GOPATH is another environment variable. It defines the locations of your go source codes.
If this variable includes more than one location, these locations are separated by columns (semicolon separated for Windows).
When your code has any imported package, this package has to be located here.

\* Thanks to Go modules, which is a new method to manage dependencies, 
you don't have to use $GOPATH/src for your projects since Go 1.11.
However, pkg and bin folders are still used.

\* zero value of rune is 0

\* You can use just return for returning values of the function in go. But as far as code readability is concerned, it is better to write their names when returning variables.

\* In go we don't have try catch and throw exception. Instead, where ever we have error, we should return it from the function.

\* If we specify the length of the slice, it become array not slice. So later we cannot change its size and its size is fixed. If you can choose both array and a slice for a purpose, it is better to use array because it is safer.