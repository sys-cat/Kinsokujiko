package Kinsokujiko

// Target is Mask Target
type Target struct {
	Surf string
	Pos  string
	Proc string
}

// Targets is Slice any Target
type Targets struct {
	Name    string   // ターゲット名
	Tag     []string // タグ名リスト
	Targets []Target // ターゲットリスト
}
