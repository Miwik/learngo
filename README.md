# learngo
Official website: https://golang.org/

## Installation on Ubuntu
Go to https://golang.org/dl/ and download the latest Linux version, then untar it to /usr/local:  
```sudo tar -C /usr/local -xzf goVERSION.linux-amd64.tar.gz```  

Create the working Go directory in ~/go, export the PATH to the go binaries and set the GOPATH to our working dir:  
```mkdir ~/go && echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile && echo "export GOPATH=$HOME/go >> ~/.profile && source ~/.profile```  

Create the directory where we will clone the learngo repo:  
```mkdir -p ~/go/src/github/$USER && cd ~/go/src/github/$USER && git clone https://github.com/Miwik/learngo.git```
