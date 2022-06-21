package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		wasInCache = c.Set("ddd", 300)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		val, ok = c.Get("ddd")
		require.True(t, ok)
		require.Equal(t, 300, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		_ = c.Set("a", 1)
		_ = c.Set("b", 2)
		_ = c.Set("c", 3)
		_ = c.Set("d", 4)

		nilItem := new(ListItem)
		val, _ := c.Get("a")
		if val == nilItem {
			val, ok := c.Get("a")
			require.Nil(t, val)
			require.False(t, ok)
		}
		val, ok := c.Get("d")
		require.True(t, ok)
		require.Equal(t, 4, val)

		// Проверка логики выталкивания давно используемых элементов
		_, _ = c.Get("b")
		_, _ = c.Get("c")
		_, _ = c.Get("d")
		_, _ = c.Get("b")
		_ = c.Set("c", 5) // самый старый элемент на момент завершения манипуляций с кешом
		_ = c.Set("d", 6)
		_, _ = c.Get("b")

		wasInCache := c.Set("f", 7)
		require.False(t, wasInCache)
		val, ok = c.Get("f")
		require.True(t, ok)
		require.Equal(t, 7, val)

		val, _ = c.Get("c")
		if val == nilItem {
			val, ok := c.Get("c")
			require.Nil(t, val)
			require.False(t, ok)
		}
	})

	t.Run("Clear method", func(t *testing.T) {
		c := NewCache(3)
		_ = c.Set("a", 1)
		_ = c.Set("b", 2)
		_ = c.Set("c", 3)
		c.Clear()
		val, ok := c.Get("a")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
