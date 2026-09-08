module example.com/hello

go 1.27.1

//this was placed in here by using the go mod edit command 'go mod edit -replace example.com/greetings=../greetings'
// this is used for when you are using a library that has not been published to a remote repository
replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
