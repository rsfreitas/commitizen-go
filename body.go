package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// capitalize capitalizes the first letter of a phrase.
func capitalize(phrase string) string {
	a := []rune(phrase)
	a[0] = unicode.ToUpper(a[0])

	return string(a)
}

// getSentences splits a text into sentences, lines. Each sentence will have
// a hyphen as prefix.
func getSentences(body string) []string {
	s := strings.Split(body, "\n")
	var sentences []string

	for _, r := range s {
		if r != "" {
			line := fmt.Sprintf("- %s", capitalize(r))
			sentences = append(sentences, line)
		}
	}

	return sentences
}

type AlignedSentence struct {
	max      int
	sentence string
	offset   int
	buf      bytes.Buffer
}

func (a *AlignedSentence) align(s string, offset int) int {
	if len(s) > a.max {
		tmp := s

		for offset < len(s) {
			chunk := tmp[:strings.LastIndex(tmp[:a.max], " ")]
			tmp = strings.TrimSpace(tmp[strings.LastIndex(tmp[:a.max], " "):])

			line := chunk
			if !strings.HasPrefix(line, "-") {
				line = fmt.Sprintf("  %s", chunk)
			}

			offset += len(line) + 1
			a.buf.WriteString(line + "\n")
			offset = a.align(tmp, offset)
		}
	} else {
		line := s
		if !strings.HasPrefix(line, "-") {
			line = fmt.Sprintf("  %s", s)
		}

		offset += len(line)
		a.buf.WriteString(line)
	}

	return offset
}

func (a *AlignedSentence) alignSentence() string {
	a.align(a.sentence, a.offset)
	return a.buf.String()
}

func alignSentence(sentence string, max int) string {
	if len(sentence) <= max {
		return sentence
	}

	a := &AlignedSentence{
		sentence: sentence,
		max:      max,
	}

	return a.alignSentence()
}

// bodyToBulletPoints converts a text body into bullet point sentences.
func bodyToBulletPoints(body string) string {
	text := bytes.NewBufferString("\n")
	sentences := getSentences(body)

	for _, s := range sentences {
		text.WriteString(alignSentence(s, 72))
		text.WriteString("\n\n")
	}

	return text.String()
}
