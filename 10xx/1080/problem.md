# 1080. Insufficient Nodes in Root to Leaf Paths
#### Medium

Given the root of a binary tree, consider all root to leaf paths: paths from the root to any leaf.  (A leaf is a node with no children.)

A node is insufficient if every such root to leaf path intersecting this node has sum strictly less than limit.

Delete all insufficient nodes simultaneously, and return the root of the resulting binary tree.

 

### Example 1:


![p1.png](p1.png)
```
Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
```
![p2.png](p2.png)
```
Output: [1,2,3,4,null,null,7,8,9,null,14]
```


### Example 2:
![p3.png](p3.png)
```
Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
```
![p4.png](p4.png)
```
Output: [5,4,8,11,null,17,4,7,null,null,null,5]
```
 

### Example 3:


![p5.png](p5.png)
```
Input: root = [1,2,-3,-5,null,4,null], limit = -1
```
![p6.png](p6.png)
```
Output: [1,null,-3,4]
```
 

### Note:

```
The given tree will have between 1 and 5000 nodes.
-10^5 <= node.val <= 10^5
-10^9 <= limit <= 10^9
```

* Tree
* DFS