\* if the requested key is not existed in the map, it will return zero value of the type of the value for result

\* In go, we have the extend abbility. For example if you have a type "A" which in that type another type "B" is one of its fields, if type "B" implements some methods, you have access to type "B" methods in variables of type "A".

\* Its better to place utility functions like parsing, serializing, marshalling and ... in a separate package. It can be named utils or pkg or anything else. Placing them in the enitty package is not good practice. Think to this in this way that the responsibility or the functionality of that entity is not marshalling and unmarshalling.

\* One of the benefits of go is that it is cross-platform and we code on mac but build an executable file for windows. You can use "go tool dist list" to see the standard platform that go can be built on them. In normal build, we build the application using "go build main.go". If you want to build for other os, you can use this command "GOOS=android GOARCH=amd64 go build main.go"

\* Pay attention to encapsolation for struct entity fields.
