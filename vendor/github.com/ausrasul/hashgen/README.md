# HashGen package for Golang

### Installation
    go get github.com/ausrasul/hashgen

### Usage

    package main
    import "github.com/ausrasul/hashgen"
    
    func main(){
        randomString := hashgen.Get(10) // returns 10 letters random strings a-zA-Z
    }
