package models

type MealsResponse struct {
	Meals []Meal `json:"meals"`
}

type Meal struct {
	IDMeal            string      `json:"idMeal"`
	StrMeal           string      `json:"strMeal"`
	StrDrinkAlternate interface{} `json:"strDrinkAlternate"`
	StrCategory       string      `json:"strCategory"`
	StrArea           string      `json:"strArea"`
	StrInstructions   string      `json:"strInstructions"`
	StrMealThumb      string      `json:"strMealThumb"`
	StrTags           interface{} `json:"strTags"`
	StrYoutube        string      `json:"strYoutube"`
	StrIngredient1    string      `json:"strIngredient1"`
	StrIngredient2    string      `json:"strIngredient2"`
	StrIngredient3    string      `json:"strIngredient3"`
	StrMeasure1       string      `json:"strMeasure1"`
	StrMeasure2       string      `json:"strMeasure2"`
	StrMeasure3       string      `json:"strMeasure3"`
	StrSource         string      `json:"strSource"`
	DateModified      interface{} `json:"dateModified"`
}
