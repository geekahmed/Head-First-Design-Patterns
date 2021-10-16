from icache import CacheStrategy

class Cache():
    def __init__(self, cache_strategy : CacheStrategy) -> None:
        self.cache_strategy = cache_strategy
    
    def get_cache_strategy(self) -> CacheStrategy:
        return self.cache_strategy
    
    def set_strategy(self, cache_strategy: CacheStrategy) -> None:
        self.cache_strategy = cache_strategy
    

    def do_some_business_logic(self) -> None:
        self.cache_strategy.evict()