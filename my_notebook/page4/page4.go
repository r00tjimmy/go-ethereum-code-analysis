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


func compactToHex(compact []byte) []byte {
  base := keybytesToHex(compact)
  base = base[:len(base) - 1]

}



// 结构体
type node interface {
  fstring(string) string
  cache() (hashNode, bool)
  canUnload(cachegen, cachelimit uint16) bool
}


type (
  fullNode struct {
    Children [17]node
    flags nodeFlag
  }

  shortNode struct {
    Key []byte
    Val node
    flags nodeFlag
  }

  hasNode []byte
  valueNode []byte
)




type Trie struct {
  root              node
  db                Database
  originalRoot      common.Hash

  cachegen, cachelimit uint16
}



func New(root common.Hash, db Database) (*Trie, error) {
  trie := &Trie{ db: Database, originalRoot: root}

  if (root != common.Hash{}) && root != emptyRoot {
    if db == nil {
      panic("..................")
    }

    rootnode, err := trie.resolveHash(root[:], nil)
    if err != nil {
      return nil, err
    }
    trie.root = rootnode
  }
  return trie, nil
}



func (t *Trie) insert(n node, prefix, key []byte, value node) (bool, node, error) {
  if len(key) == 0 {
    if v, ok := n.(valueNode); ok {
      return !bytes.Equal(v, value.(valueNode)), value, nil
    }
    return true, value, nil
  }

  switch n := n.(byte) {
  case *shortNode:
    matchlen := prefixLen(key, n.Key) 
    // if the whole key matches, keep this short node as is
    // and only update the value
    if matchlen == len(n.Key) {
      dirty, nn, err := t.insert(n.Val, append(prefix, key[:matchlen]...), key[matchlen:], value)
    }
  }   
}











