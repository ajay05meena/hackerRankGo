package cats_and_a_mouse

import "math"

func CatAndMouse(catOnePos, catTwoPos, ratPos int32) string{
	distanceBetweenCatOneAndRat := math.Abs(float64(catOnePos - ratPos));
	distanceBetweenCatTwoAndRat := math.Abs(float64(catTwoPos - ratPos));

	if distanceBetweenCatOneAndRat == distanceBetweenCatTwoAndRat {
		return "Mouse C"
	}else if distanceBetweenCatOneAndRat > distanceBetweenCatTwoAndRat {
		return "Cat B"
	}else{
		return "Cat A"
	}
}
