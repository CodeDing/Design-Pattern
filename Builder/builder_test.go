package Builder

import "testing"

func TestMealFactory_PrepareVegMeal(t *testing.T) {
	mf := MealFactory{}
	t.Logf("Prepare vegmeal: %s", mf.PrepareNonVegMeal().ShowItems())
}

func TestMealFactory_PrepareNonVegMeal(t *testing.T) {
	mf := MealFactory{}
	t.Logf("Prepare non-vegmeal: %s", mf.PrepareNonVegMeal().ShowItems())
}
