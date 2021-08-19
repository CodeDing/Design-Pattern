package Null

import `testing`

func TestNullObject(t *testing.T) {
    f := CustomerWithSpecifiedName("Rob", "Jack", "Tom")

    if f("Rob").IsNil() {
        t.Fatalf("Rob should be a real customer, but the result is false!")
    }

    if !f("Smith").IsNil() {
        t.Fatalf("Smith should be a null customer, but result is true!")
    }
    t.Logf("Real customer: %s", f("Rob").GetName())
    t.Logf("Null customer: %s", f("Smith").GetName())

}
