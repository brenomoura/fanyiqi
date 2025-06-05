import time

from transformers import M2M100ForConditionalGeneration, M2M100Tokenizer


def main():
    start_time = time.time()

    model = M2M100ForConditionalGeneration.from_pretrained("facebook/m2m100_418M")
    tokenizer = M2M100Tokenizer.from_pretrained("facebook/m2m100_418M")

    text = "Type shit"
    src_lang = "en"
    target_lang = "pt"
    tokenizer.src_lang = src_lang
    encoded = tokenizer(text, return_tensors="pt") # pt is not the src lang, it is the TensorType

    generated = model.generate(
        **encoded, forced_bos_token_id=tokenizer.get_lang_id(target_lang)
    )
    print(tokenizer.decode(generated[0], skip_special_tokens=True))

    print(f"Time elapsed {time.time() - start_time:.2f} seconds")


if __name__ == "__main__":
    main()
