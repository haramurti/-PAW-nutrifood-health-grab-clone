package helper

import "github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"

type gradeInfo struct {
	Title       string
	Color       string
	Image       string
	URL         string
	Description map[string]string
}

var grades = map[string]gradeInfo{
	"A": {
		Title: "Nutri-Score A",
		Color: "#008439",
		Image: "https://get.apicbase.com/nutri-score-science-based-nutritional-value-labelling-system/",
		URL:   "/images/grade.png",
		Description: map[string]string{
			"food":     "Foods with a nutrition score from high to 0. This means that foods in this category have the highest nutritional quality.",
			"beverage": "Drinks with a nutrition score from 1 to 0. This means that drinks in this category have the highest nutritional quality.",
		},
	},
	"B": {
		Title: "Nutri-Score B",
		Color: "#75BE00",
		Image: "https://get.apicbase.com/nutri-score-science-based-nutritional-value-labelling-system/",
		URL:   "/images/grade.png",
		Description: map[string]string{
			"food":     "Foods with a nutrition score from 1 to 2. The nutritional quality is still good, although slightly lower than category A.",
			"beverage": "Beverages with a nutrition score from 1 to 2. The nutritional quality is still good, although slightly lower than category A.",
		},
	},
	"C": {
		Title: "Nutri-Score C",
		Color: "#FFC800",
		Image: "https://get.apicbase.com/nutri-score-science-based-nutritional-value-labelling-system/",
		URL:   "/images/grade.png",
		Description: map[string]string{
			"food":     "Foods with a nutrition score from 3 to 10. The nutritional quality is decent, but not as good as categories A and B.",
			"beverage": "Beverages with a nutrition score from 3 to 10. The nutritional quality is decent, but not as good as categories A and B.",
		},
	},
	"D": {
		Title: "Nutri-Score D",
		Color: "#FF7800",
		Image: "https://get.apicbase.com/nutri-score-science-based-nutritional-value-labelling-system/",
		URL:   "/images/grade.png",
		Description: map[string]string{
			"food":     "Foods with a nutrition score from 11 to 18. Their nutritional quality is lower than the previous category.",
			"beverage": "Beverages with a nutrition score from 11 to 18. Their nutritional quality is lower than the previous category.",
		},
	},
	"E": {
		Title: "Nutri-Score E",
		Color: "#FA2300",
		Image: "https://get.apicbase.com/nutri-score-science-based-nutritional-value-labelling-system/",
		URL:   "/images/grade.png",
		Description: map[string]string{
			"food":     "Food with a nutrition score higher than 18. This is the category with the lowest nutritional quality.",
			"beverage": "Beverages with a nutrition score above 18. This is the category with the lowest nutritional quality.",
		},
	},
}

func GetGrade(grade, foodType string) entity.GradeDetail {
	g, ok := grades[grade]
	if !ok {
		return entity.GradeDetail{Grade: grade}
	}
	return entity.GradeDetail{
		Grade:       grade,
		Title:       g.Title,
		Description: g.Description[foodType],
		Color:       g.Color,
		Image:       g.Image,
		URL:         g.URL,
	}
}
