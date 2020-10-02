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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// GetRacecardByPath parsing the racecard page. Example: /horse-racing/racecards/fontwell-park/2020-10-03/1419/19/1/bigmore-associates-mortgages-conditional-jockeys-handicap-hurdle
func GetRacecardByPath(path string, proxy string) *Racecard {

	replacer := strings.NewReplacer("(", "", ")", "")

	var datetime *time.Time = nil
	var types *RaceType = nil

	racecourse := ""
	distance := float64(-1)
	winner := float64(-1)
	currency := ""
	tRT := ""
	going := ""
	entries := make(map[int64]Entry, 0)

	doc := get(path, proxy)

	if res, err := time.Parse("15:04 Mon 02 January 2006Z-0700", strings.TrimSpace(doc.Find("h2.rp-header-text[title='Date and time of race']").First().Text())+"Z+0100"); err == nil {

		t := res.In(time.UTC)

		datetime = &t
	}

	title := strings.TrimSpace(doc.Find("h3.rp-header-text[title='Race title']").First().Text())

	tRCN := strings.TrimSpace(doc.Find("h1[title='Racecourse name']").First().Text())

	if r := regexp.MustCompile(`[0-9]+:[0-9]+$`).FindStringIndex(tRCN); len(r) > 0 {

		racecourse = nameRacecourse(strings.ToUpper(strings.TrimSpace(tRCN[:r[0]])))
	}

	if res, err := distanceToMeters(doc.Find("span[title='Distance expressed in miles, furlongs and yards']").Text()); err == nil {

		distance = res
	}

	if c, w, err := winnerAndCurrency(strings.TrimSpace(doc.Find("span[title='Prize money to winner']").First().Text())); err {

		currency = c
		winner = w
	}

	if g, err := convertGoingToAbbreviation(strings.ToUpper(strings.TrimSpace(doc.Find("td.rp-goingweather-td span.rp-header-text").First().Text()))); err {
		going = g
	}

	rt := raceTypeFromTitle(title)

	types = &rt

	tRT = raceTypeToString(rt)

	class := raceClassFromTitle(title)

	surface := strings.ToUpper(strings.TrimSpace(doc.Find("span[title='Surface of the race']").First().Text()))

	doc.Find("table tbody.rp-table-row").Each(func(i int, s *goquery.Selection) {

		num := int64(-1)
		draw := int64(-1)
		age := int64(-1)
		rating := int64(-1)
		horseID := int64(-1)
		jockeyID := int64(-1)
		trainerID := int64(-1)

		sHORSE := s.Find("a.rp-horse").First()
		horse := strings.ToUpper(strings.TrimSpace(sHORSE.Text()))

		sJOCKEY := s.Find("a[title='Jockey']").First()
		jockey := strings.ToUpper(strings.TrimSpace(sJOCKEY.Text()))

		sTRAINER := s.Find("a[title='Trainer']").First()
		trainer := strings.ToUpper(strings.TrimSpace(sTRAINER.Text()))

		if i, err := strconv.ParseInt(strings.TrimSpace(s.Find("span.rp-entry-number").First().Text()), 10, 64); err == nil {

			num = i
		}

		if i, err := strconv.ParseInt(strings.TrimSpace(replacer.Replace(s.Find("span.rp-draw").First().Text())), 10, 64); err == nil {

			draw = i
		}

		if res, err := strconv.ParseInt(strings.TrimSpace(replacer.Replace(s.Find("span[title='Official rating given to this horse']").First().Text())), 10, 64); err == nil {

			rating = res
		}

		if h, err := sHORSE.Attr("href"); err {

			if r := regexp.MustCompile(`^\/horse-racing\/horse\/form\/[a-z\_\-]+\/[0-9]+\/`).FindStringIndex(h); len(r) > 0 {

				trimH := h[r[0]:r[1]]

				if r := regexp.MustCompile(`\/[0-9]+\/$`).FindStringIndex(trimH); len(r) > 0 {

					if i, err := strconv.ParseInt(trimH[r[0]+1:r[1]-1], 10, 64); err == nil {

						horseID = i
					}
				}
			}
		}

		if h, err := sJOCKEY.Attr("href"); err {

			if r := regexp.MustCompile(`\/[0-9]+$`).FindStringIndex(h); len(r) > 0 {

				if i, err := strconv.ParseInt(h[r[0]+1:r[1]], 10, 64); err == nil {

					jockeyID = i
				}
			}
		}

		if h, err := sTRAINER.Attr("href"); err {

			if r := regexp.MustCompile(`\/[0-9]+$`).FindStringIndex(h); len(r) > 0 {

				if i, err := strconv.ParseInt(h[r[0]+1:r[1]], 10, 64); err == nil {

					trainerID = i
				}
			}
		}
		if res, err := strconv.ParseInt(strings.TrimSpace(s.Find("td[title='Horse age']").Text()), 10, 64); err == nil {
			age = res
		}

		wgt := wgtToKg(s.Find("td[title='Weight the horse is carrying']").Text())

		if wgt > 0 && num > 0 && age > 0 && horseID > 0 && jockeyID > 0 && trainerID > 0 && len(horse) > 0 && len(jockey) > 0 && len(trainer) > 0 {

			entries[num] = Entry{
				Wgt:       wgt,
				Number:    num,
				Draw:      draw,
				Age:       age,
				Rating:    rating,
				HorseID:   horseID,
				JockeyID:  jockeyID,
				TrainerID: trainerID,
				Horse:     horse,
				Jockey:    jockey,
				Trainer:   trainer,
			}
		}

	})

	if datetime != nil && types != nil && distance > 0 && len(title) > 0 && len(currency) > 0 && winner > 0 && len(tRT) > 0 && len(racecourse) > 0 && len(going) > 0 && len(entries) > 0 {

		racecard := Racecard{
			Datetime:   *datetime,
			Distance:   distance,
			Class:      class,
			Title:      title,
			Currency:   currency,
			Winner:     winner,
			RaceType:   tRT,
			Surface:    surface,
			Racecourse: racecourse,
			Going:      going,
			Types:      *types,
			Entries:    entries,
		}

		return &racecard
	}

	return nil
}
