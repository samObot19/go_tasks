package calculate

import "errors"


func Grade(score float64) (error, string){
	switch {
	case score > 0 && score < 40:
		return nil, "F"
	case score >= 40 && score < 50:
		return nil, "D"
	case score >= 50 && score < 60:
		return nil, "C"
	case score >= 60 && score < 65:
		return nil, "C+"
	case score >= 65 && score < 70:
		return nil, "B-"
	case score >= 70 && score < 75:
		return nil, "B"
	case score >= 75 && score < 80:
		return nil, "B+"
	case score >= 80 && score < 85:
		return nil, "A-"
	case score >= 85 && score < 90:
		return nil, "A"
	case score >= 90 && score <= 100:
		return nil, "A+"
	default:
		return errors.New("Invalid score : none"), ""

	}

	return nil, ""
}

