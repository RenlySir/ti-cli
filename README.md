# ti-cli
目的：这工具目的是方便TiDB DBA快速的获取一些排查问题所需信息，降低运维门槛吧。

原因：因为TiDB是严格分布式数据库，与传统单机数据库不同，它涉及store，region，table_id等概念，这些信息的映射关系在排查问题时经常会给人带来困扰。

当前状况：只实现了简单几个接口其实，还有很多东西没加，算是个demo状态吧。

计划：计划后续可以访问到pd的etcd，查询一些相关有助于问题排查的信息；

1.getinfo
获取tidb集群状态信息，是一个试验的接口，如下：
```
sirlan@sirdeMacBook-Pro ti-cli % ./ti-cli getinfo --host 127.0.0.1  --port 10080                                {
    "connections": 0,
    "version": "5.7.25-TiDB-v5.0.0",
    "git_hash": "bdac0885cd11bdf571aad9353bfc24e13554b91c"
}
```

2.getStore
快速获取store id和tikv实例映射关系
```
sirlan@sirdeMacBook-Pro ti-cli % ./ti-cli getStore --pdHost 127.0.0.1 --pdPort 2379 
this tidb cluster store count is :  3
store id and status address relation and tikv status is :
1, 172.16.7.229:20180, Up 
5, 172.16.7.200:20180, Up 
604, 172.16.7.176:20180, Up 
```

3.getRegions
根据库名表名，快速查到涉及到的region，以及涉及到的region中的信息
```
sirlan@sirdeMacBook-Pro ti-cli % ./ti-cli getRegions --host 172.16.7.150 --port 10080 --db test --table loan_contract_info
table name is : loan_contract_info，table id is : 275 
region id is : 324201, store id is : 1 
index name is : idx_cui01_loan_contract_info, index id is : 1 
start key is : dIAAAAAAAAEOX3KAAAAAAAZE9w==,end key is: <nil> 

table region ,db name is : test,table name is: test,table id is: 273 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui01_loan_contract_info ,index id is: 1 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui02_loan_contract_info ,index id is: 2 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui03_loan_contract_info ,index id is: 3 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui04_loan_contract_info ,index id is: 4 
table region ,db name is : test,table name is: loan_contract_info,table id is: 275 
table region ,db name is : sbtest1,table name is: sbtest1,table id is: 270 
index region id is : 324201, store id is: 1 

index name is : idx_cui02_loan_contract_info, index id is : 2 
start key is : dIAAAAAAAAEOX3KAAAAAAAZE9w==,end key is: <nil> 
table region ,db name is : test,table name is: test,table id is: 273 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui01_loan_contract_info ,index id is: 1 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui02_loan_contract_info ,index id is: 2 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui03_loan_contract_info ,index id is: 3 
index region ,it db name is : test,table name is: loan_contract_info,table id is: 275 ,index name is: idx_cui04_loan_contract_info ,index id is: 4 
table region ,db name is : test,table name is: loan_contract_info,table id is: 275 
table region ,db name is : sbtest1,table name is: sbtest1,table id is: 270 
index region id is : 324201, store id is: 1 
```
