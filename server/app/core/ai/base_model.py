from abc import ABC, abstractmethod
from typing import List, Tuple


class BaseTranslationModel(ABC):
    @abstractmethod
    def get_model_name(self):
        raise NotImplementedError

    @abstractmethod
    def get_languages(self) -> List[Tuple[str, str]]:
        raise NotImplementedError
    
    @abstractmethod
    def translate(self, text: str, source_lang: str, target_lang: str) -> str:
        raise NotImplementedError