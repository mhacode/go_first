
# Create a temp directory

~$ mkdir temp

~$ cd temp

## Download go from official site

https://golang.org/dl/


## Extract and install go in the temp dirctory:

~/temp$ tar -xvf go1.16.4.linux-amd64.tar.gz

~/temp$ ls go

~/temp$ clear

## Move file from temp directory to local directory 
~/tmp$ sudo mv go /usr/local

~/tmp$ rm -rf go

cd ~

## Check go version

$ go version

## If go version is not found then export these command
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
## Check go verion again

go version

## If go verion not present yet then  open nano and add these command

nano .bashrc
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
## Write this command if verion is not shown

source ~/.bashrc


