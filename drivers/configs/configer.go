package configs

// Configer はConfigを取得する為に利用するinterfaceです。
// package内でConfigをメンバーに持つStructは全てこのinterfaceを実装してください。
type Configer interface {
	// Get() はConfigを取得します。
	Get() Config
}
