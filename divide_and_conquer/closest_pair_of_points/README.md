# The Closest Pair Problem

Given a set of points in a two-dimensional space, you will have to find the distance between the closest two points.

## Input Format

The input file contains several sets of input.
Each set of input starts with an integer $N$ ($0 \le N \le 10000$), which denotes the number of points in this set.
The next $N$ lines contain the coordinates of $N$ two-dimensional points. The first of the two numbers denotes the X-coordinate and the latter denotes the Y-coordinate.
The input is terminated by a set whose $N=0$. This set should not be processed.
The value of the coordinates will be less than $40000$ and non-negative.

## Output Format

For each set of input produce a single line of output containing a floating-point number (with four digits after the decimal point) which denotes the distance between the closest two points.
If there is no such two points in the input whose distance is less than $10000$, print the line 'INFINITY'.