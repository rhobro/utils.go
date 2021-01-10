package httputil

import (
	"fmt"
	"github.com/Bytesimal/goutils/pkg/util"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func RQUntil(cli *http.Client, rq *http.Request, count int) (rsp *http.Response, err error) {
	err = fmt.Errorf("tmp")
	for i := 0; i < count; i++ {
		if err != nil {
			if i > 0 {
				// Random sleep in millisecond
				time.Sleep(time.Duration(util.Rand.Int63n(int64(time.Millisecond))))
			}
			rsp, err = cli.Do(rq)
			continue
		}
		break
	}
	return
}

var macVs []string
var iosVs []string
var iosDevices = []string{"iPod", "iPhone", "iPad"}

var webkitVs = map[string][]string{
	"mac":     {},
	"windows": {},
	"ios":     {},
}

var lazyHeaderInit sync.Once

const std = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15"

func initHeaders() {
	var wg sync.WaitGroup
	var mainErr error

	// scrape macOS versions
	wg.Add(1)
	go func() {
		defer wg.Done()

		rq, _ := http.NewRequest("GET", "https://apple.fandom.com/wiki/List_of_Mac_OS_versions", nil)
		rq.Header.Set("User-Agent", std)
		rsp, err := http.DefaultClient.Do(rq)
		if err != nil {
			if mainErr == nil {
				mainErr = err
			}
			return
		}
		defer rsp.Body.Close()
		page, err := goquery.NewDocumentFromReader(rsp.Body)
		if err != nil {
			if mainErr == nil {
				mainErr = err
			}
			return
		}

		var watchOut bool
		page.Find("div#content > div#mw-content-text > div *").Each(func(_ int, sl *goquery.Selection) {
			if sl.Is("h3") {
				// Check to see if contains version to scrape
				if strings.Contains(sl.Text(), "OS X 10") ||
					strings.Contains(sl.Text(), "macOS 1") {
					watchOut = true
				}
			} else {
				// If to watch out and is ul, parse ul
				if sl.Is("ul") && watchOut {
					sl.Find("li").Each(func(_ int, sl *goquery.Selection) {
						txt := sl.Text()

						// Check if text has valid version and parse it
						for _, e := range strings.Split(txt, " ") {
							_, err := strconv.Atoi(strings.Split(e, ".")[0])
							if strings.Contains(e, ".") && err == nil {
								macVs = append(macVs, e)
								break
							}
						}
					})
				}
			}
		})
	}()

	// scrape iOS versions
	wg.Add(1)
	go func() {
		defer wg.Done()

		rq, _ := http.NewRequest("GET", "https://www.gkgigs.com/list-apple-ios-version-history/", nil)
		rq.Header.Set("User-Agent", std)
		rsp, err := http.DefaultClient.Do(rq)
		if err != nil {
			if mainErr == nil {
				mainErr = err
			}
			return
		}
		defer rsp.Body.Close()
		page, err := goquery.NewDocumentFromReader(rsp.Body)
		if err != nil {
			if mainErr == nil {
				mainErr = err
			}
			return
		}

		page.Find("figure.wp-block-table > table").Each(func(_ int, sl *goquery.Selection) {
			headings := sl.Find("thead > tr > th")
			if headings.Length() > 0 {
				title := sl.Find("thead > tr > th").Get(0).FirstChild.Data

				if title == "Version" {
					sl.Find("tbody > tr > td").Each(func(_ int, sl *goquery.Selection) {
						vsNode := sl.Get(0).FirstChild.FirstChild
						if vsNode != nil {
							iosVs = append(iosVs, vsNode.Data)
						}
					})
				}
			}
		})
	}()

	// scrape webkit versions
	wg.Add(1)
	go func() {
		defer wg.Done()

		rq, _ := http.NewRequest("GET", "https://en.wikipedia.org/wiki/Safari_version_history", nil)
		rq.Header.Set("User-Agent", std)
		rsp, err := http.DefaultClient.Do(rq)
		if err != nil {
			if mainErr == nil {
				mainErr = err
			}
			return
		}
		defer rsp.Body.Close()
		page, err := goquery.NewDocumentFromReader(rsp.Body)
		if err != nil {
			if mainErr == nil {
				mainErr = err
			}
			return
		}

		var started bool
		var osFocus string
		var watchOut bool
		page.Find("div.mw-parser-output * ").Each(func(_ int, sl *goquery.Selection) {
			// Wait for release history
			if started {
				// Found new OS section
				if sl.Is("h3") {
					os := strings.ToLower(sl.Text())
					if strings.Contains(os, "[") {
						os = os[:strings.Index(os, "[")]
					}
					var in bool

					for k := range webkitVs {
						if k == os {
							in = true
						}
					}
					if in {
						osFocus = os
					}
				}
				// Found title before table
				if sl.Is("h4") && strings.Contains(sl.Text(), "Safari") {
					watchOut = true
				} else {
					if watchOut && sl.Is("table") {
						// Found table
						watchOut = false
						sl.Find("tbody > tr").Each(func(i int, sl *goquery.Selection) {
							if i < 2 {
								return
							}

							// Find col with correct version
							sl.Find("td").Each(func(_ int, sl *goquery.Selection) {
								// Check if text has valid version and parse it
								for _, e := range strings.Split(sl.Get(0).FirstChild.Data, " ") {
									v, err := strconv.Atoi(strings.Split(e, ".")[0])

									if strings.Contains(e, ".") && err == nil && v > 40 {
										version := strings.TrimSpace(sl.Text())
										if strings.Contains(version, "[") {
											version = version[:strings.Index(version, "[")]
										}

										// Add to specific os list
										webkitVs[osFocus] = append(webkitVs[osFocus], version)
										break
									}
								}
							})
						})
					}
				}
			} else {
				// Wait for release history
				if sl.Is("h2") && strings.Contains(sl.Text(), "Release history") {
					started = true
				}
			}
		})
	}()

	// scrape safari versions TODO
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	// scrape chrome versions TODO
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	wg.Wait()
	if mainErr != nil {
		log.Printf("unable to scrape header lists: %s", mainErr)
		lazyHeaderInit = sync.Once{}
	}
}

const (
	safari = iota
	chrome
	firefox
)

func RandUA() string {
	lazyHeaderInit.Do(initHeaders)

	ua := "Mozilla/5.0 ("
	switch util.Rand.Intn(3) {
	case 0:
		// macOS
		// bool
		browser := util.Rand.Intn(3) // 0 = safari, 1 = firefox, 2 = chrome

		ua += "Macintosh; Intel Mac OS X "
		vStr := macVs[util.Rand.Intn(len(macVs))]

		// add device vs
		if browser == firefox {
			ua += fmt.Sprintf("%s; rv:70.0", vStr)
		} else {
			ua += strings.ReplaceAll(vStr, ".", "_")
		}
		ua += ") "

		if browser == firefox {
			ua += "Gecko/20100101 Firefox/70.0"
			return ua
		} else {
			// Webkit
			ua += fmt.Sprintf("AppleWebKit/%s  (KHTML, like Gecko) ", webkitVs["mac"][util.Rand.Intn(len(webkitVs["mac"]))])
		}

		// TODO safari and chrome version history
	case 1:
		// Windows
		return std
	case 2:
		// iOS
		ua += fmt.Sprintf("%s; CPU iPhone OS %s like Mac OS X) ",
			iosDevices[util.Rand.Intn(len(iosDevices))],
			strings.ReplaceAll(iosVs[util.Rand.Intn(len(iosVs))], ".", "_"))
	}

	return ua
}
