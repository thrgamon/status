package main
import (
    "os"
    "github.com/fatih/color"
)

func main() {
  status := os.Args[1]
  switch status {
    case "100":
      color.Blue("Continue                      ")
    case "101":
      color.Blue("Switching Protocols           ")
    case "200":
      color.Blue("OK                            ")
    case "201":
      color.Blue("Created                       ")
    case "202":
      color.Blue("Accepted                      ")
    case "203":
      color.Blue("Non-Authoritative Information ")
    case "204":
      color.Blue("No Content                    ")
    case "205":
      color.Blue("Reset Content                 ")
    case "206":
      color.Blue("Partial Content               ")
    case "300":
      color.Blue("Multiple Choices              ")
    case "301":
      color.Blue("Moved Permanently             ")
    case "302":
      color.Blue("Found                         ")
    case "303":
      color.Blue("See Other                     ")
    case "304":
      color.Blue("Not Modified                  ")
    case "305":
      color.Blue("Use Proxy                     ")
    case "307":
      color.Blue("Temporary Redirect            ")
    case "400":
      color.Blue("Bad Request                   ")
    case "401":
      color.Blue("Unauthorized                  ")
    case "402":
      color.Blue("Payment Required              ")
    case "403":
      color.Blue("Forbidden                     ")
    case "404":
      color.Blue("Not Found                     ")
    case "405":
      color.Blue("Method Not Allowed            ")
    case "406":
      color.Blue("Not Acceptable                ")
    case "407":
      color.Blue("Proxy Authentication Required ")
    case "408":
      color.Blue("Request Timeout               ")
    case "409":
      color.Blue("Conflict                      ")
    case "410":
      color.Blue("Gone                          ")
    case "411":
      color.Blue("Length Required               ")
    case "412":
      color.Blue("Precondition Failed           ")
    case "413":
      color.Blue("Payload Too Large             ")
    case "414":
      color.Blue("URI Too Long                  ")
    case "415":
      color.Blue("Unsupported Media Type        ")
    case "416":
      color.Blue("Range Not Satisfiable         ")
    case "417":
      color.Blue("Expectation Failed            ")
    case "426":
      color.Blue("Upgrade Required              ")
    case "500":
      color.Blue("Internal Server Error         ")
    case "501":
      color.Blue("Not Implemented               ")
    case "502":
      color.Blue("Bad Gateway                   ")
    case "503":
      color.Blue("Service Unavailable           ")
    case "504":
      color.Blue("Gateway Timeout               ")
    case "505":
      color.Blue("HTTP Version Not Supported    ")
  }
}
