Resources:
https://www.csd.uoc.gr/~hy460/pdf/p650-lehman.pdf
https://iq.opengenus.org/b-tree-search-insert-delete-operations/

Notes from Lehman and Yao's Efficient Locking for Concurrent Operations on B-trees
he disk is partitioned into sections of a fixed size“(physical pages; in this paper, these will correspond to logical nodes of the tree)

A process is allowed to lock and unlock a disk page, lock gives exclusive modifications to that page
but allowing other processes to read

Lowercase symbols (x, t, current, etc.) are used to refer to variables (including pointers) in the primary storage of a process. Uppercase symbols (A, B, C) are used to refer to blocks of primary storage. 


lo&(x) denotes the operation of locking the disk page to which x points. If 
the page is already locked then the process will wait for the unlock

unlock(x) releases a held lock

A <- get(x) reading memory block A, the contents of the disk page to which x points

put(A, x) writes contents of x into the memory block A

process to modify a page x

lock(x)
A <- get(x)
modify data in A
put(A, x)
unlock(x)

B*-tree described by Wedekind [15] (based on the B-tree defined by Bayer and McCreight [2] ). 

Criteria for B*-Tree
1) The root is a leaf or must have at least 2 children
2) each node except root must have at least k + 1 child(ren)
 - k is a tree parameter
 - 2k is the maximum number of elements in a leaf
3) Each path to the leaf is equal length in h
4) Each node has 2k + 1 children
5) keys of data stored in leaf nodes, contains pointer to records of the database

Sequencing
Provided a lower and upper bound on values stored between points to subtree.  With pointers in between to subtrees to
associated records.


search(15)
examine C; get ptr to y


--------->problem occurs here the insertion process has altered the tree

C <-  read(y)
error: not found

Insert(9)
A <- read(x)
examine A and to get ptr to y
A <- read(y)
insert 9, split 9 into A and B
put(B, y')
put(A, y)
add node x to point to node y'

Bayer and Schkoinick
"First modifiers lock upper sections of the tree with 
writer-exclusion locks (which only lock out other writers, 
not readers). When the actual modifications must be performed, exclusive 
locks are applied, mostly in lower sections of the tree. This sparse use of exclusive locks enhances the concurrency 
of the algorithm."

Miller and Snyder
pioneer and follower locks


Blink-tree for concurrency
modifies the b*tree by adding a link pointer field to each node
link points to the next node in the same tree level

provide additional methods for reaching a node during splitting
