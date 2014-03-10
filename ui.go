package main

import (
  "github.com/mattn/go-gtk/gtk"
  "github.com/mattn/go-gtk/gdk"
  "github.com/mattn/go-gtk/glib"
  "os"
  "fmt"
)


func renderUI(cfg *Configuration) (error) {

  gtk.Init(&os.Args)
  window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
  window.Connect("destroy", gtk.MainQuit)

  table := gtk.NewTable(5, 5, false)
  window.Add(table)
  window.ShowAll()
  window.SetSizeRequest(200, 200)

  glib.TimeoutAdd(5000, func() bool {
    var jenkins Jenkins
    err := jenkins.GetBuilds(cfg.Host)
    if err != nil {
      fmt.Println(err)
    }

    children := table.GetChildren()
    for n := uint(0); n < children.Length(); n++ {
      switch widget := children.NthData(n).(type) {
      case gtk.IWidget:
       table.Remove(widget)
      }
    }

    var column uint = 0
    var row uint = 0
    for _, build := range jenkins.Builds {
      eventbox := gtk.NewEventBox()
      label := gtk.NewLabel(build.Name)

      eventbox.ModifyBG(gtk.STATE_NORMAL, gdk.NewColorRGB(0, 255, 72))
      if build.LastBuildStatus != "Success" {
        eventbox.ModifyBG(gtk.STATE_NORMAL, gdk.NewColorRGB(178, 0, 0))
      }
      if build.Activity == "Building" {
        eventbox.ModifyBG(gtk.STATE_NORMAL, gdk.NewColorRGB(255, 202, 0))
      }

      eventbox.Add(label)
      table.AttachDefaults(eventbox, column, (column + 1), row, (row + 1))
      column++
      if column == 4 {
        column = 0
        row++
      }
    }
    window.ShowAll()
    return true
  })

  gtk.Main()
  return nil
}
