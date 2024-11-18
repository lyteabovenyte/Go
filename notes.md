#### notes while reading top books in golnag:

##### Go programming language book.
###### chapter 1.
- **Type safety** in the source code is a programming language control that ensures that any variable access only its authorized memory locations in a well-defined and permissible way. In other words, the type safety feature ensures that the code doesn't perform any invalid operation on the underlying object.
- In CSP , a program is a parallel composition of processes that have *no shared state*; the processes communicate and synchronize using **channels**
- System calls are low-level functions provided by the operating system for accessing resources such as files, networks, and other hardware devices. Go provides a convenient way to use system calls by providing a package called “syscall” which provides an **interface** to call system calls. (system call interface)[more](https://byteshiva.medium.com/using-system-calls-in-go-for-low-level-control-part-1-df390f8c268e)
- golang lacks on constructor and destructor, (oop definition lacks some how)
- the `for` loop is the only loop statement in golang.
- `strings.Join` uses builder pattern under the hood for performance-wise issues for concatenation.
- 