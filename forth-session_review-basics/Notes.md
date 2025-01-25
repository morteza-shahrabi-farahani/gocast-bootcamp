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
