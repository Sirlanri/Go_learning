create table course(
    cno char(4) primary key,
    cname char(40) ,
    cpno char(4),
    ccredit smallint,
    foreign key(cpno) references course(cno)
);

create table student(
    sno char(9) primary key,
    ssex char(20) unique,
    sage smallint,
    sdept char(20)
);


INSERT INTO `course`(`"cno","cname","cpno","ccredit"`) VALUES
(1,"数据库",5,4);

insert into course
values
(10,"数据库3",9,4);


INSERT INTO `sc` VALUES
("201215121",1,92),
("201215121",2,85),
("201215121",3,88),
("201215121",2,90),
("201215121",3,80)
;

/*mysql里面没有模式这个东西，不能使用的*/
create schema test authorization testuser
create table tab1 (
    col1 smallint,
    col2 int,
    col3 char(20),
    col4 numeric(10,3),
    col5 decimal(5,2)
);

/*创建删除数据库*/
create database test1;
DROP DATABASE `test1`;

/*向Student表增加“入学时间”列，其数据类型为日期型*/
ALTER TABLE `student` ADD `s_entrance` date;

/*将年龄的数据类型由字符型*/
alter table student alter column sage int;

/*增加课程名称必须取唯一值的约束条件*/
alter table course add unique(canme);

/*删除表*/
DROP TABLE student cascade;

/*创建一个视图*/
create view is_student 
    as
        SELECT
          sno,sname,sage
        FROM `student`
        WHERE `sdept` = 'is';

/*建立索引*/
create unique index stusno on student(sno);

/*修改索引*/
alter index scno rename to scsno;

/**删除索引*/
drop index stusname;

/*查询*/
SELECT
    `student`.sno,sname,sdept
  FROM `student`;

select * from student;

/*查询全体学生的姓名、出生年份和所在的院系*/
SELECT Sname,'Year of Birth: ',2020-Sage,LOWER(Sdept)
FROM Student;

/*查询选修了课程的学生学号*/
select sno from sc;

/*去除重复行*/
select distinct sno from sc;

/*查询计算机科学系全体学生的名单*/
SELECT Sname
    FROM Student
    WHERE Sdept='cs';

/*查询所有年龄在20岁以下的学生姓名及其年龄。*/
SELECT Sname,Sage
    FROM Student
    WHERE Sage < 20;

/*查询考试成绩有不及格的学生的学号*/
SELECT DISTINCT Sno
FROM SC
WHERE Grade<60;

/*查询年龄在20~23岁（包括20岁和23岁）之间的学生的姓名、系别和年龄*/
SELECT Sname, Sdept, Sage
FROM Student
WHERE Sage BETWEEN 20 AND 23;

/*查询年龄不在20~23岁之间的*/
SELECT Sname, Sdept, Sage
FROM Student
WHERE Sage NOT BETWEEN 20 AND 23;

/*查询计算机科学系（CS）、数学系（MA）和信息系（IS）学生的姓名和性别。*/
SELECT Sname, Ssex
FROM Student
WHERE Sdept IN ('CS','MA','is' );

/*查询既不是计算机科学系、数学系，也不是信息系的学生的姓名和性别*/
SELECT Sname, Ssex
FROM Student
WHERE Sdept NOT IN ('IS','MA','cs');

/*查询学号为201215121的学生的详细情况。*/
SELECT *
FROM Student
WHERE Sno LIKE '201215121';

/*查询所有姓刘学生的姓名、学号和性别*/
SELECT Sname, Sno, Ssex
FROM Student
WHERE Sname LIKE '刘%';

/*查询姓"欧阳"且全名为三个汉字的学生*/
SELECT Sname
FROM Student
WHERE Sname LIKE '欧阳_';

/*查询名字中第2个字为"阳"字的学生*/
SELECT Sname,Sno
FROM Student
WHERE Sname LIKE '__阳%';

/*查询所有不姓刘的学生姓名、学号和性别*/
SELECT Sname, Sno, Ssex
FROM Student
WHERE Sname NOT LIKE '刘%';

/*查询DB_Design课程的课程号和学分。*/
SELECT Cno，Ccredit
FROM Course
WHERE Cname LIKE 'DB\_Design' ESCAPE '\ ' ;

/*查询以"DB_"开头，且倒数第3个字符为 i的课程的详细情况*/
SELECT *
FROM Course
WHERE Cname LIKE 'DB\_%i_ _' ESCAPE '\ ' ;

/*查询缺少成绩的学生的学号和相应的课程号。*/
SELECT Sno,Cno
FROM SC
WHERE Grade IS NULL;

/*查所有有成绩的学生学号和课程号。*/
SELECT Sno,Cno
FROM SC
WHERE Grade IS NOT NULL;

/*查询计算机系年龄在20岁以下的学生姓名。*/
SELECT Sname
FROM Student
WHERE Sdept= 'CS' AND Sage<20;

/*查询计算机科学系（CS）、数学系（MA）和信息系（IS）学生的姓名和性别。*/
SELECT Sname, Ssex
FROM Student
WHERE Sdept IN ('CS ','MA ','IS');

SELECT Sname, Ssex
FROM Student
WHERE Sdept= ' CS' OR Sdept= ' MA' OR Sdept= 'IS ';

/*查询选修了3号课程的学生的学号及其成绩*/
SELECT Sno, Grade
FROM SC
WHERE Cno= '3'
ORDER BY Grade DESC;

/*查询全体学生情况*/
SELECT *
FROM Student
ORDER BY Sdept, Sage DESC;

/*查询学生总人数。*/
SELECT COUNT(*)
FROM Student;

/*查询选修了课程的学生人数。*/
SELECT COUNT(DISTINCT Sno)
FROM SC;

/*计算1号课程的学生平均成绩。*/
SELECT AVG(Grade)
FROM SC
WHERE Cno= ' 1 ';

/*查询选修1号课程的学生最高分数。*/
SELECT MAX(Grade)
FROM SC
WHERE Cno='1';

/*查询学生201215012选修课程的总学分数。*/
SELECT SUM(Ccredit)
FROM SC,Course
WHERE Sno='201215012' AND SC.Cno=Course.Cno;

/*求各个课程号及相应的选课人数。*/
SELECT Cno,COUNT(Sno)
FROM SC
GROUP BY Cno;

/*查询选修了3门以上课程的学生学号。*/
SELECT Sno
FROM SC
GROUP BY Sno
HAVING COUNT(*) >3;

/*查询平均成绩大于等于90分的学生学号和平均成绩*/
select sno,avg(grade)
from sc
group by sno
having avg(grade) >= 90;



/*第二节*/



/*查询每个学生及其选修课程的情况*/
SELECT Student.*, SC.*
FROM Student, SC
WHERE Student.Sno = SC.Sno;

/*查询选修2号课程且成绩在90分以上的所有学生的学号和姓名。*/
SELECT Student.Sno, Sname
FROM Student, SC
WHERE Student.Sno=SC.Sno AND
SC.Cno=' 2 ' AND SC.Grade>90;

/*查询每一门课的间接先修课*/
SELECT FIRST.Cno, SECOND.Cpno
FROM Course FIRST, Course SECOND
WHERE FIRST.Cpno = SECOND.Cno;

SELECT Student.Sno,Sname,Ssex,Sage,Sdept,Cno,Grade
FROM Student LEFT OUT JOIN sc ON
(Student.Sno=SC.Sno);

/*嵌套查询*/
SELECT Sname /*外层查询/父查询*/
FROM Student
WHERE Sno IN
(SELECT Sno /*内层查询/子查询*/
FROM SC
WHERE Cno= '2');

/*确定“刘晨”所在系名*/
SELECT Sdept
FROM Student
WHERE Sname= '刘晨';

/*查找所有在CS系学习的学生*/
SELECT Sno, Sname, Sdept
FROM Student
WHERE Sdept= 'CS';

/*和她一个专业的学生*/
SELECT Sno, Sname, Sdept
FROM Student
WHERE Sdept IN
    (SELECT Sdept
        FROM Student
        WHERE Sname= '刘晨');

/*用自身连接*/
SELECT S1.Sno, S1.Sname,S1.Sdept
FROM Student S1,Student S2
    WHERE S1.Sdept = S2.Sdept AND
    S2.Sname = '刘晨';

/*查询选修了课程名为“信息系统”的学生学号和姓名*/
/*这是我自己写的，没看屁屁踢*/
select sno,sname
    from student
    where sno in (
        select sno from sc
            where cno = (
                select cno from course
                    where cname = '信息系统'
            )
    );

/*每个学生超过他选修课程平均成绩的课程号*/
SELECT Sno, Cno
FROM SC x
    WHERE Grade >=(SELECT AVG（Grade）
        FROM SC y
        WHERE y.Sno=x.Sno);

/*查询非计算机科学系中比计算机科学系任意一个学生年龄小的学生姓名和年龄*/
SELECT Sname,Sage
FROM Student
    WHERE Sage < ANY (
        SELECT Sage
        FROM Student
        WHERE Sdept= ' CS ');

SELECT Sname,Sage
FROM Student
WHERE Sage <
    (SELECT MAX(Sage)
        FROM Student
        WHERE Sdept= 'CS ')
            AND Sdept <> ' CS ';

/*查询所有选修了1号课程的学生姓名。*/
SELECT Sname
FROM Student
WHERE EXISTS
    (SELECT *
    FROM SC
    WHERE Sno=Student.Sno AND Cno= '1');