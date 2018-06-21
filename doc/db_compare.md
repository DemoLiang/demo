# 数据库特性比较


mongoDB
- 特性
```
    面向集合
            易存储对象类型的数据，包括文档内嵌对象及数组，支持二进制及大型对象
    模式自由
            无需知道存储数据的任何结构定义，支持动态查询、完全索引
    文档型
            以键-值对形式存储，支持数组，支持文档之间嵌套
    支持B+索引，全文索引，地理空间索引

```
- 适用场景
```
    更高的写负载
    不可靠环境保证高可用
    数据量超大规模，大尺寸，低价值的数据 
    基于位置的数据查询      
    非结构化数据的爆发增长
    常用的场景包括
            Web应用程序
            敏捷开发
            分析和日志(目标原子更新，定长集合)
            缓存
            可变Schema
    适用于需要动态查询支持；需要使用索引而不是 map/reduce功能；需要对大数据库有性能要求
    MongoDB适合做读写分离场景中的读取场景
    Cassandra与MongoDB之间使用定时同步，适合一致性要求不是特别强的业务。
```
- 不适用场景
```
    高度事务性，强一致性业务系统(银行，证券等)
    传统商业智能应用
    极为复制的业务逻辑查询
```


Cassandra
- 特性
```
数据最终一致性，
而HBase靠Master节点管理数据的分配，将过热的节点上的Region动态分配给负载较低的节点
Cassandra通过一致性哈希来决定数据存储的位置
强调CAP中的A（availability），和尽量满足的C（consistency），理论上很美好可以用read replica + write replica> total replica
```
- 适用场景
```
Cassandra的scan效率比HBase低，但是可以支持更高的并发写与读，并且高并发写和读可以在不同的配置的情况下在一个表上同时实现（由于读写的replica number是可以基于per request 定义的）如果你需要高并发可调节读写，scan需求少，那么Cassandra则比HBase更合适。
Cassandra适合读写分离的场景，写入场景使用Cassandra，比如插入操作日志，或领域事件日志的写入
Cassandra与MongoDB之间使用定时同步，适合一致性要求不是特别强的业务。
```

HBase
- 特性
```
HBase基于Big Table，和Hadoop MapReduce完美integrate，有一个master node储存metadata，同时保证write总是发给shard的leader所以可以保证写的顺序和强一致性
```

- 适用场景
```
需要大量scan操作 或者需要经常配合MapReduce
HBase比较中庸些，适合各种场景
```

- 不适用场景
```
由于HBase存在Master节点，因此会存在单点问题
并且需要对大数据进行随机、实时访问的场合。
```