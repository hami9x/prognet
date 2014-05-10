prognet
=======

Anti-js js Golang programmer network web code.  
This project is developed alongside Go-AngularJs as a glorious dogfood eater.  

##Installing:
    
    go get github.com/gopherjs/gopherjs  
    go get github.com/gopherjs/go-angularjs  
    go get github.com/phaikawl/prognet/client
    go get github.com/phaikawl/prognet
    go get github.com/pilu/fresh
    cd $GOPATH/src/github.com/pilu/fresh  
    git remote add pk https://github.com/phaikawl/fresh.git  
    git pull pk master  
    go build && go install  
  

##Running:  
cd to the `client` directory:
    
    gopherjs build -o="../public/app/scripts/app.js"
`cd ..` to the project directory, run:
    
    fresh
the site is usually on http://localhost:3000

##Developing:  
Install [sass](http://sass-lang.com/install)  
In the directory,
run

    fresh
Then run
    
    ./start
It starts all the automatic compilers (including gopherjs and sass).  

##Interested?  
Your help is needed, contact me at phaikawl[at]gmail[dot]com to join the team.
