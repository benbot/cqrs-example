package broadcast

type Broadcaster struct {
	readers  []chan interface{}
	listener chan interface{}
	caster   chan interface{}
}

func NewBroadcaster() *Broadcaster {
	b := &Broadcaster{
		[]chan interface{}{},
		make(chan interface{}),
		make(chan interface{})}

	go update(b)

	return b
}

func (b *Broadcaster) Write(v interface{}) {
	b.listener <- v
}

func (b *Broadcaster) Listen() chan interface{} {
	ch := make(chan interface{})
	b.readers = append(b.readers, ch)

	return ch
}

func update(b *Broadcaster) {
	for {
		val := <-b.caster
		for _, c := range b.readers {
			c <- val
		}
	}
}
