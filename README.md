# ipinfo

This is an UNOFFICIAL ipinfo.io wrapper


## Features
- [x] Get basic information about ip
- [ ] Get the Summary, Whois, IP Blocks, Related Networks according with asn
## GetInfo

Create a instance of API, and retrieve information about on IP, this return a IPInfo struct.
```go

api := ipinfo.NewAPI()
info, err := api.GetInfo("8.8.8.8")

```
### IPInfo struct

This is the information marshall in the `IPInfo` struct

|  Field | type  
|---|---|
|  City |  string |    
|  Country | string   |   
|  Hostname | string   |    
|  IP |  string  |   
|  Loc | string   |    
|  Org | string   |     
|  Region| string   |       
