The file formatter serializes and deserializes tuple data.  The two main components are the reader
and writer.

Reader

The reader follows the similar format as the node in executor by implementing Next() and Execute() since file scan 
is one of the operating node that is invoked by other operators.  Next() will check if there is
a follow tuple, if there is, the function will invoke readTuple() and store the tuple in FileScanner.next

Writer

Serialized the header and tuples into a binary format. The length of the header and tuples are
prepended before the data for the reader to utilize

TODO:
 - handle bit rot
 - Handle overflow of file and write overflowed data to a new file
 - Compression
 - Handle files directly
 - Implement a multifilescan to read from multiple files 
 