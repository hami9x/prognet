prognet
=======

Anti-js js Golang programmer network web code

Installing:
go get github.com/gopherjs/gopherjs  
go get github.com/gopherjs/go-angularjs  
go get github.com/phaikawl/prognet  
cd $GOPATH/src/github.com/gopherjs/go-angularjs  
git remote add pk https://github.com/phaikawl/go-angularjs.git  
git pull pk master  
cd $GOPATH/src/github.com/phaikawl/prognet/public  
run grunt   
go get github.com/phaikawl/hotreloader  
go get github.com/pilu/fresh  
cd $GOPATH/src/github.com/pilu/fresh  
git remote add pk https://github.com/phaikawl/fresh.git  
git pull pk master  

Running:  
in the directory:  
delete `public/app/scripts/app.js` to be sure that it doesn't work if it doesn't work  
run `fresh`  
make a new terminal tab  
cd `client`  
run `hotreloader -p=gopherjs -a="build -o ../public/app/scripts/app.js"`  

the site is usually on http://localhost:3000
