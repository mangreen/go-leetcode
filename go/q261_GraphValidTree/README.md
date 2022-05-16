# [261 - Graph Valid Tree](https://leetcode.ca/all/261.html)

- https://leetcode.ca/2016-08-17-261-Graph-Valid-Tree/
- https://blog.csdn.net/qq_37821701/article/details/104168177

## Level
Medium

## Similar Question
- 323 - Number of Connected Components in an Undirected Graph

## Description
Given n nodes labeled from 0 to n-1 and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.

Example 1:
input: n = 5, edges = [[0,1], [0,2], [0,3], [1,4]]
output: true
  0
 /|\
1 2 3
|
4

Example 2:
input: n = 5, edges = [[0,1], [1,2], [2,3], [1,3], [1,4]]
output: false
    0
    | 
4 - 1 - 2
     \ /
      3

Example 2:
input: n = 4, edges = [[0,1], [2,3], [1,4]]
output: false
    0
    | 
4 - 1   2 - 3

Note: you can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0,1] is the same as [1,0] and thus will not appear together in edges.