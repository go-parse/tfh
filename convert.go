/*
 * Copyright 2020 Vasiliy Vdovin
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tfh

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func raceTypeFromTitle(title string) RaceType {

	title = strings.ToUpper(title)

	title = strings.ReplaceAll(title, ".", "")
	title = strings.NewReplacer("INH FLAT", "BUMPER", "NATIONAL HUNT FLAT", "BUMPER", "APPRENTICE", "CONDITIONAL", "HANDICAPS", "HANDICAP", "CONDITIONS", "CONDITIONAL", "HURLE", "HURDLE", "GOLD CUP", "GOLDCUP").Replace(title)

	title = strings.NewReplacer("CHASER", "", "CHASE DAY", "", "PARK CHASE", "", "POWERS GOLD CUP", "CHASE").Replace(title)

	mt := make(map[string]bool)
	rTitle := make([]string, 0)

	isFind := false
	findNext := true

	sTitle := strings.Split(title, " ")

	for i := range sTitle {
		n := sTitle[len(sTitle)-1-i]
		rTitle = append(rTitle, n)
	}

	for _, t := range rTitle {

		if r := regexp.MustCompile(`[A-Z]+`).FindStringIndex(t); findNext && len(r) > 0 {

			t = t[r[0]:r[1]]

			switch t {
			case "CHASE":
				isFind = true
				mt[t] = true
			case "BUMPER":
				isFind = true
				mt[t] = true
			case "STAKES":
				isFind = true
				mt[t] = true
			case "HURDLE":
				isFind = true
				mt[t] = true
			case "AUCTION":
				isFind = true
				mt[t] = true
			case "MEDIAN":
				isFind = true
				mt[t] = true
			case "MAIDEN":
				isFind = true
				mt[t] = true
			case "NOVICE":
				isFind = true
				mt[t] = true
			case "SELLER":
				isFind = true
				mt[t] = true
			case "CLAIMER":
				isFind = true
				mt[t] = true
			case "HANDICAP":
				isFind = true
				mt[t] = true
			case "SALES":
				isFind = true
				mt[t] = true
			case "PATTERN":
				isFind = true
				mt[t] = true
			case "BLACK":
				isFind = true
				mt[t] = true
			case "CLAIMING":
				isFind = true
				mt[t] = true
			case "CLASSIFIED":
				isFind = true
				mt[t] = true
			case "CONDITIONAL":
				isFind = true
				mt[t] = true
			case "AMATEUR":
				isFind = true
				mt[t] = true
			default:

				if isFind {
					findNext = false
				}
			}
		}
	}

	rt := ""
	for t := range mt {

		rt += " " + t
	}

	raceType := RaceType{
		Bumper:      0,
		Hurdle:      0,
		Handicap:    0,
		Stakes:      0,
		Chase:       0,
		Classified:  0,
		Auction:     0,
		Sales:       0,
		Maiden:      0,
		Novice:      0,
		Seller:      0,
		Claimer:     0,
		Pattern:     0,
		Black:       0,
		Claiming:    0,
		Conditional: 0,
		Amateur:     0,
	}

	if strings.Contains(rt, "BUMPER") {
		raceType.Bumper = 1
	}

	if strings.Contains(rt, "HURDLE") {
		raceType.Hurdle = 1
	}

	if strings.Contains(rt, "HANDICAP") {
		raceType.Handicap = 1
	}

	if strings.Contains(rt, "STAKES") {
		raceType.Stakes = 1
	}

	if strings.Contains(rt, "CHASE") {
		raceType.Chase = 1
	}

	if strings.Contains(rt, "CLASSIFIED") {
		raceType.Classified = 1
	}

	if strings.Contains(rt, "AUCTION") {
		raceType.Auction = 1
	}

	if strings.Contains(rt, "SALES") {
		raceType.Sales = 1
	}

	if strings.Contains(rt, "MAIDEN") {
		raceType.Maiden = 1
	}

	if strings.Contains(rt, "NOVICE") {
		raceType.Novice = 1
	}

	if strings.Contains(rt, "SELLER") {
		raceType.Seller = 1
	}

	if strings.Contains(rt, "CLAIMER") {
		raceType.Claimer = 1
	}

	if strings.Contains(rt, "PATTERN") {
		raceType.Pattern = 1
	}

	if strings.Contains(rt, "BLACK") {
		raceType.Black = 1
	}

	if strings.Contains(rt, "CLAIMING") {
		raceType.Claiming = 1
	}

	if strings.Contains(rt, "CONDITIONAL") {
		raceType.Conditional = 1
	}

	if strings.Contains(rt, "AMATEUR") {
		raceType.Amateur = 1
	}

	return raceType
}

func raceTypeToString(raceType RaceType) string {

	rt := ""

	if raceType.Bumper == 1 {
		rt += " BUMPER"
	}

	if raceType.Hurdle == 1 {
		rt += " HURDLE"
	}

	if raceType.Handicap == 1 {
		rt += " HANDICAP"
	}

	if raceType.Stakes == 1 {
		rt += " STAKES"
	}

	if raceType.Chase == 1 {
		rt += " CHASE"
	}

	if raceType.Classified == 1 {
		rt += " CLASSIFIED"
	}

	if raceType.Auction == 1 {
		rt += " AUCTION"
	}

	if raceType.Sales == 1 {
		rt += " SALES"
	}

	if raceType.Maiden == 1 {
		rt += " MAIDEN"
	}

	if raceType.Novice == 1 {
		rt += " NOVICE"
	}

	if raceType.Seller == 1 {
		rt += " SELLER"
	}

	if raceType.Claimer == 1 {
		rt += " CLAIMER"
	}

	if raceType.Pattern == 1 {
		rt += " PATTERN"
	}

	if raceType.Black == 1 {
		rt += " BLACK"
	}

	if raceType.Claiming == 1 {
		rt += " CLAIMING"
	}

	if raceType.Conditional == 1 {
		rt += " CONDITIONAL"
	}

	if raceType.Amateur == 1 {
		rt += " AMATEUR"
	}

	if len(rt) <= 0 {
		rt += "OTHER"
	}

	return strings.TrimSpace(rt)
}

func convertGoingToAbbreviation(going string) (string, bool) {

	if res := strings.SplitAfterN(going, "(", 2); len(res) > 0 {

		going = strings.ToTitle(strings.ReplaceAll(strings.TrimSpace(strings.ReplaceAll(res[0], "(", "")), " ", "_"))

		_going := ""

		switch strings.ToTitle(strings.ReplaceAll(going, " ", "_")) {
		case "FAST":
			_going = "FST"
		case "FIRM":
			_going = "FM"
		case "GOOD":
			_going = "GD"
		case "GOOD_TO_FIRM":
			_going = "GD-FM"
		case "GOOD_TO_SOFT":
			_going = "GD-SFT"
		case "GOOD_TO_YIELDING":
			_going = "GD-YLD"
		case "HARD":
			_going = "HD"
		case "HEAVY":
			_going = "HVY"
		case "SLOW":
			_going = "SLW"
		case "SOFT":
			_going = "SFT"
		case "SOFT_TO_HEAVY":
			_going = "SFT-HVY"
		case "STANDARD":
			_going = "STD"
		case "STANDARD_TO_FAST":
			_going = "STD-FST"
		case "STANDARD_TO_SLOW":
			_going = "STD-SLW"
		case "YIELDING":
			_going = "YLD"
		case "YIELDING_TO_SOFT":
			_going = "YLD-SFT"
		default:
			_going = strings.ToTitle(strings.ReplaceAll(going, "/", "-"))
		}

		if len(_going) > 0 {

			return _going, true
		}
	}

	return "", false
}

func nameRacecourse(racecourse string) string {

	racecourse = regexp.MustCompile(`(\(.*\))`).ReplaceAllString(racecourse, "")

	replacer := strings.NewReplacer("PARK", "", "CITY", "", "BRIDGE", "", "ON DEE", "", "ON AVON", "")

	racecourse = strings.ReplaceAll(racecourse, "-", " ")

	racecourse = strings.ReplaceAll(strings.TrimSpace(replacer.Replace(racecourse)), " ", "-")

	return racecourse
}

func raceClassFromTitle(title string) int64 {

	if r := regexp.MustCompile(`\(\s*([0-9])\s*\).*$`).FindStringSubmatch(title); len(r) > 0 {

		if i, err := strconv.ParseInt(r[1], 10, 64); err == nil {

			return i
		}
	}

	return 0
}

func winnerAndCurrency(winner string) (string, float64, bool) {

	replacer := strings.NewReplacer(" ", "", ",", "", "£", "", "€", "")

	currency := ""

	// winner = strings.ReplaceAll(winner, ",", "")

	if strings.Contains(winner, "£") {

		currency = "GBP"

	} else if strings.Contains(winner, "€") {
		currency = "EUR"
	}

	if f, err := strconv.ParseFloat(replacer.Replace(winner), 64); err == nil && len(currency) > 2 {

		return currency, f, true
	}

	return currency, -1, false
}

func distanceToMeters(distance string) (float64, error) {

	meters := float64(0.0)

	distance = strings.ToUpper(strings.ReplaceAll(distance, " ", ""))

	switch distance {
	case "5F":
		meters = 1005.852231163131
	case "5½F":
		meters = 1106.4374542794442
	case "6F":
		meters = 1207.0226773957572
	case "6½F":
		meters = 1307.6079005120703
	case "7F":
		meters = 1408.1931236283835
	case "7½F":
		meters = 1508.7783467446966
	case "1M":
		meters = 1609.3635698610096
	case "1M½F":
		meters = 1709.9487929773227
	case "1M1F":
		meters = 1810.534016093636
	case "1M1½F":
		meters = 1911.119239209949
	case "1M2F":
		meters = 2011.704462326262
	case "1M2½F":
		meters = 2112.2896854425753
	case "1M3F":
		meters = 2212.8749085588884
	case "1M3½F":
		meters = 2313.4601316752014
	case "1M4F":
		meters = 2414.0453547915145
	case "1M4½F":
		meters = 2514.6305779078275
	case "1M5F":
		meters = 2615.2158010241405
	case "1M5½F":
		meters = 2715.8010241404536
	case "1M6F":
		meters = 2816.386247256767
	case "1M6½F":
		meters = 2916.97147037308
	case "1M7F":
		meters = 3017.556693489393
	case "1M7½F":
		meters = 3118.1419166057062
	case "2M":
		meters = 3218.7271397220193
	case "2M½F":
		meters = 3319.3123628383323
	case "2M1F":
		meters = 3419.8975859546454
	case "2M1½F":
		meters = 3520.4828090709584
	case "2M2F":
		meters = 3621.068032187272
	case "2M2½F":
		meters = 3721.653255303585
	case "2M3F":
		meters = 3822.238478419898
	case "2M3½F":
		meters = 3922.823701536211
	case "2M4F":
		meters = 4023.408924652524
	case "2M4½F":
		meters = 4123.994147768837
	case "2M5F":
		meters = 4224.579370885151
	case "2M5½F":
		meters = 4325.164594001463
	case "2M6F":
		meters = 4425.749817117777
	case "2M6½F":
		meters = 4526.335040234089
	case "2M7F":
		meters = 4626.920263350403
	case "2M7½F":
		meters = 4727.505486466715
	case "3M":
		meters = 4828.090709583029
	case "3M½F":
		meters = 4928.675932699342
	case "3M1F":
		meters = 5029.261155815655
	case "3M1½F":
		meters = 5129.8463789319685
	case "3M2F":
		meters = 5230.431602048281
	case "3M2½F":
		meters = 5331.016825164595
	case "3M3F":
		meters = 5431.602048280907
	case "3M3½F":
		meters = 5532.187271397221
	case "3M4F":
		meters = 5632.772494513534
	case "3M4½F":
		meters = 5733.357717629847
	case "3M5F":
		meters = 5833.94294074616
	case "3M5½F":
		meters = 5934.528163862473
	case "3M6F":
		meters = 6035.113386978786
	case "3M6½F":
		meters = 6135.698610095099
	case "3M7F":
		meters = 6236.2838332114125
	case "3M7½F":
		meters = 6336.869056327725
	case "4M":
		meters = 6437.454279444039
	case "4M½F":
		meters = 6538.039502560352
	case "4M1F":
		meters = 6638.624725676665
	case "4M1½F":
		meters = 6739.209948792978
	case "4M2F":
		meters = 6839.795171909291
	case "4M2½F":
		meters = 6940.380395025604
	case "4M3F":
		meters = 7040.965618141917
	case "4M3½F":
		meters = 7141.55084125823
	case "4M4F":
		meters = 7242.136064374544
	default:
		meters = -1
	}

	if meters < 0 {

		meters = 0

		m := regexp.MustCompile(`[0-9]+M`).FindStringIndex(distance)
		f := regexp.MustCompile(`[0-9]+F`).FindStringIndex(distance)
		y := regexp.MustCompile(`[0-9]+Y`).FindStringIndex(distance)

		if len(m) > 0 {

			d, err := strconv.ParseFloat(distance[m[0]:m[1]-1], 64)

			if err == nil && d > 0 {

				meters += d / 0.00062137
			}
		}

		if len(f) > 0 {

			d, err := strconv.ParseFloat(distance[f[0]:f[1]-1], 64)

			if err == nil && d > 0 {
				meters += d / 0.0049710
			}
		}

		if len(y) > 0 {

			d, err := strconv.ParseFloat(distance[y[0]:y[1]-1], 64)

			if err == nil && d > 0 {
				meters += d / 1.0936
			}
		}
	}

	if meters <= 0 {

		return -1, errors.New("failed convert distance to meters")
	}

	return meters, nil
}

func wgtToKg(wgt string) float64 {

	kg := float64(0.0)

	stonePound := strings.SplitAfterN(wgt, "-", 2)

	if len(stonePound) > 1 {

		if res, err := strconv.ParseFloat(stonePound[0][:len(stonePound[0])-1], 64); err == nil && res > 0 {

			kg += res / 0.15747
		}

		if res, err := strconv.ParseFloat(stonePound[1], 64); err == nil && res > 0 {

			kg += res / 2.2046
		}
	}

	return kg
}
