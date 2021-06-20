A simple hash table should be used to figure out what frame a given disk page occupies.

The hash table should be implemented (entirely in main memory) by using an array of pointers to 
lists of <page number, frame number> pairs. 

The array is called the directory and each list of pairs is called a bucket. Given a page number,
you should apply a hash function to find the directory entry pointing to the bucket that contains the frame number 
for this page, if the page is in the buffer pool. If you search the bucket and don't find a pair containing this page
number, the page is not in the pool. If you find such a pair, it will tell you the frame in which the page resides.
This is illustrated in Figure 1. 

Directory
------------------------------
|  <pageNum,frameNum>        |   <-----bucket
-----------------------------
|  <pageNum,frameNum>        |
-----------------------------
|  <pageNum,frameNum>        |
-----------------------------
|  <pageNum,frameNum>        |
-----------------------------


FrameTable
--------------------------------------
|  <pageNum,dirty,pincount>          |   <-----frame
-------------------------------------
|  <pageNum,dirty,pincount>          |
-------------------------------------
|  <pageNum,dirty,pincount>          |
-------------------------------------
|  <pageNum,dirty,pincount>          |
-------------------------------------

friend class BufHashTbl;

private:
BufHTEntry    *next;      // The next entry in this hashtable bucket.

    int            pageNo;    // This page number.
    int            frameNo;   // The frame we're stored in.


BufferManager
-------------------------------------------
|  BufHashTableEntry *ht[HTSIZE]          |       LinkList of HashTableEntries
-------------------------------------------
hash function
-------------------------------------------