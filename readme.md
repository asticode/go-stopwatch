# About

This is a goroutine manager for the GO programming language (http://golang.org).

# Dependencies

    github.com/asticode/go-toolbox

# Installing

## Using `go get`

    $ go get github.com/asticode/go-stopwatch
    
After this command `go-stopwatch` is ready to use. Its source will be in:

    $GOPATH/src/github.com/asticode/go-stopwatch/stopwatch
    
# Configuration

Best is to use JSON configuration:

    {
        "id": "ID",
        "enabler": {
            "headers": [],
            "ip_addresses": []
        }
    }
    
And decode it:

    oStopWatchConfiguration = json.UnMarshall(sConfiguration)
    
An example of configuration would be:

    {
        "id": "my-stopwatch",
        "enabler": {
            "headers": {
                "X-MyStopWatch": "1234"
            },
            "ip_addresses": [
                "127.0.0.1"
            ]
        }
    }
    
# Example

    import (
        "github.com/asticode/go-stopwatch/stopwatch"
    )

    // Create stop watch.
    oStopwatch, oErr := stopwatch.NewStopwatch(oStopWatchConfiguration)
    
    // Push
    oStopwatch.Push("Initialization", "The Stopwatch object has been created")
    
    // Get stopwatch as json
    aJson = oStopwatch.Json()
    