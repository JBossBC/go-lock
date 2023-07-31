# go-lock
distributed lock


## Introduce

The design style is similar to the low-level locking mechanism of Golang, which allows developers to easily migrate to the use of distributed locks, use hierarchical sleep mechanisms to prevent hunger in extreme states, and use a rollback mechanism at the user level to ensure the effectiveness of the locking mechanism



## Warning

Currently, only Redis standalone machines are supported