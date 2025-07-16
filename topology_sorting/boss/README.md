Hierarchy
A group of graduated students decided to establish a company; however, they don't agree on who is going to be who's boss.

Generally, one of the students will be the main boss, and each of the other students will have exactly one boss (and that boss, if he is not the main boss, will have a boss of his own too). Every boss will have a strictly greater salary than all of his subordinates - therefore, there are no cycles. Therefore, the hierarchy of the company can be represented as a rooted tree.

In order to agree on who is going to be who's boss, they've chosen
K
K most successful students, and each of them has given a statement: I want to be the superior of him, him, and him (they could be successful or unsuccessful). And what does it mean to be a superior? It means to be the boss, or to be one of the boss' superiors (therefore, a superior of a student is not necessary his direct boss).

Help this immature company and create a hierarchy that will satisfy all of the successful students' wishes. A solution, not necessary unique, will exist in all of the test data.

Dữ liệu nhập
In the first line of input, read positive integers
N
N (
N
≤
1
0
0
0
0
0
N≤100000), total number of students, and
K
K (
K
<
N
K<N), the number of successful students. All students are numbered
1
.
.
N
1..N, while the successful ones are numbered
1
.
.
K
1..K. Then follow
K
K lines. In Ath of these lines, first read an integer
W
W (the number of wishes of the student A,
1
≤
W
≤
1
0
1≤W≤10), and then
W
W integers from the range
[
1
,
N
]
[1,N] which denote students which student A wants to be superior to.

Dữ liệu xuất
Output
N
N integers. The Ath of these integers should be
0
0 if student A is the main boss, and else it should represent the boss of the student A.