ethdb源码解释  page 5 的一些备注
========================================


### --------------------

源码所在的目录在ethereum/ethdb目录。代码比较简单， 分为下面三个文件

database.go levelDB的封装代码
memory_database.go  供测试用的基于内存的数据库，不会持久化为文件，仅供测试
interface.go  定义了数据库的接口
database_test.go  测试案例




