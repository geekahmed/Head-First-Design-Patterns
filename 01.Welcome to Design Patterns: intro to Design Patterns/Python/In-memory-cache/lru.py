from icache import CacheStrategy

class LRU(CacheStrategy):
    def evict(self):
        print("LRU is working within this cache")