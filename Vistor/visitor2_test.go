package Vistor

import `testing`

func TestVisitor2(t *testing.T) {
    info := &Info{}
    var v Visitor = info
    v = NameVisitor{v}
    v = OtherVisitor{v}
    v = LogVisitor{v}
    infoFunc := func(info *Info, err error) error {
        info.Namespace = "namespace1"
        info.Name = "name1"
        info.Action = DELETE
        return nil
    }
    v.Visit(infoFunc)

    t.Logf("=======================================")
    v2 := &Info{Namespace: "namespace2", Name: "name2", Action: GET}
    d := NewDecorator(v2, NameVisitorFunc(), LogVisitorFunc())
    d.Visit(func(info *Info, err error) error {
        t.Logf("Decorator: namespace=%s, name=%s, action=%s", info.Namespace, info.Name, info.Action)
        return err
    })
}
