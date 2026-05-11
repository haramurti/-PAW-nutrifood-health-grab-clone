package grade

// Sentinel value pengganti 9999999 di JS
const maxVal = 9999999.0

// null di JS (fruit_vegetable_legumes punya slot kosong) → kita skip dengan -1
// Logic: kalau nilai di tabel -1, point untuk index itu tidak dihitung (skip ke next)
// Tapi karena kita cuma butuh "cari index pertama dimana nilai <= threshold",
// null di sini artinya threshold itu tidak valid → kita pakai maxVal biar tidak pernah tercapai lebih awal.

var PointsFood = map[string][]float64{
	"energy":                  {335, 670, 1005, 1340, 1675, 2010, 2345, 2680, 3015, 3350, maxVal},
	"sugar":                   {3.4, 6.8, 10, 14, 17, 20, 24, 27, 31, 34, 37, 41, 44, 48, 51, maxVal},
	"saturated_fatty":         {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, maxVal},
	"salt":                    {0.2, 0.4, 0.6, 0.8, 1, 1.2, 1.4, 1.6, 1.8, 2, 2.2, 2.4, 2.6, 2.8, 3, 3.2, 3.4, 3.6, 3.8, 4, maxVal},
	"protein":                 {2.4, 4.8, 7.2, 9.6, 12, 14, 17, maxVal},
	"fibres":                  {3, 4.1, 5.2, 6.3, 7.4, maxVal},
	"fruit_vegetable_legumes": {40, 60, 80, maxVal, maxVal, maxVal},
}

var PointsBeverage = map[string][]float64{
	"energy":                  {30, 90, 150, 210, 240, 270, 300, 330, 360, 390, maxVal},
	"sugar":                   {0.5, 2, 3.5, 5, 6, 7, 8, 9, 10, 11, maxVal},
	"saturated_fatty":         {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, maxVal},
	"salt":                    {0.2, 0.4, 0.6, 0.8, 1, 1.2, 1.4, 1.6, 1.8, 2, 2.2, 2.4, 2.6, 2.8, 3, 3.2, 3.4, 3.6, 3.8, 4, maxVal},
	"protein":                 {1.2, 1.5, 1.8, 2.1, 2.4, 2.7, 3.0, maxVal},
	"fibres":                  {3, 4.1, 5.2, 6.3, 7.4, maxVal},
	"fruit_vegetable_legumes": {40, maxVal, 60, maxVal, 80, maxVal, maxVal},
}
