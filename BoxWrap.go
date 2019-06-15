package uis

import (
	"errors"
	"github.com/andlabs/ui"
)

type BoxBuilder struct {
	*ui.Box
	stretch bool
	children map[string]ui.Control
}

func NewHBox() *BoxBuilder{
	return BoxWrap(ui.NewHorizontalBox())
}
func NewVBox() *BoxBuilder{
	return BoxWrap(ui.NewVerticalBox())
}

func BoxWrap(u *ui.Box) *BoxBuilder{
	r:=new(BoxBuilder)
	r.Box=u
	u.SetPadded(true)
	r.stretch=false
	return r
}
func (self *BoxBuilder) SetStretch(s bool) *BoxBuilder{
	self.stretch=s
	return  self
}
func (self *BoxBuilder) append(u ui.Control) {
	self.Box.Append(u,self.stretch)
}
func (self *BoxBuilder) Appends(controls ...ui.Control) *BoxBuilder{
	for _,i:=range  controls{
		self.append(i)
	}
	return self
}
func (self *BoxBuilder) addNamedControl(named string,control ui.Control) error{
	if _,ok:=self.children[named];ok{
		return errors.New("can't overwrite old named control")
	}
	self.children[named]=control
	return nil
}
func (self *BoxBuilder)AppendBtn(named,text string,f func (*ui.Button))(*BoxBuilder,error,* ui.Button) {
	ret:=ui.NewButton(text)
	_,e:=self.AppendControl(named,ret)
	ret.OnClicked(f)
	return self,e,ret
}
func (self *BoxBuilder)AppendEntry(named  string)(*BoxBuilder,error,*ui.Entry) {
	ret:=ui.NewEntry()
	_,e:=self.AppendControl(named,ret)
	return self,e,ret
}
func (self *BoxBuilder)AppendControl(named string,control ui.Control)(*BoxBuilder,error) {
	err:=self.addNamedControl(named,control)
	self.append(control)
	return self,err
}

func (self *BoxBuilder) ToBox() *ui.Box{
	return self.Box
}

func (self *BoxBuilder) GetChild(named string) (ui.Control,bool) {
	r,ok:= self.children[named]
	return r,ok
}