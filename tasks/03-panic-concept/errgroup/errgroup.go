package errgroup

import (
	"context"
	"sync"
)

// Group представляет собой набор горутин, работающих над подзадачами,
// которые являются частью одной и той же общей задачи.
type Group struct {
	cancel func()

	wg sync.WaitGroup

	err     error
	errOnce sync.Once
}

// WithContext возвращает новую группу и связанный с ней контекст, наследованный от ctx.
//
// Контекст группы будет отменён, когда функция, переданная в Go вернёт ненулевую ошибку или запаникует,
// или когда завершится вызов Wait (в зависимости от того, что произойдет первым).
func WithContext(ctx context.Context) (*Group, context.Context) {
	// Реализуй меня.
	return nil, nil
}

// Wait является блокирующим вызовом. Ожидает, пока все функции, запущенные с помощью Go,
// завершат свою работу, а затем возвращает первую ненулевую ошибку (если она была) от них.
func (g *Group) Wait() error {
	// Реализуй меня.
	return nil
}

// Go запускает функцию f в новой горутине.
//
// Первая ненулевая ошибка (или отловленная паника) от очередного вызова f
// отменяет контекст группы, эта же ошибка будет возвращена в Wait.
func (g *Group) Go(f func() error) {
	// Реализуй меня.
}
