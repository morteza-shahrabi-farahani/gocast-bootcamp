\* The init function excecuted without calling. If you need some codes to be run without calling the function, you can write it in init function. The init function, excecuted before the main function. The order of executing the init function, depends on the order of initializing the packages. 

\* In tests, using fatalf is better than using errorf

\* We can use init function for things like connecting to the database or oading the source file or ...

\* Name of the packages should be singular instead of plural.

\* Entity is different from model. Model is what we use in database. But entity is the thing we use and we have in our business logic. For example maybe we don't need created at or updated at in the business logic, but obviously we need them in database.

\* The method receiver name can be small word or abbreviation words. Because the scope of using the variable is small and we don't use that variable on external files.

\* If we want to defer a function which has return value and error and we should check it for safety, we can use defer function in which that function is called and the returned variables of the function is checked.

```
defer file.Close()

///////

defer func() {
    err := file.Close()
    if err != nil {
        fmt.Println("error")
    }
}()
```

\* We can use constructor for the structs which needs to have private fields for initializing

\* If it is possible, it is better to use methods instead of functions.

\* dependency injection: It means that we depens package to each other. For example we say that handler package need a variable from storage package so we use dependency injection between them. We can define a type for handler package and set a field for storage init. And for example we set this variable in constructor part of the program. The alternative way for handing this functionality is using global variables which is not good enough. Moreover, in handler package, we cannot change the fields of repository variable and it is not good practice. All the changes for the repository, must be done at the repository package and we just use repo variable in handler package.