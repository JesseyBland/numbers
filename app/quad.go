package app

//Quad is the struck of our number image
type Quad struct {
	l1 string
	l2 string
	l3 string
	l4 string
}

//CNUMBER Quadrant varriations
var (
	hx      = "   #   "
	hl      = "#"
	q1zero  = Quad{"   ", "   ", "   ", "   "}
	q1one   = Quad{"###", "   ", "   ", "   "}
	q1two   = Quad{"   ", "   ", "   ", "###"}
	q1three = Quad{"   ", "#  ", " # ", "  #"}
	q1four  = Quad{"  #", " # ", "#  ", "   "}
	q1five  = Quad{"###", " # ", "#  ", "   "}
	q1six   = Quad{"  #", "  #", "  #", "  #"}
	q1seven = Quad{"###", "  #", "  #", "  #"}
	q1eight = Quad{"  #", "  #", "  #", "###"}
	q1nine  = Quad{"###", "  #", "  #", "###"}

	q2zero  = Quad{"   ", "   ", "   ", "   "}
	q2one   = Quad{"###", "   ", "   ", "   "}
	q2two   = Quad{"   ", "   ", "   ", "###"}
	q2three = Quad{"   ", "  #", " # ", "#  "}
	q2four  = Quad{"#  ", " # ", "  #", "   "}
	q2five  = Quad{"###", " # ", "  #", "   "}
	q2six   = Quad{"#  ", "#  ", "#  ", "#  "}
	q2seven = Quad{"###", "#  ", "#  ", "#  "}
	q2eight = Quad{"#  ", "#  ", "#  ", "###"}
	q2nine  = Quad{"###", "#  ", "#  ", "###"}

	q3zero  = Quad{"   ", "   ", "   ", "   "}
	q3one   = Quad{"   ", "   ", "   ", "###"}
	q3two   = Quad{"###", "   ", "   ", "   "}
	q3three = Quad{"#  ", " # ", "  #", "   "}
	q3four  = Quad{"   ", "  #", " # ", "#  "}
	q3five  = Quad{"   ", "  #", " # ", "###"}
	q3six   = Quad{"#  ", "#  ", "#  ", "#  "}
	q3seven = Quad{"#  ", "#  ", "#  ", "###"}
	q3eight = Quad{"###", "#  ", "#  ", "#  "}
	q3nine  = Quad{"###", "#  ", "#  ", "###"}

	q4zero  = Quad{"   ", "   ", "   ", "   "}
	q4one   = Quad{"   ", "   ", "   ", "###"}
	q4two   = Quad{"###", "   ", "   ", "   "}
	q4three = Quad{"  #", " # ", "#  ", "   "}
	q4four  = Quad{"   ", "#  ", " # ", "  #"}
	q4five  = Quad{"   ", "#  ", " # ", "###"}
	q4six   = Quad{"  #", "  #", "  #", "  #"}
	q4seven = Quad{"  #", "  #", "  #", "###"}
	q4eight = Quad{"###", "  #", "  #", "  #"}
	q4nine  = Quad{"###", "  #", "  #", "###"}
)
