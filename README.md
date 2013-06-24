RabbitMQGoDemo
==============

To run this example, you need a correctly configured go environment.
For example, put this in your .profile or .bashrc:

`export GOPATH=$GOPATH:~/myGoWorkspace`  
`export PATH=$PATH:~/myGoWorkspace/bin`  

You'll also need to get the following dependencies dependencies (requires git and mercurial):

`go get github.com/streadway/amqp`  
`go get code.google.com/p/goprotobuf/proto`  

Now get this example and build it:

`go get github.com/Sophiacom/RabbitMQGoDemo`  
`go install github.com/Sophiacom/RabbitMQGoDemo`
