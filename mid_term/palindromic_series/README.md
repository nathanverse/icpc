# Palindromic Series Problem

## Problem Description

Adobe's got a game for you! You'll get a number `N`, and your mission is to turn it into a special string of lowercase letters. Once you've got that string, you need to tell if it's a palindrome.

Here's the breakdown of how the string magic happens:

1.  **Digit-to-Letter Map:** Think of 'a' as 0, 'b' as 1, 'c' as 2, and so on, all the way up to 'j' for 9. This is your decoder ring.

2.  **Building the Base:** Take your number `N` and convert each of its digits into its corresponding letter based on the map. This sequence of letters forms your "base substring."
    * **Example:** If `N` is `61`, the `6` becomes `g`, and the `1` becomes `b`. So, your base substring is "gb".

3.  **Growing the String:** Now, sum up all the digits in `N`. This sum tells you exactly how long your *final* string needs to be. Take your base substring and repeat it until it reaches that total length. If it goes over, just snip it at the right spot.
    * **Example (continuing from `N = 61`):** The sum of digits for `61` is $6 + 1 = 7$. Your base substring "gb" needs to become 7 characters long. Repeating "gb" gives you "gbgbgbg".

4.  **The Palindrome Check:** Is your newly formed string a palindrome? That means it reads the same forwards and backward.
    * **Example (continuing):** "gbgbgbg" is indeed a palindrome!

## Input Format

* The first line of input contains an integer `T`, which is the number of test cases.
* Each of the following `T` lines will contain a single integer `N`.

**Important Notes:**
* Numbers `N` will *not* start with zero (e.g., `01` is not a valid input).
* Only digits from 0 to 9 will