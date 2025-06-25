import os
from huggingface_hub import snapshot_download

models = ["michaelfeil/ct2fast-m2m100_418M", "facebook/m2m100_418M"]

target_dir = os.path.join(
    os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "data", "models"
)

print(target_dir)

os.makedirs(target_dir, exist_ok=True)

for model_id in models:
    print(f"Downloading {model_id}...")
    snapshot_download(
        repo_id=model_id,
        local_dir=os.path.join(target_dir, model_id.replace("/", "_")),
        local_dir_use_symlinks=False,
    )
    print(
        f"Downloaded {model_id} to {os.path.join(target_dir, model_id.replace('/', '_'))}"
    )

print("All models downloaded.")
