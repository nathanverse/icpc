# Graph Coloring Problem

## Problem Description

You are given an undirected graph $G = (V, E)$, where $V$ is a set of vertices and $E$ is a set of edges. Your task is to assign a "color" to each vertex such that no two adjacent vertices (i.e., vertices connected by an edge) have the same color. The goal is to find the minimum number of colors required to color the entire graph. This minimum number of colors is called the chromatic number of the graph, often denoted as $\chi(G)$.

**Constraints:**
* Two vertices connected by an edge must have different colors.
* You want to minimize the total number of colors used.

## Input Format

The input will consist of multiple test cases. Each test case will start with two integers:
* `N`: The number of vertices in the graph ($1 \le N \le 15$).
* `M`: The number of edges in the graph ($0 \le M \le N*(N-1)/2$).

Following these two integers, there will be `M` lines, each containing two integers:
* `u, v`: Representing an an edge between vertex `u` and vertex `v` ($1 \le u, v \le N$, $u \ne v$).

The vertices are 1-indexed.

## Output Format

For each test case, output a single integer representing the minimum number of colors required to color the graph (i.e., its chromatic number).
