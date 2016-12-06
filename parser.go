package andrea

import (
	"fmt"
	"io"
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct{}


type zifstatement struct{}


type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}


func (p *Parser) Parse() (*SelectStatement, error) {
	stmt := &SelectStatement{}


	//Erasmo

	if tok, lit := p.scanIgnoreWhitespace(); tok != zIF {
		return nil, fmt.Errorf("encontro %q, esperaba zif", lit)
	}
	return stmt,nil
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != no {
		return nil, fmt.Errorf("encontro %q, esperaba no", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//Suma global.

	if tok, lit := p.scanIgnoreWhitespace(); tok != var {
		return nil, fmt.Errorf("encontro %q, esperaba var", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != var {
		return nil, fmt.Errorf("encontro %q, esperaba var", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != funcn {
		return nil, fmt.Errorf("encontro %q, esperaba funcn", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zmain {
		return nil, fmt.Errorf("encontro %q, esperaba zmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//suma normal

	if tok, lit := p.scanIgnoreWhitespace(); tok != funcn {
		return nil, fmt.Errorf("encontro %q, esperaba funcn", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zmain {
		return nil, fmt.Errorf("encontro %q, esperaba zmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	// CICLO para

	if tok, lit := p.scanIgnoreWhitespace(); tok != para {
		return nil, fmt.Errorf("encontro %q, esperaba para", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba para", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zcase

	if tok, lit := p.scanIgnoreWhitespace(); tok != zcase {
		return nil, fmt.Errorf("encontro %q, esperaba zcase", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//sw

	if tok, lit := p.scanIgnoreWhitespace(); tok != sw {
		return nil, fmt.Errorf("encontro %q, esperaba sw", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != break {
		return nil, fmt.Errorf("encontro %q, esperaba break", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//while

	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != while {
		return nil, fmt.Errorf("encontro %q, esperaba while", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//du ciclo
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != du {
		return nil, fmt.Errorf("encontro %q, esperaba du", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != while {
		return nil, fmt.Errorf("encontro %q, esperaba while", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//tri
	if tok, lit := p.scanIgnoreWhitespace(); tok != tri {
		return nil, fmt.Errorf("encontro %q, esperaba tri", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != cach {
		return nil, fmt.Errorf("encontro %q, esperaba cach", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zfinally

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFINALLY {
		return nil, fmt.Errorf("encontro %q, esperaba zfinally", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zsize

	if tok, lit := p.scanIgnoreWhitespace(); tok != zsize {
		return nil, fmt.Errorf("encontro %q, esperaba zsize", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != znew {
		return nil, fmt.Errorf("encontro %q, esperaba znew", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zout

	if tok, lit := p.scanIgnoreWhitespace(); tok != zout {
		return nil, fmt.Errorf("encontro %q, esperaba zout", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zUNCHECKED

	if tok, lit := p.scanIgnoreWhitespace(); tok != zUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba zUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba valores", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba zUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	////ARREGLO
	//
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != znew {
		return nil, fmt.Errorf("encontro %q, esperaba znew", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//
	////feach

	if tok, lit := p.scanIgnoreWhitespace(); tok != feach {
		return nil, fmt.Errorf("encontro %q, esperaba feach", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//PAQUETES
	if tok, lit := p.scanIgnoreWhitespace(); tok != zthis {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//funciones

	if tok, lit := p.scanIgnoreWhitespace(); tok != funcn {
		return nil, fmt.Errorf("encontro %q, esperaba funcn", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zmain

	if tok, lit := p.scanIgnoreWhitespace(); tok != zmain {
		return nil, fmt.Errorf("encontro %q, esperaba zmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zend {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	return stmt, nil

}
func (p *Parser) ParsezIF() (*zifstatement, error) {
	stmt := &zifstatement{}
	// First token should be a "SELECT" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != zIF {
		return nil, fmt.Errorf("encontro %q, esperaba zif", lit)
	}
	return stmt, nil
}


func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}
