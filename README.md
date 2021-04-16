# json
A fork of encoding/json that fixes omitempty behavior on structs

https://play.golang.org/p/M8AOBJgWMJy

# Example
```
T := struct {
	B struct {} `json:",omitempty"`
}{}

json.NewEncoder(os.Stdout).Encode(T) // stdlib
NewEncoder(os.Stdout).Encode(T) // this package
```

# Output
```
{"B":{}} // stdlib
{} // this package
```
