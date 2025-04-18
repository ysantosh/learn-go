# Get started with Installation
## Installation
### Installation on MAC
* Download tar from https://golang.org/doc/install
  ```
  tar -C /usr/local -xzf Downloads/go1.9.2.darwin-amd64.tgz
  ```
* OR install it via brew
  ```
  brew install go
  ```
### Installation on Ubuntu
* Install golang-go. 
  ```
  sudo apt  install golang-go
  ```
* Check git and if it is not installed then install it
  ```
  sudo apt update
  sudo apt install git
  ```
## Create your first program and execute
* Create directory and cd to it - Ref : https://go.dev/doc/tutorial/getting-started
  ```
  mkdir test1
  cd test1
  ```
* Create a file main.go
  ```
  vim main.go
     package main
      import "fmt"
      func main() {
        fmt.Println("Hello, World!")
      }
  ```
* Compile and run
  ```
  go mod init main.go
  go mod tidy
  go run main.go
  ```
*Issue1* : mod command not found when installed on ubuntu 18.x via apt. Install on new version of ubuntu

* Build executable which can run on any server where go is installed
  ```
  go build main.go
  ```
*Note* : Create a main.go file which can be renamed and transffered to any VM.

