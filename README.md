prognet
=======

Anti-js js Golang programmer network web code

Installing:  
go get github.com/phaikawl/prognet  
Ignore the errors  
cd to $GOPATH/src/gopherjs/go-angularjs  
git remote add pk https://github.com/phaikawl/go-angularjs.git  
git pull pk master  
go to $GOPATH/src/github.com/phaikawl/prognet/public  
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
cd to `client`  
run `hotreloader -p=gopherjs -a="build -o ../public/app/scripts/app.js"`  

the site is usually on http://localhost:3000
