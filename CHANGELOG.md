
#1-Two Sum

基本算法：循环遍历 时间复杂度 O（N*N）
优化算法：hash表 O(N)

#2-AddTwo Numbers

基本算法：模拟整数加法，直接链表操作

#3-LengthOfLongestSubstring

基本算法：二重循环查找，最多做一些循环的优化

#4-MedianofTwoSortedArrays

考察归并排序的算法

#5- Longest Palindromic Substring

最长回文判断
从某个点往前做偶数串和奇数串的回文判断，然后记录开始和结束的点，最后输出。

#6 - ZigZag Conversion

说实话，没看别人的解答时根本不知道这题是要做啥
找图形规律

发现所有行的重复周期都是 2 * nRows - 2
对于首行和末行之间的行，还会额外重复一次，重复的这一次的索引为 (当前遍历值 / 周期 + 1) * 周期 - 本行第一个索引

#7 -Reverse Integer

主要是边界处理 溢出问题

#8 - String to Integer

主要问题还是边界处理

#9 -Palindrome Number

回文数判断，O(1)的空间限制

直接反转数字生成新的数字，注意一下溢出

#10 - Regular Expression Matching

第一眼想到了动态规则，没弄出状态转移方程

看了提示写出来的，递归和dp

#11 Container With Most Water

穷举或者从外往内取最大值

#12 Integer to Roman

简单题

其实开始看题的时候我也是蒙圈的，连个示例都没有。。

#13 Roman to Integer

接12题的简单题

#14 Longest Common Prefix

简单题

#15 3Sum

最开始自己写了一波，根据三种情况分别处理，但是超时了，看了其它的解题报告得到现在的程序。

首先对原数组进行排序，然后开始遍历排序后的数组，这里注意不是遍历到最后一个停止，而是到倒数第三个就可以了，然后我们还要加上重复就跳过的处理，对于遍历到的数，我们用0减去这个数得到一个sum，我们只需要再之后找到两个数之和等于sum即可，这样一来问题又转化为了求two sum，这时候我们一次扫描，找到了等于sum的两数后，加上当前遍历到的数字，按顺序存入结果中即可，然后还要注意跳过重复数字。






