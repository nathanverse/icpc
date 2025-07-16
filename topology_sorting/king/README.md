# C. King's Path

## Problem Description

The black king is positioned on an enormous chessboard that spans $10^9$ rows and $10^9$ columns. Rows are numbered $1$ to $10^9$ from top to bottom, and columns are numbered $1$ to $10^9$ from left to right. A cell at the $i$-th row and $j$-th column is denoted as $(i, j)$.

You are given a set of *allowed* cells on this vast chessboard. These allowed cells are defined by $n$ segments. Each segment is described by three integers: $r_i$, $a_i$, and $b_i$ ($a_i \le b_i$). This means that all cells in the $r_i$-th row, from column $a_i$ to column $b_i$ (inclusive), are allowed.

Your task is to find the minimum number of moves the king requires to travel from a starting square $(x_0, y_0)$ to a target square $(x_1, y_1)$. The crucial constraint is that the king can only move to and be located on *allowed* cells throughout its path.

Recall that a chess king can move to any of its 8 neighboring cells in a single move. Two cells are considered neighboring if they share at least one point (i.e., they are horizontally, vertically, or diagonally adjacent).

## Input Format

(Standard input format for competitive programming problems typically follows this structure, though not explicitly provided in your prompt)

The first line contains two integers $x_0, y_0$ representing the starting coordinates.
The second line contains two integers $x_1, y_1$ representing the target coordinates.
The third line contains a single integer $n$ ($1 \le n \le 10^5$), the number of allowed segments.
The next $n$ lines each contain three integers $r_i, a_i, b_i$ ($1 \le r_i, a_i, b_i \le 10^9$; $a_i \le b_i$), describing an allowed segment.

## Output Format

Output a single integer: the minimum number of moves required. If it's impossible for the king to reach the target square, output -1.

## Constraints

* $1 \le x_0, y_0, x_1, y_1 \le 10^9$
* $1 \le n \le 10^5$
* $1 \le r_i, a_i, b_i \le 10^9$
* $a_i \le b_i$

## Example

(An example would typically be provided here to illustrate input/output)