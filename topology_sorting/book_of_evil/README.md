# Book of Evil

Paladin Manaao caught the trail of the ancient Book of Evil in a swampy area. This area contains $n$ settlements numbered from $1$ to $n$. Moving through the swamp is very difficult, so people tramped exactly $n-1$ paths. Each of these paths connects some pair of settlements and is bidirectional. Moreover, it is possible to reach any settlement from any other one by traversing one or several paths.

The distance between two settlements is the minimum number of paths that have to be crossed to get from one settlement to the other one. Manaao knows that the Book of Evil has got a damage range $d$. This means that if the Book of Evil is located in some settlement, its damage (for example, emergence of ghosts and werewolves) affects other settlements at distance $d$ or less from the settlement where the Book resides.

Manaao has heard of $m$ settlements affected by the Book of Evil. Their numbers are $p_1, p_2, \ldots, p_m$. Note that the Book may be affecting other settlements as well, but this has not been detected yet. Manaao wants to determine which settlements may contain the Book. Help him with this difficult task.

## Input

The first line contains three space-separated integers $n$, $m$ and $d$ ($1 \le m \le n \le 100000$; $0 \le d \le n-1$). The second line contains $m$ distinct space-separated integers $p_1, p_2, \ldots, p_m$ ($1 \le p_i \le n$). Then $n-1$ lines follow, each line describes a path made in the area. A path is described by a pair of space-separated integers $a_i$ and $b_i$ representing the ends of this path.

## Output

Print a single number — the number of settlements that may contain the Book of Evil. It is possible that Manaao received some controversial information and there is no settlement that may contain the Book. In such case, print 0.