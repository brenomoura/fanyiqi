from app.core.providers.cf2ast_facebook_m2m100_418M import CF2FastM2M100
from app.core.providers.facebook_m2m100_418M import FacebookM2M100
from app.core.providers.google_translate import GoogleTranslate

PROVIDERS = {
    "facebook/m2m100_418M": FacebookM2M100,
    "michaelfeil/ct2fast-m2m100_418M": CF2FastM2M100,
    "google-translate": GoogleTranslate,
}