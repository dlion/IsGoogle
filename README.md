# IsGoogle
Performs a reverse lookup for a certain ip to check if it belongs to Google crawlers 

Golang port of [is-google](https://github.com/roccomuso/is-google) nodejs package, it allows to verify that a web crawler is visiting your server using the [Google specs](https://support.google.com/webmasters/answer/80553?hl=en).

## Get

`go get https://github.com/dlion/IsGoogle`

## Example

```go
package main

import (
	"fmt"

	. "github.com/dlion/IsGoogle"
)

const (
	ip = "66.249.66.1"
)

func main() {
	answer, err := IsGoogle(ip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The address %s is belongs to Google: %v\n", ip, answer)
}
```

## Tests

`go test`

## A little copying is better than a little dependency

Yea, I know that but in my use case is not true :yolo:

## License

MIT

## Author

Domenico Luciani [@DLion92](https://twitter.com/DLion92) https://domenicoluciani.com

## Thanks

Thank you [Rocco Musolino](https://github.com/roccomuso) for the [idea](https://github.com/roccomuso/is-google)
