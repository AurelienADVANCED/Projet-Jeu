package items

type Nourriture struct {
	Nom       string
	Stack     int
	Stack_max int
	VieRecup  float64
	Symbole   string
}

func NewNourriture(nom string, stack, stackMax int) *Nourriture {
	return &Nourriture{
		Nom:       nom,
		Stack:     stack,
		Stack_max: stackMax,
	}
}

func (n Nourriture) GetNom() string {
	return n.Nom
}

func (n Nourriture) TypeItem() int {
	return 2
}

func (n Nourriture) StackMax() int {
	return n.Stack_max
}

func (n Nourriture) GetStack() int {
	return n.Stack
}

func (n Nourriture) SetStack(newStack int) {
	n.Stack = newStack
}

func (n Nourriture) Clone() Item {
	return &Nourriture{
		Nom:       n.Nom,
		Stack:     n.Stack,
		Stack_max: n.Stack_max,
	}
}
