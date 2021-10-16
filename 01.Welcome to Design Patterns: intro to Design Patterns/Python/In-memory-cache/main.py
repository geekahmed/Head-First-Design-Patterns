from cache import Cache
from lfu import LFU
from lru import LRU

lfu_cache = LFU()
new_cache = Cache(cache_strategy=lfu_cache)
new_cache.do_some_business_logic()

# Change Cache to LRU
lru_cache = LRU()
new_cache.set_strategy(lru_cache)

new_cache.do_some_business_logic()