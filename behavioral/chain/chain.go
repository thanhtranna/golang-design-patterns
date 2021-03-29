package main

import (
	"fmt"
	"strconv"
)

type TranslatorRing interface {
	Next(TranslatorRing)
	GetNext() TranslatorRing
	KnowsWord(s string) bool
	TranslationOf(s string) string
}

type EnglishTranslator struct {
	next TranslatorRing
}

func (fr EnglishTranslator) Next(r TranslatorRing) {
	fr.next = r
}

func (fr EnglishTranslator) GetNext() TranslatorRing {
	return fr.next
}

func (fr EnglishTranslator) KnowsWord(s string) bool {
	dict := ItalianToEnglishDictionary()
	if _, ok := dict[s]; ok {
		return true
	}
	return false
}

func (fr EnglishTranslator) TranslationOf(s string) string {
	dict := ItalianToEnglishDictionary()
	if val, ok := dict[s]; ok {
		return val
	}
	panic("Oops! Wrong translation requested")
}

func ItalianToEnglishDictionary() map[string]string {
	dict := map[string]string{}
	dict["topo"] = "mouse"
	dict["cocomero"] = "watermelon"
	return dict
}

type FrenchTranslator struct {
	next TranslatorRing
}

func (sr *FrenchTranslator) Next(r TranslatorRing) {
	sr.next = r
}

func (fr *FrenchTranslator) GetNext() TranslatorRing {
	return fr.next
}
func (fr *FrenchTranslator) KnowsWord(s string) bool {
	dict := ItalianToFrancaisDictionary()
	if _, ok := dict[s]; ok {
		return true
	}
	return false
}

func (fr *FrenchTranslator) TranslationOf(s string) string {
	dict := ItalianToFrancaisDictionary()
	if val, ok := dict[s]; ok {
		return val
	}
	panic("Oops! Wrong translation requested")
}

func ItalianToFrancaisDictionary() map[string]string {
	dict := map[string]string{}
	dict["casa"] = "maison"
	return dict
}

type SpanishTranslator struct {
	next TranslatorRing
}

func (sr *SpanishTranslator) Next(r TranslatorRing) {
	sr.next = r
}

func (fr *SpanishTranslator) GetNext() TranslatorRing {
	return fr.next
}
func (fr *SpanishTranslator) KnowsWord(s string) bool {
	dict := ItalianToSpanishDictionary()
	if _, ok := dict[s]; ok {
		return true
	}
	panic("Oops! Wrong translation requested")
}

func (fr *SpanishTranslator) TranslationOf(s string) string {
	dict := ItalianToSpanishDictionary()
	if val, ok := dict[s]; ok {
		return val
	}
	panic("Oops! Wrong translation requested")
}

func ItalianToSpanishDictionary() map[string]string {
	dict := map[string]string{}
	dict["ciao"] = "hola"
	dict["cocomero"] = "sandía"
	return dict
}

type TranslatorChain struct {
	Start TranslatorRing
}

func (h *TranslatorChain) Translate(s string) (string, error) {
	r := h.Start
	for {
		if r.KnowsWord(s) {
			return r.TranslationOf(s), nil
		}
		if r.GetNext() != nil {
			r = r.GetNext()
		} else {
			panic("Oops! Missing translation")
		}
	}
}

func (h *TranslatorChain) CountRings() int {
	numOfRings := 0
	if h.Start != nil {
		numOfRings++
	}
	r := h.Start
	for r.GetNext() != nil {
		r = r.GetNext()
		numOfRings++
	}
	return numOfRings

}

func main() {
	chain := TranslatorChain{
		Start: &EnglishTranslator{
			&FrenchTranslator{
				&SpanishTranslator{},
			},
		},
	}

	fmt.Println("vim-go")

	fmt.Println(strconv.Itoa(chain.CountRings())) // 3
	fmt.Println(chain.Translate("ciao"))          // spanish: hola
	fmt.Println(chain.Translate("casa"))          // french: maison
	fmt.Println(chain.Translate("topo"))          // english mouse
	fmt.Println(chain.Translate("cocomero"))      // english watermelon
}
