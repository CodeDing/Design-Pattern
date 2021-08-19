package Filter

import (
    `bytes`
    `fmt`
    `strings`
)

type Filter func(*Person) bool

func WithMale(gender string) Filter {
    return func(p *Person) bool {
        if gender == strings.ToLower(p.Gender()) {
            return true
        }
        return false
    }
}

func WithFemale(gender string) Filter {
    return func(p *Person) bool {
        if gender == strings.ToLower(p.Gender()) {
            return true
        }
        return false
    }
}

func WithSingle(marital string) Filter {
    return func(p *Person) bool {
        if marital == strings.ToLower(p.MaritalStatus()) {
            return true
        }
        return false
    }
}

func WithSingleMale(gender, marital string) Filter {
    return func(p *Person) bool {
        if gender == strings.ToLower(p.Gender()) && marital == strings.ToLower(p.MaritalStatus()) {
            return true
        }
        return false
    }
}

func WithSingleOrFemale(gender, marital string) Filter {
    return func(p *Person) bool {
        if gender == strings.ToLower(p.Gender()) || marital == strings.ToLower(p.MaritalStatus()) {
            return true
        }
        return false
    }
}

func NewPersonsWithFilter(persons []*Person, filter Filter) []*Person {
    res := make([]*Person, 0, len(persons))
    for _, p := range persons {
        if filter(p) {
            res = append(res, p)
        }
    }
    return res
}

const (
    MALE   = "male"
    FEMALE = "female"
    SINGLE = "single"
)

func IsMale() Filter {
    return func(p *Person) bool {
        if MALE == strings.ToLower(p.Gender()) {
            return true
        }
        return false
    }
}

func IsFemale() Filter {
    return func(p *Person) bool {
        if FEMALE == strings.ToLower(p.Gender()) {
            return true
        }
        return false
    }
}

func IsSingle() Filter {
    return func(p *Person) bool {
        if SINGLE == strings.ToLower(p.MaritalStatus()) {
            return true
        }
        return false
    }
}

func NewPersonsWithFilters(persons []*Person, filters ...Filter) []*Person {
    res := make([]*Person, 0, len(persons))
    for _, p := range persons {
        match := true
        for _, f := range filters {
            if !f(p) {
                match = false
                break
            }
        }
        if match {
            res = append(res, p)
        }
    }
    return res
}

func FormatPerson(persons ...*Person) string {
    if len(persons) == 0 {
        return "<nil>"
    }
    var buf bytes.Buffer
    for _, p := range persons {
        buf.WriteString(fmt.Sprintf("Name: %s, Gender: %s, Marital: %s\n", p.Name(), p.Gender(), p.MaritalStatus()))
    }
    return buf.String()
}
