package items

type Arme struct {
	Nom       string
	Degats    float64
	Stack     int
	Stack_max int
	Classe    string
	Symbole   string
	Default   bool
}

func NewArme(nom string, degats float64, stackMax int, classe, symbole string) Arme {
	return Arme{
		Nom:       nom,
		Degats:    degats,
		Stack_max: stackMax,
		Classe:    classe,
		Symbole:   symbole,
	}
}

func (a *Arme) GetDegats() float64 {
	return a.Degats
}

func (a *Arme) GetClasse() string {
	return a.Classe
}

func (a *Arme) GetNom() string {
	return a.Nom
}

func (a *Arme) TypeItem() int {
	return 1
}

func (a *Arme) StackMax() int {
	return a.Stack_max
}

func (a *Arme) GetStack() int {
	return a.Stack
}

func (a *Arme) SetStack(newStack int) {
	a.Stack = newStack
}

func (a *Arme) GetSymbole() string {
	return "🗡️"
}

func (a *Arme) GetStackMax() int {
	return 1
}

func (a Arme) Clone() Item {
	return &Arme{
		Nom:       a.Nom,
		Degats:    a.Degats,
		Stack_max: a.Stack_max,
		Classe:    a.Classe,
		Symbole:   a.Symbole,
	}
}
