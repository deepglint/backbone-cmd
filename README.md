#Backbone-cmd

Is the command tool for Deepglint Backbone Framework

##Prepare

First, You Shoule install Golang and NVM

when install golang, just checkout golang.org

When install NVM, chekcout https://github.com/creationix/nvm

##Get Started

```
go install github.com/deepglint/backbone-cmd

backbone-cmd component create helloworld

cd helloworld

nvm use 0.10

npm install 

webpack

go run helloworld.go

```

Open Browser and check http://localhost:8004

Or You Can Just Run

```

python -m SimpleHTTPServer 

```

and open the Browser http://localhost:8000

##Use Vulcand

```
cd yourcomponent

backbone vulcand create yourname yourcomponent http://192.168.5.46:8182 http://your-ip:8004

bash gen_to_vulcand.sh

```
and open http://192.168.5.46:8182/yourname/yourcomponent/

The Bumble_bee api:

/bumble/api/....

By yanhuang@deepglint.com