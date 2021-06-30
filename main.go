package main

import (
  "io"
  "bytes"
  "strings"
  "net/http"

  log "github.com/sirupsen/logrus"
  "github.com/gorilla/mux"
)


func main() {



  log.Info("Starting the HTTP server on port 8888")

  router := mux.NewRouter().StrictSlash(true)
  router.Path("/upload").Methods("POST").HandlerFunc(uploadGPXFile)

  log.Fatal(http.ListenAndServe(":8888", router))
}

func uploadGPXFile(w http.ResponseWriter, r *http.Request) {

  // https://stackoverflow.com/a/28074084
  var maxBytes int64 = 5 * 1024  * 1024
  r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

  // 
  r.ParseMultipartForm(maxBytes)

  var buf bytes.Buffer

  file, header, err := r.FormFile("file")
  if err != nil {
    log.Warn(err)
    w.WriteHeader(http.StatusNotAcceptable)
    return
  }
  defer file.Close()

  name := strings.Split(header.Filename, ".")
  log.WithFields(log.Fields{"FileName": name[0], "FileSize": header.Size}).Info("file received")
  io.Copy(&buf, file)

  contents := buf.String()
  log.Info(contents)
}
