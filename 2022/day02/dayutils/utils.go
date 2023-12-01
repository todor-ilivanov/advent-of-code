package dayutils

func GetScoreForMove(pick string) int {
	switch pick {
	case "r":
		return 1
	case "p":
		return 2
	case "s":
		return 3
	default:
		return 0
	}
}
