# swt
My personal project of encoding data. I created this only for educational purposes and for fun.
## How to install?
In order to install swt follow the instructions:
```
go get github.com/wspirrat/swt
```
After that you are ready to use swt library!
## How it works?
Inside swt.go file is created config which contains private key. Private key is used for aes encoding of SWT struct. Key is created randomly and every time you run the script the value of key is other. EncodeSWT() function creates SWT{} struct which contains payload. SWT struct is then encoded with gob to binary and after that to string encrypted with aes and base64. DecodeSWT() reads the string decrypt it with the key and then read binary. After reading binary it is assigned to SWT{} and then returned to user.
## How to use?
SWT have two main function **EncodeSWT(value)** and **DecodeSWT(payload)**. 
- EncodeSWT encode given value of map[string]interface{} to string. In order to use it:
```go
import (
  "fmt"
  
  "github.com/wspirrat/swt/swt"
)

func main() {
  m := make(map[string]interface{})
  // m[key] = value
  m["myStringValue"] = "hello :)"
  m["myint"] = 1
  m["mybool"] = true
  // also you can use any value that you want structs, bytes etc.
  // we use custom as we are using map
  encoded := swt.EncodeSWTcustom(m)
  fmt.Println(encoded)
  // jLGi0jq98ap77J_UTdYvTdFojbytuotRrj6uPCzk3orKTGf8jyUSykvPWVDRCeQc1zMbdC4bU2BohtMjKGY32
  // JfTMwYIomeXmwhBDmoFXsfADwOu6ncEo6Yi_PqD1Xn1VJ1SI4X5J
  // REUYNxOV-fGsONsyLjpMay0TcvvlgoXBGbSSNX8lt8R0G7C6WTZFz7sl0wE7XsRNeBQ7oQBDxv1JHXWelc0fsDgK
  // Nm6zsLAwpV1xWzSPM_3ELvaQsup0rwA4dp4Embkkhi-utKZk27lJsbIvcGlMi0=
}
```
- DecodeSWT decode the given value and returns SWT token with value given in map. To get some type just write .(type).
```go
decoded := swt.DecodeSWT(encoded)
fmt.Println(decoded.Payload["mybool"].(bool))
// true
```
