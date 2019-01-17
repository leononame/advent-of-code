package day21

// 00 seti 123 0 5
// 01 bani 5 456 5
// 02 eqri 5 72 5
// 03 addr 5 2 2
// 04 seti 0 0 2
// 05 seti 0 9 5
// 06 bori 5 65536 3
// 07 seti 7586220 4 5
// 08 bani 3 255 1
// 09 addr 5 1 5
// 10 bani 5 16777215 5
// 11 muli 5 65899 5
// 12 bani 5 16777215 5
// 13 gtir 256 3 1
// 14 addr 1 2 2
// 15 addi 2 1 2
// 16 seti 27 9 2
// 17 seti 0 9 1
// 18 addi 1 1 4
// 19 muli 4 256 4
// 20 gtrr 4 3 4
// 21 addr 4 2 2
// 22 addi 2 1 2
// 23 seti 25 4 2
// 24 addi 1 1 1
// 25 seti 17 2 2
// 26 setr 1 6 3
// 27 seti 7 8 2
// 28 eqrr 5 0 1
// 29 addr 1 2 2
// 30 seti 5 0 2

func b2i(v bool) int {
	if v {
		return 1
	}
	return 0
}

// This file contains the reverse engineering of the program
// IP == R2
func rev01() {
	var R0, R1, R3, R4, R5 int
	// OP00:
	R5 = 123
OP01:
	R5 &= 456 // 5 = 72
	// OP02:
	R5 = b2i(R5 == 72) // R5 = 1
	// OP03:
	if R5 == 1 { // R2 += R5 => IP + R5 => IP++ => JUMP over next instruction
		goto OP05
	}
	// OP04:
	goto OP01 // R2 = 0 => IP = 0 => JUMP to Instruction 1
OP05:
	R5 = 0
OP06:
	R3 = R5 | 65536 // R3 = 65536
	// OP07:
	R5 = 7586220
OP08:
	R1 = R3 & 255
	// OP09:
	R5 += R1
	// OP10:
	R5 &= 16777215
	// OP11:
	R5 *= 65899
	// OP12:
	R5 &= 16777215
	// OP13:
	R1 = b2i(256 > R3)
	// OP14:
	if R1 == 1 { // R2 += R1 => IP + R1 =>
		goto OP16 // JUMP over next instruction if R1 == 1
	}
	// OP15:
	goto OP17 // R2++ => IP++ => JUMP over next instruction
OP16:
	goto OP28 // R2 = 27 => JUMP to 28
OP17:
	R1 = 0
OP18:
	R4 = R1 + 1
	// OP19:
	R4 *= 256
	// OP20:
	R4 = b2i(R4 > R3)
	// OP21:
	if R4 == 1 { // Logic is same as above
		goto OP23
	}
	// OP22:
	goto OP24
OP23:
	goto OP26
OP24:
	R1++
	// OP25:
	goto OP18
OP26:
	R3 = R1
	// OP27:
	goto OP08
OP28:
	R1 = b2i(R0 == R5)
	// OP29:
	if R1 == 1 {
		return // JUMP over next instruction => outside of program => halt
	}
	// OP30:
	goto OP06
}

func rev02() {
	var R0, R1, R3, R4, R5 int
	// OP00 - OP04
	R5 = 72
	// OP05:
	R5 = 0
OP06:
	R3 = R5 | 65536 // R3 = 65536
	// OP07:
	R5 = 7586220 // This changes for other people apparently, our "real" input
OP08:
	// OP08 - OP12
	R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
	// OP13-15
	if R3 < 256 {
		goto OP16
	}
	goto OP17
OP16:
	goto OP28 // R2 = 27 => JUMP to 28
OP17:
	R1 = 0
OP18:
	// OP18 - OP19
	R4 = (R1 + 1) * 256
	// OP20 - 22
	if R4 > R3 {
		goto OP23
	}
	goto OP24
OP23:
	goto OP26
OP24:
	R1++
	// OP25:
	goto OP18
OP26:
	R3 = R1
	// OP27:
	goto OP08
OP28:
	// OP28-OP30
	if R0 == R5 {
		return
	}
	goto OP06
}

func rev03() {
	var R0, R1, R3, R4, R5 int
	R5 = 0
OP06:
	R3 = R5 | 65536 // R3 = 65536
	R5 = 7586220    // This changes for other people apparently, our "real" input

OP08:
	R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
	if R3 < 256 {
		goto OP28
	}
	R1 = 0

OP18:
	R4 = (R1 + 1) * 256
	if R4 > R3 {
		goto OP26
	}
	R1++
	// OP25:
	goto OP18

OP26:
	R3 = R1
	goto OP08

OP28:
	if R0 == R5 {
		return
	}
	goto OP06
}

func rev04() {
	var R0, R1, R3, R4, R5 int
	R5 = 0
OP06:
	R3 = R5 | 65536 // R3 = 65536
	R5 = 7586220    // This changes for other people apparently, our "real" input

OP08:
	R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
	if R3 < 256 {
		if R0 == R5 {
			return
		}
		goto OP06
	}
	R1 = 0

OP18:
	R4 = (R1 + 1) * 256
	if R4 > R3 {
		R3 = R1
		goto OP08
	}
	R1++
	goto OP18
}

func rev05() {
	var R0, R1, R3, R4, R5 int
	R5 = 0
outer:
	for {
		R3 = R5 | 65536 // R3 = 65536
		R5 = 7586220    // This changes for other people apparently, our "real" input
		for {
			R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
			if R3 < 256 {
				if R0 == R5 {
					return
				}
				continue outer
			}
			R1 = 0
			for {
				R4 = (R1 + 1) * 256
				if R4 > R3 {
					R3 = R1
					break
				}
				R1++
			}
		}
	}
}

func rev06() {
	var R0, R1, R3, R4, R5 int
	R5 = 0
	for {
		R3 = R5 | 65536 // R3 = 65536
		R5 = 7586220    // This changes for other people apparently, our "real" input
		for {
			R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
			if R3 < 256 {
				break
			}
			R1 = 0
			for {
				R4 = (R1 + 1) * 256
				if R4 > R3 {
					break
				}
				R1++
			}
			R3 = R1
		}
		if R0 == R5 {
			return
		}
	}
}

func rev07() {
	var R0, R3, R4, R5 int
	R5 = 0
	for {
		R3 = R5 | 65536 // R3 = 65536
		R5 = 7586220    // This changes for other people apparently, our "real" input
		for {
			R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
			if R3 < 256 {
				break
			}

			for i := 0; ; i++ {
				R4 = (i + 1) * 256
				if R4 > R3 {
					R3 = i
					break
				}
			}
		}
		if R0 == R5 {
			return
		}
	}
}

func rev08() {
	var R0, R3, R5 int
	R5 = 0
	for {
		R3 = R5 | 65536 // R3 = 65536
		R5 = 7586220    // This changes for other people apparently, our "real" input
		for {
			R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
			if R3 < 256 {
				break
			}

			R3 = R3 / 256
		}
		if R0 == R5 {
			return
		}
	}
}

func rev09(input int) {
	// input = 7586220 => Hard coded value replaced by parameter
	var R0, R3, R5 int
	for {
		R3 = R5 | 65536 // R3 = 65536
		for R5 = input; ; R3 /= 256 {
			R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
			if R3 < 256 {
				break
			}
		}
		if R0 == R5 {
			return
		}
	}
}
