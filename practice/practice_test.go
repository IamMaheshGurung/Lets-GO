package practice


import (
    "testing"
    "regexp"
)

func TestRoseName(t *testing.T) {
    name := "Mahesh"
    want := regexp.MustCompile(`\b` + name + `\b`)
    msg, err := Rose("Mahesh")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Mahesh") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
func TestRoseEmpty(t *testing.T) {
    msg, err := Rose("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
