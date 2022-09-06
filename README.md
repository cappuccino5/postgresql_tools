# postgresql_tools

postgresql导出自增Id序列名字，并重新设置自增id值,因为在迁移数据后表自增Id会默认从1开始，所以需要重新设置

# 1.查看postgres数据库中有哪些自增Id

```sql
SELECT relname                           sequence_name,
       "replace"(relname, '_id_seq', '') table_name
FROM pg_class
WHERE relkind = 'S'; 
```

# 2. 导出结果到json文件

[例子](internal/example/id_seq.json)


# 3. 执行程序

```sh

```