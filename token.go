package andrea

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK      // *
	COMMA         // ,
	DosPunt       // :
	ParDer        // (
	ParIsq        // )
	PuntCom       // ;
	LlaveDer      // {
	LlaveIsq      // }
	SaltoDeL      //n
	Asignaicon    // =
	AgruppadorDER // [
	AgruppadorISQ // ]

	// Keywords
	SELECT
	FROM
	si
	sw
	zend
	zINT // Tipodedato
	funcn
	zmain
	no
	para
	zcase
	var // variables
	break
	while
	
	tri
	cach
	zFINALLY
	zsize
	znew
	zout
	feach
	clase
	zthis

	//Bonifaz

	MCCREATE
	DATABASE
	MCUSE
	MCCREATETABLE

)
