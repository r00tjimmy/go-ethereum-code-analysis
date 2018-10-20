RPC包的官方文档
=====================================


JSONRPC标准

请求的格式:

{
"jsonrpc":  "2.0",
"method":   "eth_getBlockByNumber",
"params":   ["0x1b4", true],
"id":       1
}




返回的格式:

{
"jsonrpc":  "2.0",
"id":       1,
"result":   
  {
    "difficulty":     "0x23h23hj2h3123123x",
    "gasUsed":        "0x0",
    "size":           "0x220" 
  }
}









