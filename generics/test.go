package main

import "fmt"
type BaseFields struct { MyField int }
type JustBase struct { MyField int }
type MoreThanBase struct {
        MyField int
        MyOtherField int
}

type FieldConstraint interface { BaseFields }

// constrain struct to type of FieldConstraint
func Constrained[T FieldConstraint](param T) {
        fmt.Println(param.MyField)
}

func main() {
        var _ BaseFields = BaseFields(JustBase{})    // this works in Go already!
        Constrained[JustBase](JustBase{})            // seems like this should work then!
        Constrained[MoreThanBase](MoreThanBase{})    // should this work?
}

