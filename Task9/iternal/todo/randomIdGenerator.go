package todo

import "time"

// quarterRound - это аналог шага ChaCha20, но с меньшими данными (метод рандомизации данных для айди)
func quarterRound(a, b, c, d uint16) (uint16, uint16, uint16, uint16) {
	a += b; d ^= a; d = (d << 7) | (d >> 9)
	c += d; b ^= c; b = (b << 9) | (b >> 7)
	a += b; d ^= a; d = (d << 13) | (d >> 3)
	c += d; b ^= c; b = (b << 8) | (b >> 8)
	return a, b, c, d
}

// randomIdGenerator - это аналог ChaCha20, но с меньшим кол-во шагов и обычным каунтером с счетчиком времени вместо шума OS
func randomIdGenerator(counter int) int {
	var a uint16 = 0xAAAA
	var b uint16 = uint16(time.Now().UnixNano())
	var c uint16 = uint16(counter)
	var d uint16 = 0x5555

	for i := 0; i < 8; i++	{
		a, b, c, d = quarterRound(a,b,c,d)
	}
	return int(a ^ b ^ c ^ d)
}