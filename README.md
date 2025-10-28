# Welcome to Go Quickstart

This repository is for my CS-330 class at Simmons University where my goal is to create a simple straightfoward tutorial for getting started using Go for students already comfortable with at least one programming language. This is a work-in-progress so there will be some errors, omissions and incomplete sections. When the project is complete and the semester is over, I will tag it **complete** and the repo will allow pull requests. I'm currently working in Windows 11, and while I try to give instructions for Mac and Linux, I can't test them all on my machine and so I'm mainly going by documentation sources. Your miles may vary. Until then, if you have any issues or commentary, please see the [Discussion](https://github.com/omicreativedev/go-quickstart/discussions) section of the repository.

If you need to reach me, message me on [LinkedIn](https://www.linkedin.com/in/sshunamon/)

Credits:

Side by Side Markdown Editor: [Dillinger.io](https://dillinger.io/)

Hosting: [Render.com](https://render.com/) _Free Tier_

------------------------
# Table of Contents

### Part 1

[History of the Go Programming Language](#history-of-the-go-programming-language)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Important Features of Go](#important-features-of-go)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[What is Go good for?](#what-is-go-good-for)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Famous Projects Created with Go](#famous-projects-created-with-go)<br>
<br>
[Installation and Setup](#installation-and-setup)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Installing Go](#installing-go)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Choosing an IDE](#choosing-an-ide)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Installing VSCode Go Extension](#installing-vscode-go-extension)<br>
<br>
[Creating Your First Go Project in VSCode](#creating-your-first-go-project-in-vscode)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Hello World!](#creating-your-first-go-project-in-vscode)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Commenting Your Code](#commenting-your-code)<br>
<br>
[Go Learning Resources](#go-learning-resources)<br>

### Part 2

[Type System and Variable Semantics](#type-system-and-variable-semantics)<br>
[Data Types](#data-types)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Basic Types](#basic-types)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Boolean](#boolean)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Numeric](#numeric)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Strings](#strings)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Aliases](#aliases)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Composite Types](#composite-types)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Aggregate](#aggregate): Arrays, Structs<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Reference](#reference): Slices, Maps, Channel, Pointer<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Interface Types](#interface-types)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Complex Types](#complex-types)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Type Conversion](#type-conversion)<br>
<br>
[Syntax](#syntax)<br>
[Reserved Words](#reserved-words)<br>
[Variable Naming Requirements & Conventions](#variable-naming-requirements--conventions)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Required by Compiler](#required-by-compiler)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Encouraged by Professional and Community Standards](#encouraged-by-professional-and-community-standards)<br>
[Composite Literals](#composite-literals)<br>
[Type Aliases vs. Defined Types](#type-aliases-vs-defined-types)<br>
[Operators](#operators)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Availability](#availability)<br>
[Binding](#binding)<br>
[Storage, Addresses and Lifetime](#storage-addresses-and-lifetime)<br>
[Scope](#scope)<br>
<br>
[Limitations](#limitations)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Pitfalls](#pitfalls)

Part 3
[Start Here](https://github.com/omicreativedev/go-quickstart/tree/main#control-statements-loops-selection-conditionals)

Part 4
[Function Declaration & Syntax](https://github.com/omicreativedev/go-quickstart/tree/main#function-declaration-syntax)
[Function Scope](https://github.com/omicreativedev/go-quickstart/tree/main#function-scope)
[Function Passby](https://github.com/omicreativedev/go-quickstart/tree/main#pass-by)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Side Effects](https://github.com/omicreativedev/go-quickstart/tree/main#side-effects)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Guardrails](https://github.com/omicreativedev/go-quickstart/tree/main#guardrails)
[Memory Management](https://github.com/omicreativedev/go-quickstart/tree/main#memory-management-stack-vs-heap)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[Recursive Functions](https://github.com/omicreativedev/go-quickstart/tree/main#recursive-functions)


------------------------
![Part 1](https://github.com/omicreativedev/go-quickstart/blob/main/images/part_1.png?raw=true "Part 1")

# History of the Go Programming Language

The Go programming language, often referred to as Golang, was created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson as a direct response to the frustrations experienced in software development within the company.[^1] The main catalyst for its creation was the difficult nature of using existing languages for massive systems work, lengthy compilation times for languages like C and the complexity of distributed systems built in languages like Java.[^2] The designers were also motivated by a desire to improve the scale of development for large teams of programmers working on shared codebases.[^3] Officially announced in 2009 as an open-source project, Go was designed to be a simple language that emphasized ease of use for modern hardware but with the power of other languages already in use at the company.[^4]

If you want an interesting read about the beginnings of Go, check out the official [Go Spec Blob on Github](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) by Robert Griesemer himself.

[^1]: [Go FAQ: History](https://go.dev/doc/faq#history)
[^2]: [Go FAQ: Creating a New Language](https://go.dev/doc/faq#creating_a_new_language)
[^3]: [The Go Programming Language Specification (Preface)](https://go.dev/ref/spec)
[^4]: [Official Go Announcement](https://opensource.googleblog.com/2009/11/hey-ho-lets-go.html)

### Important Features of Go

<!--
> [!NOTE]
> Unfinished
-->

- Simplicity and Minimalism
  <br>There are no classes or inheritance[^5], exceptions[^6], or higher order functions[^7] in Go. These are achieved with workarounds.
- Built-in Concurrency Primitives: Goroutines and Channels[^8]
- Fast Compilation to a Single, Static Binary[^9]
- Explicit Error Handling[^10]
- Composition over Inheritance[^11]
- Built-in Tooling[^12]
- Garbage Collection[^13]
- Generics (Recent Addition)[^14]

A full tour of Go by Russ Cox is available on YouTube [here](https://www.youtube.com/watch?v=ytEkHepK08c&t=620s).

[^5]: [Golang #3: No Classes, No Inheritance — Just Structs & Methods, and It Works](https://medium.com/@wedevare/golang-3-no-classes-no-inheritance-just-structs-methods-and-it-works-4cd4c8cb45be)
[^6]: [Why doesn’t Go have Exceptions?](https://blog.cubed.run/why-doesnt-go-have-exceptions-b9c16165d93d?gi=f43007b88b72)
[^7]: [Generic Map, Filter and Reduce in Go](https://itnext.io/generic-map-filter-and-reduce-in-go-3845781a591c)
[^8]: [Go Blog: Share Memory by Communicating](https://go.dev/blog/codelab-share)
[^9]: [On Golang, “static” binaries, cross-compiling and plugins](https://medium.com/@diogok/on-golang-static-binaries-cross-compiling-and-plugins-1aed33499671)
[^10]: [Mastering Error Handling in Go: A Guide to Explicit Error Handling](https://medium.com/@sanjay.krishna.m.27/mastering-error-handling-in-go-a-guide-to-explicit-error-handling-072eef6cf00d)
[^11]: [Why did Golang Choose Composition as it’s base rather than Inheritance?](https://medium.com/deep-golang/why-golang-choose-composition-as-its-base-rather-than-inheritance-1225d22a4798)
[^12]: [The How and Why of Go, Part 1: Tooling](https://okigiveup.net/tutorials/the-how-and-why-of-go-part-1-tooling/)
[^13]: [A Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)
[^14]: [Tutorial: Getting Started with Generics](https://go.dev/doc/tutorial/generics)

### What is Go good for?

- Cloud-Native & Distributed Systems (Microservices)[^15]
- Command-Line Interfaces (CLIs)[^16] & DevOps Tools[^17]
- Web Servers & API Backends[^18]
- Concurrent Network Services[^19]
- Data Processing & Pipelines[^20]
- Databases & Storage Systems[^21]
- Cryptography[^22] & Security Tools[^23]
- Embedded Systems[^24] & IoT[^25]
- Scripting[^26] & Automation[^27]
- Proxy[^28] and Load Balancer Infrastructure[^29]

[^15]: [Go for Cloud and Networked Services](https://go.dev/solutions/cloud)
[^16]: [Command Line Interfaces in Go](https://go.dev/solutions/clis)
[^17]: [Development Operations & Site Reliability Engineering](https://go.dev/solutions/devops)
[^18]: [Building Restful API Go Tutorial](https://codezup.com/building-restful-api-go-tutorial/)
[^19]: [Go Concurancy and Parallelism](https://www.geeksforgeeks.org/go-language/go-concurrency-and-parallelism/)
[^20]: [Go Pipeline Patterns](https://www.golang101.com/advanced-concurrency-patterns/pipeline-patterns/)
[^21]: [Databases Implemented in Go](https://awesome-go.com/databases-implemented-in-go/)
[^22]: [Cyptography in Go](https://www.golang101.com/go-standard-library/cryptography/)
[^23]: [10 Programming Languages for Cybersecurity](https://www.brillicaservices.com/blogs/cyber-security-programming-languages)
[^24]: [Go Programming for Embedded Systems](https://rathorji.in/p/Go_Programming_for_Embedded_Systems)
[^25]: [Microcontroller Development with Golang](https://medium.com/@alrazak/microcontroller-development-with-go-golang-d947f2141ebb)
[^26]: [Embeddable Scripting Languages in Go](https://awesome-go.com/embeddable-scripting-languages/)
[^27]: [20 Libraries for Automation in Golang](https://blog.devgenius.io/20-libraries-for-automation-with-golang-4935e13fac9c)
[^28]: [Building a Proxy Server with Golang](https://www.privateproxyguide.com/building-a-proxy-server-with-go-golang-for-high-performance/)
[^29]: [Built My Own Load Balancer in Go](https://medium.com/@yashbatra11111/built-my-own-load-balancer-in-go-and-benchmarked-it-against-haproxy-d22a763db8a8)

#### Famous Projects Created with Go

- [Docker](https://github.com/docker/docker) A platform to develop, ship, and run applications in containers.
- [Kubernetes](https://github.com/kubernetes/kubernetes) An open-source system for automating deployment of containerized applications.
- [Hugo](https://github.com/gohugoio/hugo) A fast and modern static site generator.
- [Terraform](https://github.com/hashicorp/terraform) A tool for building, changing, and versioning infrastructure.
- [CockroachDB](https://github.com/cockroachdb/cockroach) A cloud-native, distributed SQL database.
- [InfluxDB](https://github.com/influxdata/influxdb) An open-source time series database.
- [Ethereum Go (Geth)](https://github.com/ethereum/go-ethereum) The official Go implementation of the Ethereum protocol.
- [Caddy](https://github.com/caddyserver/caddy) An open-source web server with automatic HTTPS.
- [Syncthing](https://github.com/syncthing/syncthing) A continuous real-time file synchronization program.
- [Dgraph](https://github.com/dgraph-io/dgraph) A distributed and transactional native GraphQL database.

# Installation and Setup

### Installing Go

The first step is to download and install Go. You can get the download [here](https://go.dev/doc/install).

<details open><summary>Windows</summary>

1. Open the MSI file you downloaded and follow the prompts to install Go.
2. You can change the location of your installation as needed. After installing, you will need to close and reopen any open command prompts.
3. Verify that you've installed Go. In Windows, click the Start menu. In the menu's search box, type cmd, then press the Enter key.
4. In the Command Prompt window that appears, type the following command:

```bash
> go version
```

5. Confirm that the command prints the installed version of Go.
</details>

<details><summary>Mac</summary>

1. Open the package file you downloaded and follow the prompts to install Go.
2. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.
3. Verify that you've installed Go by opening a command prompt and typing the following command:

```bash
$ go version
```

4. Confirm that the command prints the installed version of Go.
</details>

<details><summary>Linux</summary>

1. Remove any previous Go installation by deleting the /usr/local/go folder (if it exists)
2. Extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

```
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.25.1.linux-amd64.tar.gz
```

Warning: Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

3. Add /usr/local/go/bin to the PATH environment variable:
   You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

```
export PATH=$PATH:/usr/local/go/bin
```

Note: Changes made to a profile file may not apply until the next time you log into your computer.
To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

4. Verify that you've installed Go by opening a command prompt and typing the following command:

```
$ go version
```

Confirm that the command prints the installed version of Go.

</details>

<sub>_Source: [https://go.dev/doc/install](https://go.dev/doc/install)_</sub>

### Choosing an IDE

[GoLand by Jetbrains](https://www.jetbrains.com/go/) is an IDE from JetBrains that offers a dedicated and feature-rich experience for Go development. Since it's designed specifically for Go, it has a deep understanding of the language and provides quality code suggestions, refactoring support, and integrated tooling for debugging and testing. Students can use all of Jetbrains IDEs (including GoLand) for free with [Github's Student Developer Pack](https://www.jetbrains.com/academy/student-pack/) for the length of their studies.

[Vim Go](https://go.dev/gopls/editor/vim) is a keyboard based text editor that can be enhanced with plugins like vim-go. It's a great option for advanced users and experts but has a steep learning curve for beginners. It's ideal for those who prefer a terminal based workflow or a Unix style interface. Vim is fast, customizable, and good for rapid development.

[VSCode by Microsoft](https://code.visualstudio.com/docs/languages/go) is a powerful editor that supports multiple languages, making it easy to develop full stack Go applications without switching IDEs. It features a marketplace of extensions including an official Go extension, deep Git integration, an AI co-pilot, and more to make it a great all in one solution for doing front-end, back-end, databases, APIs, and Documentation all within the same workspace.

#### Installing VSCode Go Extension

The Official Go VSCode extension provides features designed to help any beginner get started with Go, including IntelliSense code suggestions and semantic syntax highlighting. It also offers tools like hover information for detailed insights on keywords, variables, and structs, alongside efficient keyboard shortcuts for code navigation and file formatting. Developers also benefit from a custom Go test UI, as well as support for package import fixing, refactoring, and debugging. A complete list of features and explanations are available on the [VSCode Go extension Github repository](https://github.com/golang/vscode-go/wiki/features).

1. Be sure you've installed Go on your computer and confirm its working.
2. [Download VSCode](https://visualstudio.microsoft.com/free-developer-offers/) and install it on your preferred drive.
3. Open VSCode after installing.
4. Navigate to the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) page on Microsoft Marketplace and click install with VSCode open.
5. It will ask you if you'd like to install the extension. Follow the prompts to install.
6. Close your program and restart your computer.

Microsoft has provided a link to [this helpful video](https://www.youtube.com/watch?v=1MXIGYrMk80) by Google Open Studios on setting up your Go environment in VSCode and creating your first file.

# Creating Your First Go Project in VSCode

**Perequisites:**

- Go installed and working in your console
- VSCode installed
- Go official VSCode extension installed

1.  Open Visual Studio Code and launch a new terminal using Terminal > New Terminal
2.  Create a folder for your project

```
mkdir my-first-go-project
```

3. Go to the folder using cd

```
cd my-first-go-project
```

Or, navigate to it using File > Open Folder...

4. Create your first module. This creates a go.mod file that tracks your project's dependencies.

```
go mod init my-first-go-project
```

5. Create a package folder.

```
mkdir hello
```

6. Navigate to the package folder you just created.

```
cd hello
```

7. Create a file named main.go in the hello folder:

```Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

8. Run the file in the terminal

```
go run hello/main.go
```

Or click Run > Run Without Debugging

9. Compile your file into an executable (Windows)

```
go build -o hello-app.exe hello/main.go
```

10. Run your executable

**On Mac/Linux**

```
./hello-app
```

**On Windows**

```
.\hello-app.exe
```

The complete project structure should look like this:

my-first-go-project/

├── go.mod

├── hello/

│ └── main.go

└── hello-app.exe

View the full example [here](https://github.com/omicreativedev/go-quickstart/tree/main/my-first-go-project)

### Commenting Your Code

For single line comments in Go, use two forward slashes.

Single line comments can be on their own line or they can append an existing line of code.

```Go
// This is a comment

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!") // This is also a comment
}
```

For multi-line comments, use a forward slash asteriks, asteriks forward slash, with the comment in between.

```Go
/*
This is an example
of a multi-line comment
in Go
*/

```

Go has an advanced feature call go directives that attach to comments in Go. Therefore, don't use **//go:** when making regular single line comments. For more information on go directives and other comment conventions, learn more in the [Go Comments Documentation](https://go.dev/wiki/Comments) and [Go Doc Comments](https://go.dev/doc/comment).

# Go Learning Resources

#### Official Sources

- **[go.dev](https://go.dev/)**
  - [go.dev/play/](https://go.dev/play/) The Go playground for writing, running, and sharing code online.
  - [go.dev/doc/](https://go.dev/doc/) The official documentation for the Go programming language.
  - [go.dev/doc/effective_go](https://go.dev/doc/effective_go) Official documentation of clean coding in Go.
  - [go.dev/doc/faq](https://go.dev/doc/faq) Frequently Asked Questions (and answers) about the Go language.
  - [go.dev/ref/spec](https://go.dev/ref/spec) The official Go language specification.
  - [go.dev/blog/](https://go.dev/blog/) The official Go blog for news and in-depth article.
  - [pkg.go.dev](https://pkg.go.dev/) The official Go package discovery and reference site.
- **[Golang.org](https://tour.golang.org)**
- [github.com/golang](https://github.com/golang) The official GitHub organization for the Go project.

#### Websites

- [Gobyexample.com](https://gobyexample.com/)
- [Golangbot.com](https://golangbot.com/learn-golang-series/)
- [Practical-go-lessons.com](https://www.practical-go-lessons.com/)
- [Geeksforgeeks.org/go-language](https://www.geeksforgeeks.org/go-language/)
- [W3schools.com/go/](https://www.w3schools.com/go/index.php)

#### Tutorials

- [Exercism.org](https://exercism.org/tracks/go)
- [Gophercises.com](https://gophercises.com/)

#### Books

- [go101.org](https://go101.org/)
- [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Spaceship Go](https://blasrodri.github.io/spaceship-go-gh-pages/) by
- [How to Code in Go](https://assets.digitalocean.com/books/how-to-code-in-go.pdf) by Digital Ocean
- [Go from the Beginning](https://leanpub.com/go-from-the-beginning) by Chris Noring and [Code Repository](https://github.com/softchris/golang-book)
- [Anti-Textbook Go Book](https://leanpub.com/antitextbookGo/) by (???) and [Code Repository](https://github.com/thewhitetulip/web-dev-golang-anti-textbook)

#### Videos

- [youtube.com/c/Justforfunc](https://youtube.com/c/Justforfunc)
- [Go From The Beginning Youtube Playlist](https://www.youtube.com/playlist?list=PLRbYLREdOhfnLJl0xJxKEPe4BhlDU66Zj)

#### Other

- [Gophercises with Calhoun](https://courses.calhoun.io/courses/cor_gophercises)
- [Algorithms in Go](https://www.calhoun.io/lets-learn-algorithms/)

------------------------
![Part 2](https://github.com/omicreativedev/go-quickstart/blob/main/images/part_2.png?raw=true "Part 2")
# Type System and Variable Semantics

**Go is statically typed**  (as opposed to dynamically typed), similar to Java, C++, and Rust. This means the type of a variable is known and checked at compile time, unlike in Python, JavaScript, and Ruby, where types are determined and checked at runtime.

```go
// Go
count := 10
count = "hello" // Error
```
```python
# Python
count = 10
count = "hello"  # Valid in Python
```

However, Go supports type inference with the ```:=``` operator. **Go allows both explicit type declarations and implicit type inference.** The compiler infers its type from the value you assign at compile time (not runtime) when you use the ```:=``` operator, but once inferred, the variable type is fixed. 

```
# Explicit
var count int = 10
```
```
# Implicit
count := 10 // The compiler infers this as an int
```
Because **Go is strongly typed**, the language prevents operations between incompatible types and does not perform implicit type conversions. For example, you cannot add a string and an integer. This is stricter than JavaScript (weakly typed) but similar to Python (which is also strongly typed). Learn more about type conversions below.

Variables declared with ```var``` or the short declaration operator ```:=``` are **mutable***. Their value can be changed.
```
x := 5
x = 10
```
However, constants declared with the reserved ```const``` keyword are **immutable**. Their value must be known at compile time and can't be changed.
```
const pi = 3.14
pi = 2.71 // Error
```
> [!NOTE]
> Reference types like slices/maps/channels are always mutable.

# Data Types

### Basic Types

#### Boolean

[bool](https://www.w3schools.com/go/go_boolean_data_type.php) can carry only the true or false value. It's default value is always false.

```
func main() {
	var isSunny bool = true
	var isRaining bool

	if isSunny && !isRaining { // If it's sunny AND NOT raining
		fmt.Println("Let's go outside.")
	} else { // otherwise...
		fmt.Println("Let's stay indoors.")
	}
}
```


#### Numeric

Signed and unsigned integers in Go have generic types and byte specific types. For instance, ```int``` is 64 bits on a 64 bit system. However, if you want to limit it to just 8, you could use ```int8```. The same goes for unsigned integers.

Signed Integers can store both positive and negative values.

- int is generic and platform dependent. They are 32 bits in 32 bit systems and 64 bit in 64 bit systems.
- int8
- int16
- int32 is also a rune (see below)
- int64

Unsigned Integers cannot hold negative values.

- uint is also generic and platform dependent. They are 32 bits in 32 bit systems and 64 bit in 64 bit systems.
- uint8 is also a byte (see below)
- uint16
- uint32
- uint64
- uintptr is an unsigned integer type large enough to hold the bit pattern of any pointer. It is used in low-level programming with the unsafe package. It should be used with extreme caution. More on that here.

```int``` and ```uint``` are implementation-dependent. This sometimes causes portability issues across 32 bit and 64 bit systems. Therefore its convention to specify which explicity in most cases or infer its type by value.

>[!WARNING]
>uintptr is NOT garbage collected.

Read [more about integers from w3schools.com](https://www.w3schools.com/go/go_integer_data_type.php)

**Floating-Point Numbers** are like floats in python, which are used for both 32 bit and 64 bit decimals numbers.

- float (without byte specification) defaults to float64 but in Go it's convention to specify the float type explicity or infer it by value.
- float32 is similar to floats in C++ and Java
- float64 is similar to doubles in C++ and Java

**Complex Numbers** are the set of all complex numbers with float real and imaginary parts

- complex64 are float 32 real and imaginary parts
- complex128 are float 64 real and imaginary parts

Learn more about numeric types on [go.dev/ref/spec#Numeric_types](https://go.dev/ref/spec#Numeric_types).

#### Strings

In Go, a **string** is an immutable sequence of bytes that is interpreted as UTF-8 text.

There are a variety of ways to initialize a string in Go. The most common ways are:

```Go
var s1 string = "Hello, Go!"
s2 := "Hello, World!"
```

Raw strings can span multiple lines. If you want to preserve all charachters in a raw string without escapes use single back ticks

```Go
raw := `Hello,
World.
\nThis is ignored.`
fmt.Println(raw)
```
Once a string is created, it cannot be changed.
Indexing a string returns a byte, not a full Unicode (rune).
Slicing a string preserves the UTF-8 encoding so you can extract parts of it.

```Go
package main
import "fmt"
func main() {
	text := "Hello"
	fmt.Println("First byte:", text[0])  // 72
	fmt.Println("Slice [0:1]:", text[0:1]) // H (UTF-8)
	fmt.Println("Rune:", []rune(text)) // [72 101 108 108 111] (Unicode)
}
```


#### Aliases

To reduce confusion and help distinguish the intent of a value, Go has a few aliases that make handling data easier.

A **rune** is an alias for ```int32```, a rune holds a full 32-bit Unicode character making it easy for working with non-ASCII text.
```Go
var r rune = '⌘'
```

A **byte** is an alias for ```uint8```, and is used for a variable meant to be a raw piece of 8-bit data like an ASCII character.
```Go
var b byte = 'a'
```

### Composite Types

#### Aggregate

An **array** is a fixed-length sequence of elements of a single type. These types are composed of elements or fields, which are themselves other types. Arrays cannot hold different types at the same time though, like python lists. For that functionality, see interfaces (below.)

Arrays can hold 
- Basic types: int, float64, bool, string, etc.
- Composite types: struct, [n]Type (arrays), pointers, functions, etc.
- Empty interface [n]interface{} can hold any type (mixed types)

```
var numbers [5]int // An array of 5 integers initialized to [0 0 0 0 0]
var names [2]string{"Sally","Dave"} // An array of 2 strings initialized to [Sally, Dave]
```


A **struct** is a collection of named fields, where each field can be of a different type.
```Go
type Person struct { 
Name string; 
Age int
}
```
Most of the labor done by classes in OOP languages like Java are done with structs in Go. A key difference is that Go separates the struct (data) from the methods or functions (behavior). Instead they are bound with a receiver (see binding below.) **There is no inheritance**, so Go uses struct embedding. Structs can be embedded into one another to create complex data relationships.

```Go
type Person struct {
    Name string
    Age  int
}
type Employee struct {
    Person // Embedded Person struct
    EmployeeID string
}
```


### Reference

A **pointer** holds the memory address of a variable.
```Go
*int
*MyStruct
```

Each variable or object occupies storage space in memory.
You can retrieve the address of a variable using the & operator:
```Go
var a int = 10
ptr := &a
fmt.Println(ptr)
```

### Function Type
In Go, a **function** can also be a type, allowing functions to be passed as arguments and assigned to variables.
```Go
func(int, int) int
```
### Interface Types
Go uses interfaces to achieve **polymorphism**. An interface is a collection of method signatures, and any type that implements all the methods of an interface can be treated as that interface's type. This is different from class-based inheritance where a subclass must explicitly inherit from a superclass.

A variable of an interface type can hold any concrete value that implements all the methods in the interface.

Example: **Empty Interface**
```Go
var anyValue interface{}
anyValue = 42          // Can hold int
anyValue = "hello"     // Can hold string
anyValue = []float64{1.2, 3.4} // Can hold slice
```
Example: **Error Interface**
```
type error interface {
    Error() string
}
type MyError struct {
    Message string
}
func (e MyError) Error() string {
    return "Error: " + e.Message
}
```
Example: **Stringer Interface**
```Go
type Stringer interface {
    String() string
}
type Person struct {
    Name string
    Age  int
}
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}
```

### Complex Types

**Slice ([]T)** dynamic-sized, flexible view into an array. This is one of the most used data structures, replacing arrays for most use cases.
```Go
[]int
[]string
```
**Map (map[K]V)** is an unordered collection of key-value pairs similar to a python dictionary.
```Go
map[string]int
```
**Channel (chan T)** is conduit for sending and receiving values with the arrow <- operator used for communication between goroutines (lightweight threads).
```Go
chan int, chan<- string (send-only), <-chan bool (receive-only)
```

### Type Conversion

Go requires **explicit type conversions** between different types. Unlike some languages (e.g., JavaScript, Python, C), Go does not perform implicit coercion, even between numeric types. You must tell the compiler exactly how to convert one type to another using the syntax:

- **Numeric types:** All conversions between numeric types are explicit (int to float64, float64 to int, int32 to int64, etc.) Overflow and truncation can occur silently.

- String ↔ byte slice:
&nbsp;&nbsp; ```[]byte("hello")``` creates a slice of bytes from a string.
&nbsp;&nbsp; ```string([]byte{104, 101, 108, 108, 111})``` converts bytes back into a string.

- Rune ↔ string: Converting a rune (int32) to string produces a one-character string containing the UTF-8 encoding of the rune.

- Untyped constants: An untyped constant can be assigned to variables of different types without explicit conversion until it’s given a concrete type:
```
const n = 5
var a int32 = n
var b float64 = n
```
#### Invalid conversions
- You cannot directly convert a string to an int; you must use helper functions like strconv.Atoi

- Interfaces: A value of one type can be assigned to an interface if the type implements the interface methods — this is not a "conversion" but implicit interface assignment. To get the original type back, you need a type assertion:
```
var i interface{} = 42
v := i.(int) 
```
# Syntax

### Reserved Words

Go has 25 reserved words that cannot be used as identifiers such as variable names.

- break
- case
- chan
- const
- continue
- default
- defer
- else
- fallthrough
- for
- func
- go
- goto
- if
- import
- interface
- map
- package
- range
- return
- select
- struct
- switch
- type
- var

### Variable Naming Requirements & Conventions

#### Required by Compiler

- A variable name must begin with a letter or an underscore _. The remaining characters can be letters, digits, or underscores.
- Names are case-sensitive so myNum and MyNum are different variables.
- If an identifier needs to be visible outside its package (exported), it must start with a capital letter.
- In Go, identifiers that start with a capital letter are exported (public), while those starting with lowercase are unexported (package-private).
- The underscore _ blank identifier has a special role and is used to ignore values, e.g. in assignments or imports.
- You always have to specify either type or value (or both).

#### Encouraged by Professional and Community Standards

- Acronyms should be in all caps: ServeHTTP, urlAPI, etc.
- CamelCase is preferred.
- Don't use underscores for common variable names despite them being legal.
- Full meaningful words for variables specific to your program i.e. serverAWS not s1
- Use := only when introducing a new variable.
- Use = if you only want to reassign an existing variable.
- Single-method interfaces are often named with the method name plus -er: Reader, Writer, Stringer
- Exported (Public) Identifiers: If an identifier starts with a capital letter, it is exported (visible from outside its package). This is Go's mechanism for public visibility.
```
fmt.Println // Println is exported because it starts with 'P'.

http.ListenAndServe
```
- Unexported (Private) Identifiers: If an identifier starts with a lowercase letter, it is unexported and only accessible within the package it's declared in.
```
myHelperFunction

internalCounter
```
- Use common short words for readability. The more professional Go code you'll read, you'll notice some patterns appear often:<br/><br/>
&nbsp;&nbsp;i, j, etc - used often in nested loops<br/>
&nbsp;&nbsp;n - for counts or number<br/>
&nbsp;&nbsp;p - pointer<br/>
&nbsp;&nbsp;r - io.Reader<br/>
&nbsp;&nbsp;w - io.Writer<br/>
&nbsp;&nbsp;rw - io.ReadWriter<br/>
&nbsp;&nbsp;err - error<br/>
&nbsp;&nbsp;db - database<br/>
&nbsp;&nbsp;cfg - config<br/>

### Composite Literals

Go supports composite literals, which provide a concise way to construct values for arrays, slices, maps, and structs.

- ```arr := [3]int{1, 2, 3}  ```
- ```s := []string{"a", "b", "c"}  ```
- ```m := map[string]int{"one": 1, "two": 2}  ```
- ```p := Person{Name: "Alice", Age: 30}  ```

### Type Aliases vs. Defined Types

Go distinguishes between defined types and type aliases:

Defined type: creates a new, distinct type
```
type MyInt int  
var x MyInt = 10  
var y int = 20  
```

Type alias: another name for an existing type
```
type MyIntAlias = int  
var a MyIntAlias = 30  
var b int = 40  
```

### Operators

Go has a standard set of C-like operators.

- **Arithmetic:** +, -, *, /, % 

- **Comparison:** ==, !=, <, <=, >, >=

- **Logical:** &&, ||, !

- **Bitwise:** & (and), | (or), ^ (xor), &^ (and not), << (left shift), >> (right shift)

- **Assignment:** =, +=, -=, *=, /=, %=, etc.

- **Address / Pointer:** & (address of), * (dereference)

- **Channel**: (used for sending/receiving from channels)

- **Increment/Decrement operators:** ++,  --

Go only allows increment/decrement as statements, not expressions"

```
i++ // is valid
x = i++ // is invalid
```

In Go, ```:=``` can redeclare a variable if there’s at least one new variable being declared in the same statement.
This can lead to shadowing, where an inner variable hides an outer one.

```Go
package main

import "fmt"

func main() {
    x := 10
    fmt.Println("Outer x:", x)

    if true {
        x := 20 // 👈 Shadows the outer x
        fmt.Println("Inner x:", x)
    }

    fmt.Println("Outer x again:", x) // Still 10, inner x is gone
}
```

**Go does not allow mixed-type operations** without an explicit conversion. This is a core tenet of its strong typing.
```Go
var x int32 = 10
var y int64 = 20

// sum := x + y // Compile Error
sum := int64(x) + y // This works
```
The one exception is that untyped constants (like const n = 5) can be mixed in expressions until they are assigned to a variable.

### Availability

- **Numbers (int, float, complex):** + - * / % (mod only for ints), comparisons (== != < <= > >=), bitwise ops (& | ^ &^ << >>, only for ints).

- **Strings:** + (concatenation), comparisons (== != < <= > >= lex order).

- **Booleans:** && || !, comparisons (== !=).

- **Pointers:** * (dereference), & (address of), == != (compare addresses).

- **Interfaces:** == != (two interfaces equal if both dynamic type and value are equal, or both nil).

- **Structs:** == != only if all fields are comparable types.

- **Arrays:** == != if element type is comparable.

- **Slices, Maps, Functions, Channels:** only == != against nil.

- **Channels:** additionally, <- for send/receive.

### Binding


| Scenario                          | Example                                   | Binding Time                                                                                                                        |
| --------------------------------- | ----------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| **Variable Declaration**          | `var x int = 5`                           | **Compile time** (the name `x` is bound to a variable object)                                                                       |
| **Short Variable Declaration**    | `x := 10`                                 | **Runtime (within compile constraints)**; the type and object are determined at compile time, but initialization happens at runtime |
| **Function Declaration**          | `func add(a, b int) int { return a + b }` | **Compile time**; the identifier `add` is bound to the function’s entry point                                                       |
| **Constant Declaration**          | `const Pi = 3.14`                         | **Compile time**; immutable binding                                                                                                 |
| **Type Definition**               | `type MyInt int`                          | **Compile time**; binds a new name to an existing or new type                                                                       |
| **Package Imports**               | `import "fmt"`                            | **Compile time**; binds the identifier `fmt` to the imported package namespace                                                      |
| **Interface and Method Bindings** | Interface methods bound to concrete types | **Runtime (dynamic dispatch)** when the interface is assigned a concrete value                                                      |
### Storage, Addresses, and Lifetime

- **Stack Allocation:** Local variables that don’t escape their function are stored on the stack.

- **Heap Allocation:** Variables that “escape” (e.g., returned by a function or captured by a closure) are stored on the heap, managed by the garbage collector.

- **Addressability:** Not all expressions have addresses (e.g., constants, temporary values, and function results are not addressable).

Static Lifetime applies to package level variables and exist for the duration of the program:
```
var counter int
```
Local Lifetime variables exist only during the execution of their containing function:
```
func demo() {
    x := 5 
}
```
Heap Lifetime variables that returned references exist until garbage is collected:
```
func makeCounter() *int {
    c := 0
    return &c
}
```
### Scope
- **Universe Scope:** Built-in identifiers available everywhere ex.```int, true, len```

- **Package Scope:** Identifiers declared at the top level of a package file

- **File Scope:** Applies to imports and variables declared in a single file

- **Function Scope:** Names declared inside a function are visible only there

 - **Block Scope:** Identifiers within {} are visible only within that specific block
```
var global = "package scope"

func demo() {
    local := "function scope"
    {
        inner := "block scope"
        fmt.Println(inner)
    }
    // fmt.Println(inner) // Error
}
```

# Limitations

- := cannot be used at package level; must declare at least one new variable.
- const only for primitive values; cannot be slice/map/channel.
- Go’s enumerations rely on iota. That’s part of how constants are typically used in practice.
- untyped constants. These are more flexible than variables until assigned a type.

```
const n = 5
var x int32 = n   // allowed
var y float64 = n // also allowed
```
- Arrays/structs: == only if elements/fields are comparable.
- Slices/maps/functions: cannot be compared except to nil.
- Must use explicit conversion for mixed numeric types (e.g., int + float64).
- Arrays and slices are homogeneous — cannot store different types unless using interface{}.
- Type conversion syntax (T(v)) is explicit and limited — some require helper functions (strconv.Atoi for string → int).
- No operator overloading, no implicit type coercion.
- Function equality → only comparable against nil (not against other functions).

#### Pitfalls

- Zero values: variables are auto-initialized (0, "", nil, false) — can cause logic errors if assumptions differ.
- Slices and maps are reference types: assigning them copies the reference, not the underlying data.
- Nil interfaces: (type, value) pairs — an interface holding a typed nil is not equal to nil.
- nil is the zero value for reference types (slice, map, channel, pointer, interface, function).
- But their behavior differs: nil slices can still be appended to, while nil maps panic on write.
- Shadowing with := can silently redefine outer variables.

---
![Part 3](https://github.com/omicreativedev/go-quickstart/blob/main/images/part_3.png?raw=true "Part 3")

# Control Statements: Loops, Selection, Conditionals

As mentioned under the DataTypes section, [Boolean](https://github.com/omicreativedev/go-quickstart?tab=readme-ov-file#boolean) takes only true and false which is important to remember as we discuss conditionals.

Go officially supports if/else and switch statements. Unlike C++/Java, Go has no official ```else if``` keyword.  Also, Go doesn't use parenthesis around control statements like C++ and writes them directly, similar to Python.

### If/else

```Go
if x == true {
    fmt.Println("x is true")
} else {
    fmt.Println("x is false")
}

if x > 0 && y < 10 {
    fmt.Println("Both conditions met")
}
```
Else if is **emulated** using else followed by if which, is a nuanced distinction.

```Go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}
```
Source: (Go Documentation Spec. - If Statements)[https://go.dev/ref/spec#If_statements]

### For Loops

Go has only for loop that replaces while, do/while, and foreach from other languages.

```Go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

To emulate a while loop, Go is smart and recognizes when a condition is not met.
```Go
x := 5
for x > 0 {
    fmt.Println(x)
    x--
}
```
Example emulation of a **for each**
```Go
numbers := []int{1, 2, 3}
for index, value := range numbers {
    fmt.Printf("%d: %d\n", index, value)
}
```
Go also can print and infinite loop if not careful.
```Go
for {
    fmt.Println("Forever")
    break
}
```
#### Continue and Break

Inside a for loop, you can use break to stop it.

```Go
for x := 0; n <= 10; x++ {
  if x == 4 {
    break
  }
  fmt.Println(x)
}
```
Continue only stops the execution of the current iteration. It continues with the next one.

```Go
for x := 0; x <= 10; x++ {
  if x%2 == 0 {
    continue
  }
  fmt.Println(x)
}
```
```range``` here iterates through the numbers slice
```Go
for i, num := range numbers {
        if num > 25 {
            fmt.Printf("Index %d: %d is greater than 25\n", i, num)
        }
    }
```
#### Labels

Labels can be used with break and continue to control exactly from which loop we want to break or continue

```Go
func main() {
ControlBreak:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Println(i)
            break ControlBreak
        }
    }
}
```
Source: [Exercism -- For Loops](https://exercism.org/tracks/go/concepts/for-loops)
### Switch Statements

Go's switch automatically breaks after each case. It also accepts multiple values!

```Go
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Almost weekend")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek")
}
```
The word ```fallthrough`` is used to cause the switch to go to the next case.
```Go
switch number {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
        fallthrough // <<< GO TO THE NEXT CASE
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Other number")
    }
```
Source:
[Go.dev -- Switch Statements](https://go.dev/ref/spec#Switch_statements)
[Go.dev -- Fallthrough Statements](https://go.dev/ref/spec#Fallthrough_statements)

### Block Delimiting

Go uses explicit braces {} for code blocks. This pevents the **danging else** issue. Opening braces must be on the same line, never the next line. It doesn't use semicolons. 

```Go
if x > 5 {
    fmt.Println("Braces Always Required)
}
```
Source: [](https://go.dev/doc/faq#semicolons)

### Short Curcuit Evaluation

Go uses short-circuit evaluation like Java/C++. Logical operators && and || evaluate left-to-right.

```Go
package main

import "fmt"

func main() {
    if false && printEx() {
        // This won't run
    }
    if true || printEx() {
        fmt.Println("Done") 
    }
}

func printEx() bool {
    fmt.Println("This won't run either")
    return true
}
```
Source: [Go.dev -- Operators](https://go.dev/ref/spec#Operators)

### Why Is Go Missing Things?

GGo omits many features like ternary operator, parentheses around conditions, while loops, and inheritance) because its primary design goal is simplicity. As Rob Pike (Go co-designer) stated: "The key point here is our programmers are Googlers, they're not researchers..." So the designers believe simpler is more maintainable. 

Readability over writeability - Code should be clear to readers not writers

Simplicity over complexity - Fewer features mean fewer bugs

Explicit over implicit - Nothing magical that's hard to trace

Compilation speed - Minimal features enable fast builds

Source: [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)

---
![Part 4](https://github.com/omicreativedev/go-quickstart/blob/main/images/part_4.png?raw=true "Part 4")

# Function Declaration Syntax

In Go, functions are declared using the ```func``` keyword followed by a name, parameters in parentheses, and optional return types. Functions can accept multiple parameters of different types, and parameters can be grouped that share the same type declaration.

Go is compiled, so function order doesn't matter. There are no specific function placement rules. Functions can be declared anywhere in the package, above or below a function call.

Source: [Go - Function declarations](https://go.dev/ref/spec#Function_declarations)

Example with multiple parameters:

```Go
func add(a int, b int) int {
    return a + b
}
```
Example with different data types:

```Go
func printInfo(name string, age int) {
    fmt.Println(name, age)
}
```
Example with parameters grouped:

```Go
func multiply(x, y int) int {
    return x * y
}
```
Go controls visibility through capitalization rules where uppercase names are public and lowercase are private.

Source: [Go Blog](https://go.dev/doc/effective_go#names)

Example:

```Go
func PublicFunc() {}  // Accessible outside package
func privateFunc() {} // Accessible in this package
```
A special feature of Go is that functions can return multiple values, which should be specified after the parameters, and these return values can be named.

Source: [Go - Function types](https://go.dev/ref/spec#Function_types)

Example returning multiple values at the same time:

```Go
func getName() (string, string) {
    return "Sally", "Sue"
}
```

For flexibility, functions can be created that accept a variable number of arguments using the ```...``` syntax on the final parameter.

Source: [Go - Function types](https://go.dev/ref/spec#Function_types)

```Go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```
Functions in Go can be attached to types as methods using receivers, written without names as anonymous functions, and return other functions.

Source: [Go - Method declarations](https://go.dev/ref/spec#Method_declarations)

```Go
type Rectangle struct { width, height int }

func (r Rectangle) area() int {
    return r.width * r.height
}

func main() {
    square := func(x int) int { return x * x }
}

```
Go also provides helpful features like the ```defer``` keyword for cleanup actions.

Source: [Go Blog - Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)

```Go
func example() {
    defer fmt.Println("Done")
    fmt.Println("Working...")
}

```
# Function Scope

Unlike Python's function-level scope and Java's block-level scope with hoisting, Go uses strict block-level scope without hoisting, meaning variables exist only within their declared blocks and aren't accessible before they're declared.

### Scope Hierarchy

| Scope Level | Visibility | Lifetime | Unusual Characteristics |
|-------------|------------|----------|-------------------------|
| Universe | Predeclared identifiers | Entire program | Built-in types and functions like `int`, `println` |
| Package | Exported identifiers across files | Entire program | Capitalized names are accessible outside package |
| File | Imported packages | File duration | Imports are file-scoped, NOT package-wide |
| Function | Parameters, local variables | Function execution | Can shadow package-level variables |
| Block | Variables declared in blocks | Block execution | Includes if, for, switch statements, etc. |

Example Code:

```Go
var global = "global"

func main() {
    local := "local"
    if true {
        block := "block"
        fmt.Println(global, local, block) // Accessible
    }
    // fmt.Println(block) // Error: block not accessible
}
```
Source: [Go - Declarations & Scope](https://go.dev/ref/spec#Declarations_and_scope)<br>
Source: [Effective Go](https://go.dev/doc/effective_go#names)

# Pass By

Go is pass-by-value by default. When you pass arguments to functions, Go creates copies of the values. When you pass pointers, slices, maps, or channels, you're copying the reference, allowing modification of the original data.

Source: [Go FAQ - Functions](https://go.dev/doc/faq#pass_by_value)

Pass-by-Value
```Go
func doubleValue(x int) {
    x = x * 2
}

func main() {
    num := 5
    doubleValue(num)
    fmt.Println(num) // 5 - unchanged
}
```

Pass-by-Reference

```Go
func doubleReference(x *int) {
    *x = *x * 2
}

func main() {
    num := 5
    doubleReference(&num)
    fmt.Println(num) // 10 - modified
}
```

### Side Effects

Side effects occur when functions modify data outside their local scope. Go allows side effects through pointers, but provides protection with arrays by passing them as copies. This prevents accidental modifications to original array data.

```Go
package main
import "fmt"

func main() {
    nums := [3]int{1, 2, 3}
    modifyArray(nums)
    fmt.Println(nums) // [1 2 3] - unchanged
}

func modifyArray(arr [3]int) {
    arr[0] = 99
    fmt.Println(arr) // [99 2 3] - only local copy changed
}
```
Source: [Go - Assignments](https://go.dev/ref/spec#Assignments)

### Guardrails

Go has built-in protections against side effects through its pass-by-value behavior for basic types and arrays. When you pass arrays and basic types to functions, Go creates copies, preventing accidental modifications to the original data.

```Go
package main
import "fmt"

func main() {
    data := [2]int{1, 2}
    update(data)       // Array is copied
    fmt.Println(data)  // [1 2] NOT changes
}

func update(d [2]int) {
    d[0] = 99         // Modifies only the copy
}
```
# Memory Management: Stack vs Heap 

Go automatically manages memory using two primary areas: the stack and the heap. The stack provides fast access for short-lived data like function arguments and local variables, while the heap stores data that needs to persist longer or is shared across function boundaries. Go uses escape analysis to determine where to store variables. If a variable's lifetime extends beyond its function, it "escapes" to the heap.

Function Storage:

- Stack: Function arguments, parameters, return addresses, and most local variables
- Heap: Data that escapes function scope, large allocations, and shared data

During Execution:

- Local Variables: Typically on stack, unless returned or shared (then heap)
- Arguments: Copied onto stack when function is called
- Parameters: Stored on stack as local variables within the function

Source: [Go Blog - Escape Analysis](https://go.dev/doc/faq#stack_or_heap)
Source: [Go Memory Management](https://go.dev/doc/effective_go#memory)

### Recursive Functions

Recursive functions are functions that call themselves to solve smaller instances of the same problem. Each recursive call creates a new stack frame with its own parameters and local variables. 
```Go
func factorial(n int) int {
    if n <= 1 {
        return 1  // Base case stops recursion
    }
    return n * factorial(n-1)  // Recursive call
}

func main() {
    result := factorial(5)
    fmt.Println(result) 
}
```
