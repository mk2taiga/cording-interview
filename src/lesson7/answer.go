package lesson7

// 7.1

type Deck struct {
	max   int
	cards []Card
}

type Hand interface {
	cards() []Card
}

type Card interface {
	name() string
}

type BlackJackCard struct {
	Card
	val int
}

type BlackJackHand struct {
	Hand
}

// 7.2

type Employee interface {
	name() string
}

type Response struct {
	Employee
	// 仕事が終わったことを伝えられるように handler を保持しておくといいでしょう。
}

type Manager struct {
	Employee
}

type Director struct {
	Employee
}

type CallHandler struct {
	responses []*Response
	managers  []*Manager
	directors []*Director
	// 対応できなかった Call を管理するキューを持っておくと良さそう。
}

// Call は問い合わせを管理する。
type Call struct {
	caller Employee
}

func (ch *CallHandler) dispatchCall() Employee {
	//  TODO:
	//  	リストの0番目から順番に取得する。
	// 		response → manager → director の順でサーチする。
	return ch.responses[0]
}
