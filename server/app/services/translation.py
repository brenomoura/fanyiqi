from typing import List, Tuple

from app.core.providers import PROVIDERS


class NotFoundModelException(Exception): ...


class TransalationError(Exception): ...


class TranslationService:
    def __init__(self, model_name: str):
        model = PROVIDERS.get(model_name)
        if model is None:
            raise NotFoundModelException
        self.model = model()

    def translate(self, text: str, source_language: str, target_language: str) -> str:
        try:
            return self.model.translate(text, source_language, target_language)
        except Exception:
            raise TransalationError

    def get_languages(self) -> List[Tuple[str, str]]:
        return self.model.get_languages()
