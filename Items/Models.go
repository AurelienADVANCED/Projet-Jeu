package items

type Item interface {
	GetNom() string
	GetSymbole() string
	TypeItem() int
	StackMax() int
	GetStack() int
	GetStackMax() int
	SetStack(int)
	Clone() Item
}
