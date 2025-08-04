# Painting Fence

Bizon the Champion isn't just attentive, he's also very hardworking. He's decided to paint his old fence his favorite color, orange.

The fence consists of $n$ vertical planks, placed in a row with no gaps between them. The planks are numbered from left to right, starting from one. The $i$-th plank has a width of 1 meter and a height of $a_i$ meters.

Bizon the Champion bought a brush with a width of 1 meter. He can make both vertical and horizontal strokes. During a stroke, the brush's full surface must touch the fence at all times (see examples for better understanding).

What's the minimum number of strokes Bizon the Champion needs to fully paint the fence? Note that he's allowed to paint the same area of the fence multiple times.

---

### Input

The first line contains an integer $n$ ($1 \le n \le 5000$) — the number of fence planks.

The second line contains $n$ space-separated integers $a_1, a_2, \dots, a_n$ ($1 \le a_i \le 10^9$).

---

### Output

Print a single integer — the minimum number of strokes needed to paint the whole fence.