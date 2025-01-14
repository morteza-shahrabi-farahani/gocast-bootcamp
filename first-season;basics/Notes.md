# Session 1: go basics

## Compiler vs interpretable languages

The key differences between compiled and interpreted languages revolve around how they are executed, the tools they use for execution, and their performance characteristics. Here’s a breakdown:

Execution Process

Compiled Languages:

The source code is translated into machine code (binary executable) by a compiler before execution.
The resulting executable runs directly on the hardware.
Examples: C, C++, Rust, Go.

Interpreted Languages:

The source code is executed line-by-line or statement-by-statement by an interpreter at runtime.
No separate executable is created; the interpreter is required to run the code.
Examples: Python, Ruby, JavaScript (in traditional interpretation).

Performance

Compiled Languages:

Faster execution because the machine code is directly executed by the hardware.
Optimized by the compiler during the compilation process.

Interpreted Languages:

Slower execution because the interpreter needs to process each instruction at runtime.
Can have performance improvements using Just-In-Time (JIT) compilation (e.g., modern JavaScript engines like V8).

Development Workflow

Compiled Languages:

Require a compilation step before running the program.
Errors are caught at compile time, which can help prevent runtime issues.

Interpreted Languages:

Can be run immediately without a separate compilation step, which can speed up development and debugging.
Errors appear at runtime, which might make debugging slower for larger programs.

Portability

Compiled Languages:

Platform-dependent: The compiled binary works only on the platform it was built for (unless cross-compiled).
The source code needs to be recompiled for different platforms.

Interpreted Languages:

Platform-independent: As long as the interpreter is available, the same source code can run on any platform.

Error Detection

Compiled Languages:

Catch syntax and some logical errors during compilation (before running the program).

Interpreted Languages:

Errors are detected during execution, making it less predictable when issues will arise.

Examples of Dual-Behavior Languages

Some languages blur the line:

Java: Compiled to bytecode and then interpreted or JIT-compiled by the JVM.

Python: Interpreted but can be compiled to bytecode (.pyc) for faster startup.

JavaScript: Traditionally interpreted but often JIT-compiled in modern engines.


Why Go is Platform-Independent:

Cross-Compilation Support:

Go's compiler supports cross-compilation out of the box. You can compile code for different platforms and architectures by setting environment variables like GOOS (Operating System) and GOARCH (Architecture).

Static Binaries:

Go often produces statically linked binaries, meaning the binary includes all necessary dependencies (except for OS-level libraries like glibc in some cases).
This makes it easy to run the binary on the target platform without worrying about missing dependencies.

Standard Library:

The Go standard library is designed to be consistent across platforms, ensuring the code behaves the same way regardless of the target environment.

Using Go Modules for Platform-Independent Code

Go Modules help manage dependencies and ensure that cross-platform compatibility is maintained in source code. For example:

## Implicit and Explicit type conversions

Implicit type conversion, also known as type coercion, occurs automatically without explicit instructions from the programmer.

Go does not support implicit type conversion between different types. This design decision prevents unintended errors that can occur when types are automatically coerced (e.g., converting a float to an int without realizing it).
For example, you cannot directly assign a float64 to an int without explicitly converting it.

```
var x int = 10
var y float64 = 20.5
x = y // This will throw an error
```

Explicit type conversion requires the programmer to explicitly specify the type they want to convert a value to. 
Go uses explicit type conversion to promote clarity and safety.

```
var x int = 10
var y float64 = 20.5
x = int(y) // Explicit conversion
fmt.Println(x) // Outputs 20
```

Implicit conversions (type coercion) are common in other languages.

## Slice and array

\* One of the features of arrays is there size. The type of two array variables which have different sizes are not equal and we cannot assign them.


\* At the time of initializing slices, go compiler assign capacity and length of the slice equally. For example if we make a slice with length of 5, the compiler automatically set the capacity 5.
But if you append an element to the slice and overcome its length, the compiler first 
make the capacity of the slice twice the last value and then add element to the slice.

## Editional notes
\* In each folder, we must have just one package and package name.

\* Reserved keywords in go are just 26 words which is very small.

