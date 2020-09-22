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
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

var client = &http.Client{Transport: transport, Timeout: time.Minute * 5}

func get(path, proxy string) *goquery.Document {

	sl := "\n"
	sl += "get tfh \n"
	sl += "path: " + path + " \n"
	sl += "proxy: " + proxy + " \n"
	log.Println(sl)

	transport.Proxy = http.ProxyURL(&url.URL{Host: proxy})

	host := "www.timeform.com"

	url := "https://" + host + path

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Host", host)
	req.Header.Add("Accept-Languag", "en-gb")
	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15")

	if err != nil {

		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {

		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {

		log.Fatal(err)
	}

	return doc
}
