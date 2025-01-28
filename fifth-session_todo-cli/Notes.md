\* Instead of using fmt.Scanf() we can use bufio.NewScanner() and Scan() method for getting input from the terminal.

\* A zero value struct is simply a struct variable where each key’s value is set to their respective zero value. If we want to check if struct is nil or not, actually we cannot and we should check its fields. However, we can declare a pointer to a struct and in this scenario, because the zero value of pointer is nil, we can check whether a variable of struct initialized before or not.

\* Code sanitization: Sanitizing will remove any illegal character from the data.
Validating will determine if the data is in proper form. We must do sanitization 
in every part that is needed. So for each part of the program, we must sure that 
the input is completely valid and the required validations for the input in the
last parts are done before.

\* variable shadowing: You can create multiple variables with the same name
in different scopes of the program.

\* Serialization is the process of converting a data object—a combination of
code and data represented within a region of data storage—into a series of
bytes that saves the state of the object in an easily transmittable form. 
In this serialized form, the data can be delivered to another data store 
(such as an in-memory computing platform), application, or some other 
destination.

\* When storing user information to the database, it is better to hash the passwords
and then storing them to the database. So if some one accessed to the database file,
he or she cannot find the users' password. hashing functions are one-way functions.
This means that by having the password string, we can easily reach the hashed 
password, but by having the result hashed password string, we can't reach to the
original password of the user.

\* Encoding: Converting a stream of bytes to another format of steam of bytes.
