package Filter

import (
    `bytes`
    `fmt`
    `strings`
)

type Person struct {
    name          string
    gender        string
    maritalStatus string
}

func newPerson(name, gender, maritalStatus string) *Person {
    return &Person{name, gender, maritalStatus}
}

func (p Person) Name() string {
    return p.name
}

func (p Person) Gender() string {
    return p.gender
}

func (p Person) MaritalStatus() string {
    return p.maritalStatus
}

type Persons []*Person

func newPersons() *Persons {
    var p Persons
    p = make([]*Person, 0)
    return &p
}

func (p *Persons) Add(persons ...*Person) {
    if p == nil {
        panic("persons can not be nil")
    }
    *p = append(*p, persons...)
}

func (p Persons) String() string {
    if len(p) == 0 {
        return "<nil>"
    }
    var buf bytes.Buffer
    for _, i := range p {
        buf.WriteString(fmt.Sprintf("Name: %s, Gender: %s, MateriaStatus: %s\n", i.Name(), i.Gender(), i.MaritalStatus()))
    }
    return buf.String()
}

type Criteria interface {
    MeetCriteria([]*Person) []*Person
}

type CriteriaMale struct{}

func (CriteriaMale) MeetCriteria(persons []*Person) []*Person {
    res := make([]*Person, 0, len(persons))
    for _, p := range persons {
        if strings.ToLower(p.Gender()) == "male" {
            res = append(res, p)
        }
    }
    return res
}

type CriteriaFemale struct{}

func (CriteriaFemale) MeetCriteria(persons []*Person) []*Person {
    res := make([]*Person, 0, len(persons))
    for _, p := range persons {
        if strings.ToLower(p.Gender()) == "female" {
            res = append(res, p)
        }
    }
    return res
}

type CriteriaSingle struct{}

func (CriteriaSingle) MeetCriteria(persons []*Person) []*Person {
    res := make([]*Person, 0, len(persons))
    for _, p := range persons {
        if strings.ToLower(p.MaritalStatus()) == "single" {
            res = append(res, p)
        }
    }
    return res
}

type AndCriteria struct {
    criteria      Criteria
    otherCriteria Criteria
}

func (a AndCriteria) MeetCriteria(persons []*Person) []*Person {
    firstCriteriaPersons := a.criteria.MeetCriteria(persons)
    return a.otherCriteria.MeetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
    criteria      Criteria
    otherCriteria Criteria
}

func (o OrCriteria) MeetCriteria(persons []*Person) []*Person {
    firstCriteriaItems := o.criteria.MeetCriteria(persons)
    otherCriteriaItems := o.criteria.MeetCriteria(persons)
    for _, p := range otherCriteriaItems {
        exist := false
        for _, i := range firstCriteriaItems {
            if p.Name() == i.Name() {
                exist = true
                break
            }
        }
        if !exist {
            firstCriteriaItems = append(firstCriteriaItems, p)
        }
    }
    return firstCriteriaItems
}
