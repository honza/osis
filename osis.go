// Copyright 2019, Honza Pokorny <me@honza.ca>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package osis

import (
	"errors"
	"strings"
)

var singleChapterBooks = map[string]int{
	"Obad":  0,
	"Phlm":  0,
	"2John": 0,
	"3John": 0,
	"Jude":  0,
}

var books = map[string][]string{
	"Gen":           []string{"Genesis"},
	"Exod":          []string{"Exodus"},
	"Lev":           []string{"Leviticus"},
	"Num":           []string{"Numbers"},
	"Deut":          []string{"Deuteronomy"},
	"Josh":          []string{"Joshua"},
	"Judg":          []string{"Judges"},
	"Ruth":          []string{"Ruth"},
	"1Sam":          []string{"1 Samuel"},
	"2Sam":          []string{"2 Samuel"},
	"1Kgs":          []string{"1 Kings"},
	"2Kgs":          []string{"2 Kings"},
	"1Chr":          []string{"1 Chronicles"},
	"2Chr":          []string{"2 Chronicles"},
	"Ezra":          []string{"Ezra"},
	"Neh":           []string{"Nehemiah"},
	"Esth":          []string{"Esther"},
	"Job":           []string{"Job"},
	"Ps":            []string{"Psalm", "Psalms"},
	"Prov":          []string{"Proverbs"},
	"Eccl":          []string{"Ecclesiastes"},
	"Song":          []string{"Song of Solomon"},
	"Isa":           []string{"Isaiah"},
	"Jer":           []string{"Jeremiah"},
	"Lam":           []string{"Lamentations"},
	"Ezek":          []string{"Ezekiel"},
	"Dan":           []string{"Daniel"},
	"Hos":           []string{"Hosea"},
	"Joel":          []string{"Joel"},
	"Amos":          []string{"Amos"},
	"Obad":          []string{"Obadiah"},
	"Jonah":         []string{"Jonah"},
	"Mic":           []string{"Micah"},
	"Nah":           []string{"Nahum"},
	"Hab":           []string{"Habakkuk"},
	"Zeph":          []string{"Zephaniah"},
	"Hag":           []string{"Haggai"},
	"Zech":          []string{"Zechariah"},
	"Mal":           []string{"Malachi"},
	"Matt":          []string{"Matthew"},
	"Mark":          []string{"Mark"},
	"Luke":          []string{"Luke"},
	"John":          []string{"John"},
	"Acts":          []string{"Acts"},
	"Rom":           []string{"Romans"},
	"1Cor":          []string{"1 Corinthians"},
	"2Cor":          []string{"2 Corinthians"},
	"Gal":           []string{"Galatians"},
	"Eph":           []string{"Ephesians"},
	"Phil":          []string{"Philippians"},
	"Col":           []string{"Colossians"},
	"1Thess":        []string{"1 Thessalonians"},
	"2Thess":        []string{"2 Thessalonians"},
	"1Tim":          []string{"1 Timothy"},
	"2Tim":          []string{"2 Timothy"},
	"Titus":         []string{"Titus"},
	"Phlm":          []string{"Philemon"},
	"Heb":           []string{"Hebrews"},
	"Jas":           []string{"James"},
	"1Pet":          []string{"1 Peter"},
	"2Pet":          []string{"2 Peter"},
	"1John":         []string{"1 John"},
	"2John":         []string{"2 John"},
	"3John":         []string{"3 John"},
	"Jude":          []string{"Jude"},
	"Rev":           []string{"Revelation"},
	"Ps.$chapters":  []string{"Psalm", "Psalms"},
	"1Sam-2Sam":     []string{"1—2 Samuel"},
	"1Kgs-2Kgs":     []string{"1—2 Kings"},
	"1Chr-2Chr":     []string{"1—2 Chronicles"},
	"1Cor-2Cor":     []string{"1—2 Corinthians"},
	"1Thess-2Thess": []string{"1—2 Thessalonians"},
	"1Tim-2Tim":     []string{"1—2 Timothy"},
	"1Pet-2Pet":     []string{"1—2 Peter"},
	"1John-2John":   []string{"1—2 John"},
	"1John-3John":   []string{"1—3 John"},
	"2John-3John":   []string{"2—3 John"},
	"1Sam,2Sam":     []string{"1 and 2 Samuel"},
	"1Kgs,2Kgs":     []string{"1 and 2 Kings"},
	"1Chr,2Chr":     []string{"1 and 2 Chronicles"},
	"1Cor,2Cor":     []string{"1 and 2 Corinthians"},
	"1Thess,2Thess": []string{"1 and 2 Thessalonians"},
	"1Tim,2Tim":     []string{"1 and 2 Timothy"},
	"1Pet,2Pet":     []string{"1 and 2 Peter"},
	"1John,2John":   []string{"1 and 2 John"},
	"1John,3John":   []string{"1 and 3 John"},
	"2John,3John":   []string{"2 and 3 John"},
}

func isSingleChapterBook(book string) bool {
	_, ok := singleChapterBooks[book]
	return ok
}

func Format(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	result := ""

	if strings.Contains(s, "-") {
		groups := strings.Split(s, "-")
		if len(groups) != 2 {
			return "", errors.New("invalid range")
		}

		start := groups[0]
		end := groups[1]

		startParts := strings.Split(start, ".")
		endParts := strings.Split(end, ".")

		startBookList, ok := books[startParts[0]]

		if !ok {
			return "", errors.New("unknown book")
		}

		endBookList, ok := books[endParts[0]]

		if !ok {
			return "", errors.New("unknown book")
		}

		startBook := startBookList[0]
		endBook := endBookList[0]

		if len(startBookList) == 2 {
			startBook = startBookList[1]
		}

		if len(endBookList) == 2 {
			endBook = endBookList[1]
		}

		// If we're dealing with a range inside one chapter in the
		// Psalms, revert back.
		if startBook == endBook && startParts[1] == endParts[1] {
			startBook = startBookList[0]
			endBook = endBookList[0]
		}

		result += startBook
		result += " "

		if len(startParts) >= 2 {
			result += startParts[1]
		}

		if len(startParts) == 3 {
			result += ":"
			result += startParts[2]
		}

		result += "-"

		if startBook == endBook && startParts[1] == endParts[1] {
			result += endParts[2]
			return result, nil
		}

		if startBook != endBook {
			result += endBook
			result += " "
		}

		if len(endParts) >= 2 {
			result += endParts[1]
		}

		if len(endParts) == 3 {
			result += ":"
			result += endParts[2]
		}

		return result, nil
	}

	parts := strings.Split(s, ".")

	if parts == nil {
		return "", errors.New("Malformed")
	}

	bookList, ok := books[parts[0]]

	if !ok {
		return "", errors.New("unknown book")
	}

	book := bookList[0]

	result += book

	if len(parts) > 1 {
		result += " "
	}

	if len(parts) == 2 {
		if !isSingleChapterBook(book) {
			chapter := parts[1]
			result += chapter
		}
	}

	if len(parts) == 3 {
		if !isSingleChapterBook(book) {
			chapter := parts[1]
			result += chapter
			result += ":"
		}
		verse := parts[2]
		result += verse
	}

	return result, nil
}

func FormatMany(s string) ([]string, error) {
	groups := strings.Split(s, ",")
	result := []string{}

	for _, group := range groups {
		formatted, err := Format(group)

		if err != nil {
			return []string{}, err
		}

		result = append(result, formatted)
	}

	return result, nil
}
