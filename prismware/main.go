package main
import (
    "log"
    "os"
    "path/filepath"
    "strings"
    // "net/http"
)

func main() {


  apisDir := os.Args[1]

  e := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
      fullpath := path
      if info.IsDir() == false {
        prefixes := [2]string{apisDir,"/"}
        postfixes := [2]string{info.Name(),"/"}
        for _, prefix := range prefixes {
          // sum += num
          path = strings.TrimPrefix(path, prefix)
        }
        for _, postfix := range postfixes {
          // sum += num
          path = strings.TrimRight(path, postfix)
        }
        route := path

        println(fullpath,route)
      }
      return nil
  })
  if e != nil {
      log.Fatal(e)
  }
}


// /srv # curl localhost:2019/config/apps/http/servers/
// {"srv0":{"listen":[":80"],"logs":{"default_logger_name":"log0"},"routes":[{"handle":[{"handler":"vars","root":"/usr/share/caddy/"}],"match":[{"path":["/"]}]},{"handle":[{"handler":"subroute","routes":[{"handle":[{"handler":"rewrite","strip_path_prefix":"/store"}]},{"handle":[{"handler":"reverse_proxy","headers":{"request":{"set":{"Host":["{http.request.host}"],"X-Forwarded-For":["{http.request.remote.host}"],"X-Forwarded-Port":["{http.request.port}"],"X-Forwarded-Proto":["{http.request.scheme}"],"X-Real-Ip":["{http.request.remote.host}"]}}},"upstreams":[{"dial":"prism:4010"}]}]}]}],"match":[{"path":["/store*"]}]},{"handle":[{"handler":"file_server","hide":["/etc/caddy/Caddyfile"]}]}]}}
