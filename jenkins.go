package main

import (
  "encoding/xml"
  "net/http"
  "io/ioutil"
)

type Jenkins struct {
  Builds []JenkinsBuild `xml:"Project"`
}

type JenkinsBuild struct {
  Name string `xml:"name,attr"`
  LastBuildStatus string `xml:"lastBuildStatus,attr"`
  Activity string `xml:"activity,attr"`
}

func (jenkins *Jenkins) GetBuilds(host string) (error) {
  resp, err := http.Get(host)

  if err != nil {
    return err
  }

  defer resp.Body.Close()
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }
  xml.Unmarshal(data, jenkins)

  return nil
}
