package grade

// AllPoint menyimpan point masing-masing nutrisi setelah lookup tabel
type AllPoint struct {
	Energy                int
	SaturatedFatty        int
	Sugar                 int
	Salt                  int
	Protein               int
	Fibres                int
	FruitVegetableLegumes int
}

// getPoint mencari index pertama di tabel threshold di mana nilai <= threshold[i]
func getPoint(value float64, thresholds []float64) int {
	for i, t := range thresholds {
		if value <= t {
			return i
		}
	}
	return len(thresholds) - 1
}

func CalculatePointFood(p AllPoint) int {
	N := p.Energy + p.SaturatedFatty + p.Sugar + p.Salt
	P := p.Fibres + p.FruitVegetableLegumes
	if N < 11 {
		P += p.Protein
	}
	return N - P
}

func CalculatePointBeverage(p AllPoint) int {
	N := p.Energy + p.SaturatedFatty + p.Sugar + p.Salt
	P := p.Protein + p.Fibres + p.FruitVegetableLegumes
	return N - P
}

func CheckGradeFood(score int) string {
	switch {
	case score <= 0:
		return "A"
	case score <= 2:
		return "B"
	case score <= 10:
		return "C"
	case score <= 18:
		return "D"
	default:
		return "E"
	}
}

func CheckGradeBeverage(score int) string {
	switch {
	case score <= 0:
		return "A"
	case score <= 2:
		return "B"
	case score <= 6:
		return "C"
	case score <= 9:
		return "D"
	default:
		return "E"
	}
}
