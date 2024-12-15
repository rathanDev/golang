
https://www.youtube.com/watch?v=d_L64KT3SFM
Build REST api with Go 

use Gin library 

# ----- ----- # 
# Y Go

With the advancement of hardware, servers can handle multiple threads,
though most programming languages don't support multithreads by default,
when writing multithreaded programming in Java, code becomes very complex

Go was designed to 
    run on multiple cores and 
    built to support concurrency 

Concurrency in Go is cheap and easy 

Main use case of Go:
For performant applications 
Running on scaled, distributed systems 

Use to build server side / backend apps 
Microservices 
Web apps 
database services 

# ----- ----- # 
Use vscode 

install Go plugin 
Go 
from Go team at Google

# ----- ----- # 

go mod init togo 
    generates go.mod file 

# ----- ----- # 

# the syntax reflects the natural flow of how we think 
    var tickets uint = 500;
    a variable ticket which is uint

# = vs :=

    = 
        assigment operator 
        var x int
        x = 50
        x =  x + 5

    :=
        short declaration operator 
        x := 50
        x = x + 5

        generally used to contain the returned values of a func 
        age := calculateAge(dob uint)

        can't use for package level variables 

# func validate        vs       func Validate 
    validate 
        accessible within the same package 
    Validate 
        exported 
        accessible from other packages 

# go routine 

# ----- ----- # 







