//go:build ignore

package main

/*
	both channels and go routines are
	used for concurrent programming

	line 21 :-
		http.Get waits for response
		this is very slow for what we are tryin to create

	line 25 :-
		adding go creates a go routine
		and creates a illusion of go concrrency
		using a go scheduler ...
				go scheduler uses only one cpu by default
				and we need to configure it to use two or more

					true concurrency is parallelism
						when two or more cpus are involved ...

		when a http call goes out it switches to next go routine

		there is still a issue that main routine ends
			without processing other go routines
				here channels come into play

	existing program is a main routine
		and go routine is a separate process similar to go routine only




*/
