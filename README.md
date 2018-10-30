# Cache

General cache interface that provides a consistent API that various data stores can implement. This can enable easier transitioning of a cache's persistence layer as well as provide an easy way to compare performance of two competing data stores.

Check out `cache.go` to see the interface.

## Practical Usage

For more practical usage, check out the [omni-cache library](https://github.com/panoplymedia/omni-cache), which implements this interface.
