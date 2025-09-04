module example.com/greetings

go 1.24.5

replace example.com/hello => ../hello

require (
	example.com/goroutines v0.0.0-00010101000000-000000000000
	example.com/hello v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.28.0 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace example.com/goroutines => ../goroutines
