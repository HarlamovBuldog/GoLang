// Netcat1 - TCP-client for readonly.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

var clockWall = template.Must(template.New("clockWall").Parse(`
<h1>{{.TotalCount}} themes</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	var validArg = regexp.MustCompile(`[a-zA-z]+=localhost:(6553[0-5]|655[0-2][0-9]\d|65[0-4](\d){2}|6[0-4](\d){3}|[1-5](\d){4}|[1-9](\d){0,3})`)
	for _, argToParse := range os.Args[1:] {
		if validArg.MatchString(argToParse) == true {
			// go func here
			indexOfEqualSign := strings.Index(argToParse, "=")
			areaName := argToParse[:indexOfEqualSign]
			localhostStr := argToParse[indexOfEqualSign+1:]
			fmt.Printf("Area: %s\nHost: %s\n", areaName, localhostStr)
			conn, err := net.Dial("tcp", localhostStr)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn)
		} else {
			fmt.Printf("%s: invalid call\n", argToParse)
		}
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
