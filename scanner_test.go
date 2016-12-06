package andrea_test

import (
	"github.com/CarolinaRomero33/Comperas"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok andrea.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: andrea.EOF},
		{s: `#`, tok: andrea.ILLEGAL, lit: `#`},
		{s: ` `, tok: andrea.WS, lit: " "},
		{s: "\t", tok: andrea.WS, lit: "\t"},
		//{s: "\n", tok: andrea.WS, lit: "\n"},
		{s: "\r", tok: andrea.WS, lit: "\r"},

		// Misc characters
		{s: `*`, tok: andrea.ASTERISK, lit: "*"},
		{s: `:`, tok: andrea.DosPunt, lit: ":"},
		{s: `(`, tok: andrea.ParDer, lit: "("},
		{s: `)`, tok: andrea.ParIsq, lit: ")"},
		{s: `;`, tok: andrea.PuntCom, lit: ";"},
		{s: `=`, tok: andrea.Asignaicon, lit: "="},
		//{s: `[`, tok: andrea.AgruppadorDER, lit: "]"},
		//{s: `[`, tok: andrea.AgruppadorISQ, lit: "]"},
		//{s: `\n`, tok: andrea.SaltoDeL, lit: "\n"},

		// Identifiers
		{s: `foo`, tok: andrea.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: andrea.IDENT, lit: `Zx12_3U_`},

		//Keywords
		{s: `si`, tok: andrea.si, lit: "si"},
		{s: `zcase`, tok: andrea.zcase, lit: "zcase"},
		{s: `zmain`, tok: andrea.zmain, lit: "zmain"},
		{s: `no`, tok: andrea.no, lit: "no"},
		{s: `sw`, tok: andrea.sw, lit: "sw"},
		{s: `funcn`, tok: andrea.funcn, lit: "funcn"},
		{s: `zINT`, tok: andrea.zINT, lit: "zINT"},
		{s: `zend`, tok: andrea.zend, lit: "zend"},
		{s: `var`, tok: andrea.var, lit: "var"},
		{s: `while`, tok: andrea.while, lit: "while"},
		{s: `du`, tok: andrea.du, lit: "du"},
		{s: `tri`, tok: andrea.tri, lit: "tri"},
		{s: `cach`, tok: andrea.cach, lit: "cach"},
		{s: `zFINALLY`, tok: andrea.zFINALLY, lit: "zFINALLY"},
		{s: `zsize`, tok: andrea.zsize, lit: "zsize"},
		{s: `zout`, tok: andrea.zout, lit: "zout"},
		{s: `zUNCHECKED`, tok: andrea.zUNCHECKED, lit: "zUNCHECKED"},
		{s: `znew`, tok: andrea.znew, lit: "znew"},
		{s: `feach`, tok: andrea.feach, lit: "feach"},
		{s: `zIMPORT`, tok: andrea.zIMPORT, lit: "zIMPORT"},


		//Bonifaz
		//{s: `MCCREATE`, tok: andrea.MCCREATE, lit: "MCCREATE"},
		//
		//{s: `MCUSE`, tok: andrea.MCUSE, lit: "MCUSE"},
		//
		//{s: `MCCREATETABLE`, tok: andrea.MCCREATETABLE, lit: "MCCREATETABLE"},
		//
		//{s: `DATABASE`, tok: andrea.DATABASE, lit: "DATABASE"},




	}

	for i, tt := range tests {
		s := andrea.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q zt=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q zt=%q", i, tt.s, tt.lit, lit)
		}
	}
}
