# 1 - Dynamic Programming (DP) 

## 1.1 - Sources for this section:
https://www.geeksforgeeks.org/optimal-substructure-property-in-dynamic-programming-dp-2/ <br>
https://stackoverflow.com/questions/33563230/optimal-substructure-in-dynamic-programing

## 1.2 - Overview

* solves a given complex problem by breaking it into subproblems using recursions and storing the results of subproblems to avoid computing the same results again. <br>
* like [divide and conquer](https://github.com/chandrakanth-c/patterns-in-algorithms/blob/main/divide-and-conquer.md) DP combines solutions to sub-problems.  

## 1.3 - Properties of a problem that suggests that it is a DP problem

2 main properties of a problem tha suggests that the given problem can be solved using DP are: <br>
* Overlapping Subproblems <br>
* Optimal Substructure 

## 1.4 - Overlapping Subproblem Property

### 1.4.1 - Sources for this section:
https://www.geeksforgeeks.org/overlapping-subproblems-property-in-dynamic-programming-dp-1/

### 1.4.2 - Overview

Consider a `finding a fibonacci series F(n) java code` (sourced from the above mentioned link) below:

F(n)=F(n-1)+F(n-2), where n>1, n=0 F(0)=0, n=1 F(1)=1 <br>
F(2)=F(1)+F(0)=1 <br>
F(3)=F(2)+F(1)=1+1=2 <br>
F(4)=F(3)+F(2)=2+1=3 and so on... <br>

So, we have the fib seq as 0,1,1,2,3,5,8....

```
static int fib(int n)
{
    if (n <= 1)
        return n;

    return fib(n - 1) + fib(n - 2);
}
```

In the above code, time complexity is exponential / O(2^n) and space complexity is linear / O(1). 

The reduce the time complexity we can store the subresults value (here it is of the form int). 
For example, if we consider F(5)=F(4)+F(3), in this we can see that F(3) is computed while computing 
F(4) and is computed while computing F(3).

The different ways to store these values of subproblems are:
* Memoization (Top Down)
* Tabulation (Bottom Up)

### 1.4.3 - Memoization (Top Down)

In the recursive version, use a global cache (to the class) either in the form of an 
array (with default value as -1) or using a Map. 

Basically, add all the computed values to the cache at compute time and retrive it 
at the start of the function.

In the below example, I am going to use Map<Integer,Integer> as cache.

```
static int fib(int n)
{
    if (n <= 1)
        return n;

    if (cache.containsKey(n)){
        return cache.get(n);
    }

    int subResult=fib(n - 1) + fib(n - 2);
    cache.put(n,subResult);
    return subResult;
}
```
Now, using memoization we have reduced the time complexity from O(2^n) to O(n).

### 1.4.4 - Tabulation (Bottom Up)

Tabulation approach fills all the values returned by subproblems in a bottom-up fashion
in an array (or a map) and return n(th) subproblem result.

```
static int fib(int n)
{
    //res[i] where 0<=i<=n stores the value of the i(th) subproblem.
    int[] res=new int[n+1];
    res[0]=0;
    res[1]=1;

    for(int i=2;i<=n;++){
        cache[i]=cache[i-1]+cache[i-2];
    }

    return res[n];
}
```

## 1.5 - Optimal Substructure Property

### 1.5.1 - Sources
https://www.geeksforgeeks.org/optimal-substructure-property-in-dynamic-programming-dp-2/

### 1.5.2 - Overview

* Optimal Substructure Property problems, like `Overlapping Subproblem Property`,
uses recursion to find compute solutions for each subproblem but would optimise a result
value in each recursion iteration.
* Finding either max or min of all the solutions of the sub problems.
* For example, `find shortest path` problem calculates min of all the paths that are
calculated in each iteration.
























