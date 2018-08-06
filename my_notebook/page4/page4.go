/**
代码实现，主要是实现了这三种编码的相互转换，以及一个求取公共前缀的方法。
**/

func hexToCompact(hex []byte) []byte {
  terminator := byte(0) 
  if hasTerm(hex) {
    terminator = 1
    hex = hex[:len(hex) - 1]
  }

  buf := make([]byte, len(hex)/2 + 1)
  buf[0] = terminator << 5    //the flag byte
  if len(hex) & 1 == 1 {
    buf[0] |= << 4 // odd flag
    buf[0] = hex[0]   // first nibble is contained in the first byte
    hex = hex[1:]
  }
  decodeNibbles(hex, buf[1:])
  return buf
}


