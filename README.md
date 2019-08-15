# Message of the day
Basic CLI program written in go to show a message of the day (motd) or write the message to a file. 

### CLI usage
Possible valid flags.
```
$ go run main.go
Usage:
  motd [flags]

Flags:
  -g, --greeting string   phrase to use within the greeting
  -h, --help              help for motd
  -n, --name string       name to use within the message
  -v, --preview           use preview to output message without writing to ./file.txt
  -p, --prompt            use prompt to input name and message
exit status 1
```

Use `DEBUG` and preview and parameters:
```
$ DEBUG=true go run main.go --preview --name lvthillo --greeting Hello
Name: lvthillo
Greeting: Hello
Prompt: false
Preview: true
```

Use preview and short parameters:
```
$ go run main.go --preview --name lvthillo --greeting Hello 
Hello, lvthillo
```

Use preview and prompt:
```
$ go run main.go --prompt --preview
Your Greeting: Hello
Your Name: lvthillo
Hello, lvthillo
```

Use parameters:
```
$ go run main.go --name lvthillo --greeting Hello
$ cat file.txt
Hello, lvthillo
```

Use short prompt flag:
```
$ go run main.go -p
Your Greeting: Hello
Your Name: lvthillo
$ cat file.txt
Hello, lvthillo
```
