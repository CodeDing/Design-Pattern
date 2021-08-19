package Filter

import `testing`

func TestFilter2(t *testing.T) {
    persons := []*Person{
        {"Robert", "Male", "Single"},
        {"John", "Male", "Married"},
        {"Laura", "Female", "Married"},
        {"Diana", "Female", "Single"},
        {"Mike", "Male", "Single"},
        {"Bobby", "Male", "Single"},
    }
    t.Logf("Origin persons:\n %s", FormatPerson(persons...))
    male := NewPersonsWithFilter(persons, WithMale("male"))
    female := NewPersonsWithFilter(persons, WithFemale("female"))
    single := NewPersonsWithFilter(persons, WithSingle("single"))
    singleMale := NewPersonsWithFilter(persons, WithSingleMale("male", "single"))
    singleOrFemale := NewPersonsWithFilter(persons, WithSingleOrFemale("female", "single"))

    t.Logf("Male: \n%s", FormatPerson(male...))
    t.Logf("Female: \n%s", FormatPerson(female...))
    t.Logf("Single: \n%s", FormatPerson(single...))
    t.Logf("SingleMale: \n%s", FormatPerson(singleMale...))
    t.Logf("SingleOrFemale: \n%s", FormatPerson(singleOrFemale...))
}

func TestFilter3(t *testing.T) {
    persons := []*Person{
        {"Robert", "Male", "Single"},
        {"John", "Male", "Married"},
        {"Laura", "Female", "Married"},
        {"Diana", "Female", "Single"},
        {"Mike", "Male", "Single"},
        {"Bobby", "Male", "Single"},
    }

    male := NewPersonsWithFilters(persons, IsMale())
    female := NewPersonsWithFilters(persons, IsFemale())
    single := NewPersonsWithFilters(persons, IsSingle())
    singleMale := NewPersonsWithFilters(persons, IsMale(), IsSingle())

    t.Logf("Male: \n%s", FormatPerson(male...))
    t.Logf("Female: \n%s", FormatPerson(female...))
    t.Logf("Single: \n%s", FormatPerson(single...))
    t.Logf("SingleMale: \n%s", FormatPerson(singleMale...))
}
