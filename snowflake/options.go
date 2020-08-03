package snowflake

var (
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1500000001000 // milliseconds

	// NodeBits holds the number of bits to use for Node
	// Remember, you have a total 22 bits to share between Node/Step
	DefaultNodeBits uint8 = 10

	// StepBits holds the number of bits to use for Step
	// Remember, you have a total 22 bits to share between Node/Step
	DefaultStepBits uint8 = 12
)

type Options struct {
	NodeBits uint8
	StepBits uint8
	Node     int64
}

type Option func(*Options)

func newOptions(opts ...Option) *Options {
	opt := &Options{
		NodeBits: DefaultNodeBits,
		StepBits: DefaultStepBits,
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// 根据 timestamp_ms-epoch <= 2^(maxBits-nodeBits-stepBits)) 可得出 timestamp_ms <= 2^(53-nodeBits-stepBits))+epoch
// 所以使用寿命为：(2^(maxBits-nodeBits-stepBits)+epoch-now_ms)/31536000000 年

// Note: 兼容float64的ID
// 因为float64精度只占53位，所以为了兼容float64，总的ID最大有效位数(maxBits)不应超过53位
// 例如，若将node和step总位数限制在12位，则可使用的相对毫秒时间戳的位数为41位，即最大使用寿命为69年

// WithNodeBits sets default NodeBits for idgen
// Remember, you have a total 22 bits to share between Node/Step
func WithNodeBits(nodeBits uint8) Option {
	return func(opts *Options) {
		opts.NodeBits = nodeBits
	}
}

// WithStepBits sets default StepBits for idgen
// Remember, you have a total 22 bits to share between Node/Step
func WithStepBits(stepBits uint8) Option {
	return func(opts *Options) {
		opts.StepBits = stepBits
	}
}

// WithNode sets default Node for idgen
func WithNode(node int64) Option {
	return func(opts *Options) {
		opts.Node = node
	}
}
