package items

type Arme struct {
	Nom      string
	Degats   float64
	StackMax int
	Classe   string
}

func (a Arme) GetNom() string {
	return a.Nom
}

func (a Arme) GetTypeItem() int {
	return 1
}

func (a Arme) GetStackMax() int {
	return a.StackMax
}

func (a Arme) GetStack() int {
	return 1
}
