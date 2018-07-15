page 3 的一些备注
========================================


1. golang中sync包实现了两种锁Mutex （互斥锁）和RWMutex（读写锁），其中RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能

2. Recursive Length Prefix， 简称RLN， 这个是用来做比特币的通信内容序列化的

3. 注意要看所有eth源码文件的测试文件， 文件tests目录, 去测试用例里面验证自己的源码逻辑理解


### 巩固的基础
-   interface类型
-   byte类型
-   test *testing.T 的 t.Parallel() 方法的使用



