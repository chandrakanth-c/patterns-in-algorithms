# Dynamic Programming (DP) - [Status:needs several updation / not ready yet]

## Sources filtered/used:
(i) https://www.geeksforgeeks.org/optimal-substructure-property-in-dynamic-programming-dp-2/ <br>
(ii) https://stackoverflow.com/questions/33563230/optimal-substructure-in-dynamic-programing

## Overview

* solves a given complex problem by breaking it into subproblems using recursions and storing the <br>
* results of subproblems to avoid computing the same results again. <br>
* like [divide and conquer](https://github.com/chandrakanth-c/patterns-in-algorithms/blob/main/divide-and-conquer.md) DP combines solutions to sub-problems.  

## Properties of a problem that suggests that it is a DP problem

2 main properties of a problem tha suggests that the given problem can be solved using DP are: <br>
(1) Overlapping Subproblems <br>
(2) Optimal Substructure 

## Overlapping Subproblem Property

### Sources filtered/used:
(i) https://www.geeksforgeeks.org/overlapping-subproblems-property-in-dynamic-programming-dp-1/

### Overview

Consider a java code (sourced from the above mentioned link) below:

```
static int fib(int n)
{
    if (n <= 1)
        return n;

    return fib(n - 1) + fib(n - 2);
}
```

In the above code, time complexity is exponential / O(2^n) and space complexity is linear / O(1).

