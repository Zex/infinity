package truth

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"

  //cfg "truth/config"
  pb "truth_pb"
)

func FindPeer(w http.ResponseWriter, req *http.Request) {
  log.Printf("FindPeer (%v, %v)", w, req)
  w.Header().Set("Content-Type", "application/json")

  peers := []pb.Peer{}

  rsp, err := json.Marshal(map[string][]pb.Peer{
    "peers": peers})

  if err != nil {
    log.Printf("json marshal failed: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s", rsp)
}
