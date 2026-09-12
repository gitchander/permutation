# Permutation Algorithm

The main concepts of this algorithm:

1. Using simple operations like swapping elements;
2. Having a simple rule to know which elements need to be swapped at the next step.

As a result of my experiments with element permutations, I realized that a single exchange is not enough. Another operation is needed: the "flip." It consists of multiple swaps.
Example flip of $n$ elements:

```text

	1 2 3 ... n
	<---------> - flip of n elements
	n ... 3 2 1


```

Using only the flip operation, you can generate all permutations of elements.
However, a challenge arises: determining how many elements need to be flipped at each step.
To solve this problem, I used an auxiliary array.
The size of this array equals the size of the elements array.
All its elements are integers and initialized to zero.

* The first element can take values: `[0]`
* The second element: `[0, 1]`
* The third element: `[0, 1, 2]`
* The 4th element: `[0, 1, 2, 3]`
* The $n$-th element: `[0, 1, ... n-1]`
...

No actions are needed for the first permutation.
For subsequent steps, to calculate the count of elements that need to be flipped, I use the index of the element in the auxiliary array.
For every iteration, the index is initially set to zero.
Then, `1` is added to the element at this index.
If this element is less than `index + 1`, then the number of first elements that need to be flipped is `index + 1`. Otherwise, the element at this index is reset to zero, and the index is incremented.
Next, check the element with the new index.
If the index equals the index of the last element, the permutation stops.

```text

+-----+-----------+-------+
| no. | elements  | flip  |
+-----+-----------+-------+
|     |           |       |
|  1  |  1  2  3  |       |
|     |  <-->     |   2   |
|  2  |  2  1  3  |       |
|     |  <----->  |   3   |
|  3  |  3  1  2  |       |
|     |  <-->     |   2   |
|  4  |  1  3  2  |       |
|     |  <----->  |   3   |
|  5  |  2  3  1  |       |
|     |  <-->     |   2   |
|  6  |  3  2  1  |       |
|     |           |       |
+-----+-----------+-------+

+-----+--------------+-------+
| no. |   elements   | flip  |
+-----+--------------+-------+
|     |              |       |
|  1  |  1  2  3  4  |       |
|     |  <-->        |   2   |
|  2  |  2  1  3  4  |       |
|     |  <----->     |   3   |
|  3  |  3  1  2  4  |       |
|     |  <-->        |   2   |
|  4  |  1  3  2  4  |       |
|     |  <----->     |   3   |
|  5  |  2  3  1  4  |       |
|     |  <-->        |   2   |
|  6  |  3  2  1  4  |       |
|     |  <-------->  |   4   |
|  7  |  4  1  2  3  |       |
|     |  <-->        |   2   |
|  8  |  1  4  2  3  |       |
|     |  <----->     |   3   |
|  9  |  2  4  1  3  |       |
|     |  <-->        |   2   |
| 10  |  4  2  1  3  |       |
|     |  <----->     |   3   |
| 11  |  1  2  4  3  |       |
|     |  <-->        |   2   |
| 12  |  2  1  4  3  |       |
|     |  <-------->  |   4   |
| 13  |  3  4  1  2  |       |
|     |  <-->        |   2   |
| 14  |  4  3  1  2  |       |
|     |  <----->     |   3   |
| 15  |  1  3  4  2  |       |
|     |  <-->        |   2   |
| 16  |  3  1  4  2  |       |
|     |  <----->     |   3   |
| 17  |  4  1  3  2  |       |
|     |  <-->        |   2   |
| 18  |  1  4  3  2  |       |
|     |  <-------->  |   4   |
| 19  |  2  3  4  1  |       |
|     |  <-->        |   2   |
| 20  |  3  2  4  1  |       |
|     |  <----->     |   3   |
| 21  |  4  2  3  1  |       |
|     |  <-->        |   2   |
| 22  |  2  4  3  1  |       |
|     |  <----->     |   3   |
| 23  |  3  4  2  1  |       |
|     |  <-->        |   2   |
| 24  |  4  3  2  1  |       |
|     |              |       |
+-----+--------------+-------+

flip examples:

1 2 3 4
<->         flip(2)
2 1 3 4

1 2 3 4
<--->       flip(3)
3 2 1 4

1 2 3 4
<----->     flip(4)
4 3 2 1

1 2 3 4 5
<->         flip(2)
2 1 3 4 5

1 2 3 4 5
<--->       flip(3)
3 2 1 4 5

1 2 3 4 5
<----->     flip(4)
4 3 2 1 5

1 2 3 4 5
<------->   flip(5)
5 4 3 2 1


Fs(B[n]) -> { B[n+1], fs }

fs - flip size

+----+---------+-----+
| no.|    B    | fs  |
+----+---------+-----+
|  1 | 0 0 0 0 |  -  |
|  2 | 0 1 0 0 |  2  |
|  3 | 0 0 1 0 |  3  |
|  4 | 0 1 1 0 |  2  |
|  5 | 0 0 2 0 |  3  |
|  6 | 0 1 2 0 |  2  |
|  7 | 0 0 0 1 |  4  |
|  8 | 0 1 0 1 |  2  |
|  9 | 0 0 1 1 |  3  |
| 10 | 0 1 1 1 |  2  |
| 11 | 0 0 2 1 |  3  |
| 12 | 0 1 2 1 |  2  |
| 13 | 0 0 0 2 |  4  |
| 14 | 0 1 0 2 |  2  |
| 15 | 0 0 1 2 |  3  |
| 16 | 0 1 1 2 |  2  |
| 17 | 0 0 2 2 |  3  |
| 18 | 0 1 2 2 |  2  |
| 19 | 0 0 0 3 |  4  |
| 20 | 0 1 0 3 |  2  |
| 21 | 0 0 1 3 |  3  |
| 22 | 0 1 1 3 |  2  |
| 23 | 0 0 2 3 |  3  |
| 24 | 0 1 2 3 |  2  |
+----+---------+-----+
| 25 | 0 0 0 0 | 4*  |
+----+---------+-----+

```

The last flip (4*) returns the array to the initial state.

Author: Manuilov Yaroslav (Chander)
Email:  jpochander@gmail.com

```text

       ---  ---
      /  | /
      \  | \
       ---  ---
       / | / |
      /  |/  |
     /       |
    /        |
-------   -------

```