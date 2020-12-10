package medium

func predictPartyVictory(senate string) string {
	var r, d []int
	for i, c := range senate {
		if c == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	n := len(senate)
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, n+r[0])
		} else {
			d = append(d, n+d[0])
		}
		r = r[1:]
		d = d[1:]
	}

	if len(r) > 0 {
		return "Radiant"
	}

	return "Dire"

}
