\* The overhead time for the cpu for context switching between multiple threads of the program, is more time consuming than the required time for go scheduler for context switching between multiple goroutines.

\* Main function is also a goroutine. We must do something to avoid terminating the main function while we want to execute some other goroutines. If the main function terminates, the program terminates and the goroutines of the program would not execute anymore.

\* In concurent programming, the important point is that while coding, we should pay attention to write the code which does not rely on the order of execution of the goroutines. With any random order that the program have, the result and the output must be the same. The program and code must be permutation invariant :))

Notice that the last statement is different from doing something and force the program and order of execution of the goroutines. This is not good. We should try to make our program permutation invariant not force the order of executing the goroutines. This order is not setted by us in most important and real scenarios. We don't have any authority over this.

\* We should avoid having a shared variable between goroutines using defining global variables.

"Do not communicate by sharing memory; instead, share memory by communicating."

\* Map in go is not concurrently safe.

\* If we just have receiver for the channel and don't have any sender for that channel, deadlock will happened.

\* After that we sure we won't send any value to a channel, we should close it so that we won't get deadlock anymore in receiver part. However, if we try to write on a close channel, we will get panic.