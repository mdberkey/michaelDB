package main

import (
    "fmt"
)

// Internal Node Format (2B increments for memory boundry alignment)
// | checksum (2B) | nkeys (2B) | pointers (nkeys * 8B) |

// Leaf Node Format
// | checksum (2B) | nkeys (2B) | offsets (nkeys * 2B) | key-value pairs |

// KV pair Format (first key not needed since it's inherited from parent
// | klen (2B) | key | val |


