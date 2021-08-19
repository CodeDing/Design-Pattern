package Filter

import `testing`

func TestFilter(t *testing.T) {
    persons := newPersons()
    persons.Add(
        newPerson("Robert", "Male", "Single"),
        newPerson("John", "Male", "Married"),
        newPerson("Laura", "Female", "Married"),
        newPerson("Diana", "Female", "Single"),
        newPerson("Mike", "Male", "Single"),
        newPerson("Bobby", "Male", "Single"),
    )

    male := CriteriaMale{}
    female := CriteriaFemale{}
    single := CriteriaSingle{}
    singleMale := AndCriteria{single, male}
    singleOrFemale := OrCriteria{single, female}

    t.Logf("Males: %s", Persons(male.MeetCriteria([]*Person(*persons))))
    t.Logf("Females: %s", Persons(female.MeetCriteria([]*Person(*persons))))
    t.Logf("Single: %s", Persons(single.MeetCriteria([]*Person(*persons))))
    t.Logf("SingleMale: %s", Persons(singleMale.MeetCriteria([]*Person(*persons))))
    t.Logf("SingleOrFemale: %s", Persons(singleOrFemale.MeetCriteria([]*Person(*persons))))
}
