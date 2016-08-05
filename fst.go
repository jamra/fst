package fst

//Fst is a Finite state transducer
type Fst struct {
	TrieNode
}

func newFst() *Fst {
	f := &Fst{}
	return f
}

func BuildFst(scanner bufio.Scanner) error {
	fst := newFst()
	
	for scanner.Scan() {
		fst.Add(scanner.Text()) 
	}
}
