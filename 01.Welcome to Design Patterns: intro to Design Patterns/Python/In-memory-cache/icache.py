from abc import ABC, abstractmethod

class CacheStrategy(ABC):
    @abstractmethod
    def evict(self, storage):
        pass