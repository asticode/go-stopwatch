# About

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/asticode/go-stopwatch/stopwatch)

`go-stopwatch` is a multi-event stopwatch that can measure memory and time deltas between events during runtime for the GO programming language (http://golang.org).

# Install `go-stopwatch`

Run the following command:

    $ go get github.com/asticode/go-stopwatch/stopwatch
    
# Examples
## Basic example

    import (
        "github.com/asticode/go-stopwatch/stopwatch"
    )
    
    // Create the stopwatch
    //
    // "myproject" is the stopwatch id and is printed in event names to set them apart from children events.
    // true is whether the stopwatch is enabled and will actually add events in its slice. This is useful when you want to add stopwatch events conditionally in your process.
    //
    s, e := stopwatch.NewStopwatch("myproject", true)
    
    // Add the first event
    s.AddEvent("My first event", "This is my first description")
    
    // Sleep during 4ms
    time.Sleep(4 * time.Millisecond)
    
    // Add the second event
    s.AddEvent("My second event", "This is my second description")
    
    // Sleep during 10ms
    time.Sleep(10 * time.Millisecond)
    
    // Add the third event
    s.AddEvent("My third event", "This is my third description")
    
    // Print the results
    fmt.Println(s.String())
    
This will output:

    Stopwatch results:
    Id: myproject
    Number of events: 3
    Time start: 11 Nov 2015 17:27:51 +0000
    Memory start: 3.804MB
    
    Name: myproject - My first event
    Description: This is my first description
    
    +4.161ms
    +0.000MB
    
    Name: myproject - My second event
    Description: This is my second description
    
    +11.261ms
    -0.006MB
    
    Name: myproject - My third event
    Description: This is my third description
    
## JSON example

You can also retrieve the `go-stopwatch` results as a JSON instead of a string:

    // Get results as a JSON-friendly struct
    j := s.JSON()
    
    // Marshall JSON
    o, e := json.Marshal(j)
    