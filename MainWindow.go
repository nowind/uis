package uis

import "github.com/andlabs/ui"

func MainWindow(w,h int,f func (window *ui.Window) ui.Control){
	ui.Main(func() {
		w:=ui.NewWindow("",w,h,false)
		w.SetMargined(true)
		w.OnClosing(func(window *ui.Window) bool {
			ui.Quit()
			return false
		})
		ui.OnShouldQuit(func() bool {
			w.Destroy()
			return true
		})
		w.SetChild(f(w))
		w.Show()
	})
}
