from typing import List, Tuple

from app.core.ai.models import MODELS


class NotFoundModelException(Exception): ...


class TransalationError(Exception): ...


class TranslationService:
    def __init__(self, model_name: str):
        model = MODELS.get(model_name)
        if model is None:
            raise NotFoundModelException
        self.model = model()

    def translate(self, text: str, source_lang: str, target_lang: str) -> str:
        try:
            return self.model.translate(text, source_lang, target_lang)
        except Exception:
            raise TransalationError

    def get_languages(self) -> List[Tuple[str, str]]:
        return self.model.get_languages()
