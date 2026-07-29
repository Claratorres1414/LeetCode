package main

func reverseBits(n int) int {
	var result int

	for i := 0; i < 32; i++ {
		result <<= 1    //abre espaço para o próximo bit
		result |= n & 1 //copia o bit do número original
		n >>= 1         //remove bit copiado
	}

	return result
}
