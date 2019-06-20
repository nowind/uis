package main

import(
	"github.com/nowind/uis"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)


func SetUi(w *ui.Window ) ui.Control{
	main:=uis.NewVBox()
	return main.
		AppendForm("你好").
		AppendBtn("OK","OK", func(button *ui.Button) {
		nh,_:=main.GetChild("你好")
		ui.MsgBox(w,"",nh.(*ui.Entry).Text())
	})
}
func main(){
	uis.MainWindow(800,200,SetUi)
}
