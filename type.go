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

import "time"

// Runner information about runner.
type Runner struct {
	Isp       float64
	Wgt       float64
	Number    int64
	Draw      int64
	Age       int64
	Rating    int64
	HorseID   int64
	JockeyID  int64
	TrainerID int64
	Pos       string
	Horse     string
	Jockey    string
	Trainer   string
}

// Race information about race.
type Race struct {
	Datetime   time.Time
	Distance   float64
	Class      int64
	Title      string
	Currency   string
	Winner     float64
	RaceType   string
	Category   string
	Surface    string
	Racecourse string
	Going      string
	Types      RaceType
	Runners    map[int64]Runner
}

// RaceType from title
type RaceType struct {
	Bumper      int64
	Hurdle      int64
	Handicap    int64
	Stakes      int64
	Chase       int64
	Classified  int64
	Auction     int64
	Sales       int64
	Maiden      int64
	Novice      int64
	Seller      int64
	Claimer     int64
	Pattern     int64
	Black       int64
	Claiming    int64
	Conditional int64
	Amateur     int64
}
