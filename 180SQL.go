package leetcode

//编写一个 SQL 查询，查找所有至少连续出现三次的数字。
//+----+-----+
//| Id | Num |
//+----+-----+
//| 1  |  1  |
//| 2  |  1  |
//| 3  |  1  |
//| 4  |  2  |
//| 5  |  1  |
//| 6  |  2  |
//| 7  |  2  |
//+----+-----+
//例如，给定上面的 Logs 表， 1 是唯一连续出现至少三次的数字。
//+-----------------+
//| ConsecutiveNums |
//+-----------------+
//| 1               |
//+-----------------+

//select distinct l1.Num as ConsecutiveNums from Logs as l1,Logs as l2,Logs as l3 where l1.Id=l2.Id-1 and l2.Id=l3.Id-1 and l1.Num=l2.Num and l2.Num=l3.Num
