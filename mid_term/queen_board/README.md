# 8-Queens Puzzle with Maximum Score

This problem is a variation of the classic N-Queens puzzle, specifically applied to an 8x8 chessboard. The goal is not just to find a valid placement of queens, but to find one that yields the highest possible score.

## Problem Description

1.  **8-Queens Constraint:** Place exactly 8 chess queens on an 8x8 chessboard such that no two queens attack each other. A queen attacks any piece in the same row, column, or along any diagonal.
2.  **Assigned Square Scores:** Each square on the 8x8 board has an integer score assigned to it, ranging from 0 to 99. These scores are part of the input.
3.  **Maximum Total Score:** Among all valid 8-queens configurations, find the one where the sum of the scores of the 8 squares occupied by the queens is maximized.
4.  **Output:** The maximum total score achievable. Optionally, the specific board configuration that results in this maximum score can also be outputted.

## Input

The input will consist of 8 lines, each containing 8 space-separated integers. These integers represent the scores assigned to each square of the 8x8 chessboard. The first line corresponds to row 0, the second to row 1, and so on.
