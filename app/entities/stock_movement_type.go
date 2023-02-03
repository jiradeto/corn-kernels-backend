package entities

// StockMovementType is enum type for stock movement type
type StockMovementType int

// StockMovementType enum constants
const (
	StockMovementTypeIn StockMovementType = iota + 1
	StockMovementTypeOut
)

var StockMovementTypeIndexMapper = map[string]StockMovementType{
	"in":  StockMovementTypeIn,
	"out": StockMovementTypeOut,
}

var StockMovementTypeStringMapper = map[StockMovementType]string{
	StockMovementTypeIn:  "in",
	StockMovementTypeOut: "out",
}

// Parse converts string to StockMovementType
func (c StockMovementType) Parse(StockMovementType string) StockMovementType {
	return StockMovementTypeIndexMapper[StockMovementType]
}

func (c StockMovementType) String() string {
	return StockMovementTypeStringMapper[c]
}

// Is is a function to check whether StockMovementType is equal to expected StockMovementType
func (c StockMovementType) Is(expected StockMovementType) bool {
	return c.String() == expected.String()
}
