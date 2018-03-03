package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			fmt.Println("url need prefix http://")
			os.Exit(1)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		/* version 2*/
		//_, err = io.Copy(os.Stderr, resp.Body) // Body to Status
		//if err != nil {
		//	fmt.Println(err)
		//}

		/* version 3 */
		for k, v := range resp.Header {
			fmt.Printf("%10s\t%s\n", k, v)
		}

		/* version 1 */
		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//
		//fmt.Printf("%s", b)
	}
}
