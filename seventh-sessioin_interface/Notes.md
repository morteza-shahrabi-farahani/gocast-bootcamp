## interface

Why we use interfaces? By using interfaces, we can prevent main components from becoming dependent on subcomponents. We call this situation, dependency inversion.  

```
type App struct {
    Name string 
    StorageFilePath string
    InMemoryStorage storage.Memory
}

// it become

type App struct {
    Name string
    UserStorage UserStore
}

type UserStore interface {
    CreateUser(u user.User)
    ListUser() []user.User
}
```

In the previous way, the App module is related to the storage package. But in this new way, the storage layer is dependent to the App module which is better. So we change from dependency injection, to dependency inversion!

In this new way, if the storage field of the app changed, the code of the App would not change at all. The cost of change in our system, become very low and the change in the app layer become easy. We must just change the UserStorage variable and we do not need to change anything else.

\* interface is like an contract. We specify the behavior and the methods which we expect from an entity init.

\* Unlike other programming languages, we don't specifically say that each type implements which interfaces. However, if the type has all the required methods of the interface, it will automatically can be used as a variable of that interface. It is a good capability of go over other programming languages. Consider a situation, in which we define an interface and one of the types of the public packages can be used as variable of this interface. But if we have this scenario in other languages, it become harder to handle this situation because in the implementation of that public type or method, we must explicitly mention that type implements that interface. We call this feature in go, composition over inheritance.

\* Interface helps us to compose different functionalities and behaviors. Each interface just focouses on implementing each methods. But by composing different interfaces, we can achieve new behaviors and functionalities.

\* When we use a variable of interface in a type, we should just use the functionality of that interface from that variable. We just call its methods.

\* We don't have inheritance in go.

\* For printing in the desired format, we can implement stringer interface.

\* error is an interface in golang and not type. Every type which implement ```Error() string``` method, become type of error. One of the benefits of being interface, is that there zero value is nil. So if we don't have error, we can return nil. 

## type assertion 

```
typeValue := tempString.(tempType)
```