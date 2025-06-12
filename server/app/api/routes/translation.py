from fastapi import APIRouter, HTTPException, Query
from pydantic import BaseModel

from app.core.ai.models import MODELS
from app.services.translation import TranslationService

DEFAULT_TRANSLATION_MODEL = "michaelfeil/ct2fast-m2m100_418M"

router = APIRouter(prefix="/translation", tags=["translation"])


class TranslationRequest(BaseModel):
    text: str
    source_language: str
    target_language: str
    model: str = DEFAULT_TRANSLATION_MODEL


class TranslationResponse(BaseModel):
    translated_text: str


def get_translation_service_with_param(param):
    return TranslationService(param)


@router.post("/", response_model=TranslationResponse)
async def translate(
    request: TranslationRequest,
):
    try:
        service = TranslationService(request.model)
        translated_text = service.translate(
            text=request.text,
            source_language=request.source_language,
            target_language=request.target_language,
        )
        return TranslationResponse(translated_text=translated_text)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/languages", response_model=list[tuple[str, str]])
async def get_available_languages(
    model: str = Query(
        default=DEFAULT_TRANSLATION_MODEL,
        description="Model parameter for TranslationService",
    ),
):
    try:
        service = TranslationService(model)
        languages = service.get_languages()
        return languages
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/models", response_model=list[str])
async def get_available_models():
    try:
        models = [model for model, _ in MODELS.items()]
        return models
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
