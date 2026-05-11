package grade

import (
	"errors"

	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"
)

type GradeResult struct {
	Score int
	Grade string
}

func GenerateContent(foodType string, nutrition entity.Nutrition) (*GradeResult, error) {
	var points map[string][]float64
	var calculate func(AllPoint) int
	var check func(int) string

	switch foodType {
	case "food":
		points = PointsFood
		calculate = CalculatePointFood
		check = CheckGradeFood
	case "beverage":
		points = PointsBeverage
		calculate = CalculatePointBeverage
		check = CheckGradeBeverage
	default:
		return nil, errors.New("invalid type: must be 'food' or 'beverage'")
	}

	allPoint := AllPoint{
		Energy:                getPoint(nutrition.Energy, points["energy"]),
		SaturatedFatty:        getPoint(nutrition.SaturatedFatty, points["saturated_fatty"]),
		Sugar:                 getPoint(nutrition.Sugar, points["sugar"]),
		Salt:                  getPoint(nutrition.Salt, points["salt"]),
		Protein:               getPoint(nutrition.Protein, points["protein"]),
		Fibres:                getPoint(nutrition.Fibres, points["fibres"]),
		FruitVegetableLegumes: getPoint(nutrition.FruitVegetableLegumes, points["fruit_vegetable_legumes"]),
	}

	score := calculate(allPoint)
	return &GradeResult{
		Score: score,
		Grade: check(score),
	}, nil
}
