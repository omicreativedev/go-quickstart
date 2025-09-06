

# Welcome to Go Quickstart

This repository is for my CS-330 class at Simmons University where my goal is to create a simple straightfoward tutorial for getting started using Go for students already comfortable with at least one programming language. This is a work-in-progress so there will be some errors, omissions and incomplete sections. When the project is complete and the semester is over, I will tag it **complete** and the repo will allow pull requests. I'm currently working in Windows 11, and while I try to give instructions for Mac and Linux, I can't test them all on my machine and so I'm mainly going by documentation sources. Your miles may vary. Until then, if you have any issues or commentary, please see the [Discussion](https://github.com/omicreativedev/go-quickstart/discussions) section of the repository.

If you need to reach me, message me on [LinkedIn](https://www.linkedin.com/in/sshunamon/)

Credits:

Side by Side Markdown Editor: [Dillinger.io](https://dillinger.io/)

Hosting: [Render.com](https://render.com/) _Free Tier_

<!-- Testing Browser IDE -->
<!-- Testing Drafts between Wip and Main -->
------------------------
🔸Part 1




# History of the Go Programming Language

The Go programming language, often referred to as Golang, was created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson as a direct response to the frustrations experienced in software development within the company.[^1] The main catalyst for its creation was the difficult nature of using existing languages for massive systems work, lengthy compilation times for languages like C and the complexity of distributed systems built in languages like Java.[^2] The designers were also motivated by a desire to improve the scale of development for large teams of programmers working on shared codebases.[^3] Officially announced in 2009 as an open-source project, Go was designed to be a simple language that emphasized ease of use for modern hardware but with the power of other languages already in use at the company.[^4]

If you want an interesting read about the beginnings of Go, check out the official [Go Spec Blob on Github](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) by Robert Griesemer himself.

[^1]: [Go FAQ: History](https://go.dev/doc/faq#history)
[^2]: [Go FAQ: Creating a New Language](https://go.dev/doc/faq#creating_a_new_language)
[^3]: [The Go Programming Language Specification (Preface)](https://go.dev/ref/spec )
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
[^14]: [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)


### What is Go good for?

- Cloud-Native & Distributed Systems (Microservices)
- Command-Line Interfaces (CLIs) & DevOps Tools
- Web Servers & API Backends
- Concurrent Network Services
- Data Processing & Pipelines
- Databases & Storage Systems
- Cryptography & Security Tools
- Embedded Systems & IoT
- Scripting & Automation
- Proxy and Load Balancer Infrastructure

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

#### Perequisites:
- Go installed and working in your console
- VSCode installed
- Go official VSCode extension installed

 1. Open Visual Studio Code and launch a new terminal using Terminal > New Terminal
 2. Create a folder for your project
```
mkdir my-first-go-project
```
3. Go to the folder using cd
```
cd my-first-go-project
```
And, navigate to it using File > Open Folder...

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

On Mac/Linux
```
./hello-app
```
On Windows
```
.\hello-app.exe
```
The complete project structure should look like this:

my-first-go-project/

├── go.mod

├── hello/

│   └── main.go

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
