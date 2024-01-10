<div align="center">
  
  <h1>
    <img src="https://media.tenor.com/TCMWkxIkF9IAAAAi/dancing-gopher.gif" width=40> 
    BACKEND WITH GO
    <img src="https://media.tenor.com/TCMWkxIkF9IAAAAi/dancing-gopher.gif" width=40>
  </h1>
   
</div>

<div>
  <h2>Table of content</h2>
  <ol>
    <li>
      <a href="#installation"> Installation </a>
    </li>
    <ol>
      <li>
        <a href="#macos-using-homebrew">MacOS using Homebrew</a>
      </li>
      <li>
        <a href="#windows-using-chocolatey">Windows using Chocolatey</a>
      </li>
      <li>
        <a href="#linux-using-apt">Linux using apt</a>
      </li>
      <li>
        <a href="#download-go">Download Go</a>
      </li>
      <li>
        <a href="#check-installed-go-version">Check installed Go version</a>
      </li>
    </ol>
    <li>
      <a href="#introduction">Introduction</a>
    </li>
    <ol>
      <li>
        <a href="#what-is-go">What is Go?</a>
      </li>
    </ol>
  </ol>
</div>

<div>
  <h2>Installation</h2>  
  <h3>MacOS using Homebrew</h3>
  <p>To install Go on MacOS using the terminal you can use the following command (in case you have previously installed homebrew).</p>
</div>

```bash
  brew install go
```

<div>
  <h3>Windows using Chocolatey</h3>
  <p>To install Go on Windows using PowerShell you can use the following command (in case you have previously installed Chocolatey).</p>
</div>

```bash
  choco install -y golang
```

<div>
  <h3>Linux using apt</h3>
  <p>To install Go on Ubuntu using the terminal you can use the following command.</p>
</div>

```bash
  sudo apt install golang-go
```
<div>
  <h3>Download Go</h3>
  <p>To download and install Go from the documentation, you can access the link below.</p>
  <a href="https://go.dev/dl/">Download Go</a>
</div>

<div>
  <h3>Check installed Go version</h3>
  <p>Run the following command in order to check the version installed</p>
</div>

```bash
  go version
```
<div>
  <br/>
</div>

**[⬆ Back to top](#table-of-content)**

<div align="justify">
  <h2>Introduction</h2>
  <h3>What is Go?</h3>
  <img align="right" src="https://openupthecloud.com/wp-content/uploads/2020/01/Golang.png" alt="Golang Logo" width=400/>
  <p>Go is an open-source programming language focused on simplicity, reliability, and efficiency. Go was originally designed at Google in 2007 by Rob Pike, and Ken Thompson. At the time, Google was growing quickly, and the code being used to manage its infrastructure was also growing quickly in both size and complexity. Programmers began to feel that this large and complex codebase was slowing them down so they decided that they needed a new programming language focused on simplicity and quick performance. Go became an open-source project and was released publicly in 2012. It quickly gained a surprising level of popularity and has become one of the leading modern programming languages.

</p>
  <br> <br>
  <img align="left" src="https://fingers-site-production.s3.eu-central-1.amazonaws.com/uploads/images/szLui8773HimqPgfZcnOSt1jcqsUYcJlnaHepZ50.jpg" width=250/>
  <p>Go is also a compiled and concurrent programming language, its syntax is a mix between C and Python. It is also statically typed, which means that the variable type cannot change its type throughout the program (it's not necessary to declare the type, meaning it is dynamically typed or uses duck typing). It doesn't use the “;” but it does use curly braces to define a block of code. It stands out at the level of concurrency, in Go there are no threads as such but there are "Go routines" that are lighter threads (this helps optimize resources). It's more a systems language than of application. It's more useful for the Backend (at microservices or data processing). It allows object-oriented programming but not in the conventional way, it has no inheritance and polymorphism will be done through interfaces. It does not have exceptions which adds unnecessary complexity, you have to validate a lot so that the application does not break.</p>
</div>
