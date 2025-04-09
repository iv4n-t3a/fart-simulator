package util

func NewAndOldValue[T comparable](old T, x1 T, x2 T) (T, T) {
	if x1 == old {
		return x1, x2
	}
	if x2 == old {
		return x2, x1
	}
  print(old)
  print(x1)
  print(x2)
  panic("No mathcing value")
}
