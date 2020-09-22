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

// GetResultByPath parsing the results page. Example: /horse-racing/result/chelmsford-city/2020-09-20/1300/28/1
func GetResultByPath(path string, proxy string) *Race {

	var datetime *time.Time = nil
	var types *RaceType = nil

	distance := float64(-1)
	winner := float64(-1)
	currency := ""
	racecourse := ""
	tRT := ""
	going := ""

	replacer := strings.NewReplacer("(", "", ")", "")

	doc := get(path, proxy)

	if res, err := time.Parse("15:04 Monday 02 January 2006Z-0700", strings.TrimSpace(doc.Find("h2.rp-header-text[title='Date and time of race']").First().Text())+"Z+0100"); err == nil {

		t := res.In(time.UTC)

		datetime = &t
	}

	title := strings.TrimSpace(doc.Find("h3.rp-header-text[title='Race title']").First().Text())

	category := strings.ToUpper(strings.TrimSpace(doc.Find("span[title='The type of race']").First().Text()))
	surface := strings.ToUpper(strings.TrimSpace(doc.Find("span[title='Surface of the race']").First().Text()))

	tRCN := strings.TrimSpace(doc.Find("h1[title='Racecourse name']").First().Text())

	class := raceClassFromTitle(title)

	rt := raceTypeFromTitle(title)

	types = &rt

	tRT = raceTypeToString(rt)

	if c, w, err := winnerAndCurrency(strings.TrimSpace(doc.Find("span[title='Prize money to winner']").First().Text())); err {

		currency = c
		winner = w

	}

	if res, err := distanceToMeters(doc.Find("span[title='Distance expressed in miles, furlongs and yards']").Text()); err == nil {

		distance = res
	}

	if r := regexp.MustCompile(`[0-9]+:[0-9]+$`).FindStringIndex(tRCN); len(r) > 0 {

		racecourse = nameRacecourse(strings.ToUpper(strings.TrimSpace(tRCN[:r[0]])))
	}

	if g, err := convertGoingToAbbreviation(strings.ToUpper(strings.TrimSpace(doc.Find("span[title='Race going']").First().Text()))); err {
		going = g
	}

	runners := make(map[int64]Runner, 0)

	doc.Find("table tbody.rp-table-row").Each(func(i int, s *goquery.Selection) {

		isp := float64(-1)
		num := int64(-1)
		draw := int64(-1)
		age := int64(-1)
		rating := int64(-1)
		horseID := int64(-1)
		jockeyID := int64(-1)
		trainerID := int64(-1)

		pos := strings.ToUpper(strings.TrimSpace(s.Find("span.rp-entry-number").First().Text()))
		horse := ""

		sHORSE := s.Find("a.rp-horse").First()
		tHORSE := sHORSE.Text()

		sJOCKEY := s.Find("a[title='Jockey']").First()
		jockey := strings.ToUpper(strings.TrimSpace(sJOCKEY.Text()))

		sTRAINER := s.Find("a[title='Trainer']").First()
		trainer := strings.ToUpper(strings.TrimSpace(sTRAINER.Text()))

		if i, err := strconv.ParseInt(strings.TrimSpace(replacer.Replace(s.Find("span.rp-draw").First().Text())), 10, 64); err == nil {
			draw = i
		}

		if h, err := sHORSE.Attr("href"); err {

			if r := regexp.MustCompile(`\/[0-9]+$`).FindStringIndex(h); len(r) > 0 {

				if i, err := strconv.ParseInt(h[r[0]+1:r[1]], 10, 64); err == nil {

					horseID = i
				}
			}
		}

		if r := regexp.MustCompile(`^\s*[0-9]+`).FindStringIndex(tHORSE); len(r) > 0 {

			horse = strings.TrimSpace(tHORSE[r[1]+1:])

			if i, err := strconv.ParseInt(tHORSE[r[0]:r[1]], 10, 64); err == nil {

				num = i
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

		if res, err := strconv.ParseInt(strings.TrimSpace(replacer.Replace(s.Find("td.rp-ageequip-hide[title='Official rating given to this horse']").Text())), 10, 64); err == nil {

			rating = res
		}

		if res, err := strconv.ParseFloat(strings.TrimSpace(s.Find("td.rp-result-sp span.price-decimal").First().Text()), 64); err == nil {

			isp = res
		}

		wgt := wgtToKg(s.Find("tr:nth-child(1) > td:nth-child(13)").Text())

		if isp > 1 && wgt > 0 && num > 0 && age > 0 && horseID > 0 && jockeyID > 0 && trainerID > 0 && len(pos) > 0 && len(pos) > 0 && len(horse) > 0 && len(jockey) > 0 && len(trainer) > 0 {

			runners[num] = Runner{
				Isp:       isp,
				Wgt:       wgt,
				Number:    num,
				Draw:      draw,
				Age:       age,
				Rating:    rating,
				HorseID:   horseID,
				JockeyID:  jockeyID,
				TrainerID: trainerID,
				Pos:       pos,
				Horse:     horse,
				Jockey:    jockey,
				Trainer:   trainer,
			}
		}
	})

	if datetime != nil && types != nil && distance > 0 && len(title) > 0 && len(currency) > 0 && winner > 0 && len(tRT) > 0 && len(category) > 0 && len(racecourse) > 0 && len(going) > 0 && len(runners) > 0 {

		race := Race{
			Datetime:   *datetime,
			Distance:   distance,
			Class:      class,
			Title:      title,
			Currency:   currency,
			Winner:     winner,
			RaceType:   tRT,
			Category:   category,
			Surface:    surface,
			Racecourse: racecourse,
			Going:      going,
			Types:      *types,
			Runners:    runners,
		}

		return &race
	}

	return nil
}

// GetResultByDate parsing the results for the whole day.
func GetResultByDate(date time.Time, proxy string) []Race {

	races := make([]Race, 0)

	doc := get("/horse-racing/results/latestresultsarchive/"+date.In(time.UTC).Format("2006-01-02")+"?region=0", proxy)

	doc.Find("div.w-results-holder section a.results-title").Each(func(i int, s *goquery.Selection) {

		if res, err := s.Attr("href"); err {

			if result := GetResultByPath(res, proxy); result != nil {

				races = append(races, *result)

				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	})

	return races
}
