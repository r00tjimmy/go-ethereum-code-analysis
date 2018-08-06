trie源码解释  page 4 的一些备注
========================================

包trie 实现了Merkle Patricia Tries，这里用简称MPT来称呼这种数据结构，这种数据结构实际上是一种Trie树变种，MPT是以太坊中一种非常重要的数据结构，用来存储用户账户的状态及其变更、交易信息、交易的收据信息。MPT实际上是三种数据结构的组合，分别是Trie树， Patricia Trie， 和Merkle树。下面分别介绍这三种数据结构。


### Trie树
简单来说  Trie树 就是一个节点就是一个字母


### Patricia Tries (前缀树)
Patricia 树 就是在原来Trie树的基础上， 把能共用的一些字符串封装起来， 作为一个节点， 然后根据这个节点的属性（hash值）去寻找接下来的其他字母


### Markle树
简单来说就是 叶子节点就是具体的数据块， 非叶子节点就是储存一些hash值， 这个hash值通常作为验证叶子节点的依据。




### BLOCK 的详细数据结构


struct {
  prevhash,
  state root, 
  tx root,
  receipt root,
}
