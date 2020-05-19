# 959. Regions Cut By Slashes
#### Medium

In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a /, \, or blank space.  These characters divide the square into contiguous regions.

(Note that backslash characters are escaped, so a \ is represented as "\\".)

Return the number of regions.

 

### Example 1:
```
Input:
[
  " /",
  "/ "
]
Output: 2
```
Explanation: The 2x2 grid is as follows:
![1.png](1.png)

### Example 2:

```
Input:
[
  " /",
  "  "
]
Output: 1
```
Explanation: The 2x2 grid is as follows:
![2.png](2.png)

### Example 3:

```
Input:
[
  "\\/",
  "/\\"
]
Output: 4
```
Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
The 2x2 grid is as follows:
![3.png](3.png)

### Example 4:

```
Input:
[
  "/\\",
  "\\/"
]
Output: 5
```
Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
The 2x2 grid is as follows:
![4.png](4.png)

### Example 5:

```
Input:
[
  "//",
  "/ "
]
Output: 3
```
Explanation: The 2x2 grid is as follows:
![5.png](5.png)

 

### Note:
```
1 <= grid.length == grid[0].length <= 30
grid[i][j] is either '/', '\', or ' '.
```

* DFS
* Union Find
* Graph