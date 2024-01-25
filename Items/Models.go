package items

type Item interface {
	GetNom() string
	TypeItem() int
	StackMax() int
	GetStack() int
	SetStack(int)
	Clone() Item
}
