from icache import CacheStrategy

class LFU(CacheStrategy):
    def evict(self):
        print("LFU is working within this cache")