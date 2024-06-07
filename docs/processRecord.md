# 数据清洗的过程记录



#### 1. 过滤不包含密码的记录

```shell
$ awk '$6!=""' fakessh-2024.log >> fakessh-filter-empty-password.log

$ wc -l fakessh-filter-empty-password.log
357981 fakessh-filter-empty-password.log
```



#### 2. 初步统计

```shell
$ awk '{print $6}' fakessh-filter-empty-password.log | sort | uniq -c | sort -r| head -n 10
  22253 123456
   6952 123
   3950 1234
   2957 admin
   2927 12345
   2803 root
   2435 test
   2170 password
   2042 12345678
   2016 1
```

